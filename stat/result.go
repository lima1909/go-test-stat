package stat

import "fmt"

const (
	actionRun    string = "run"
	actionBench  string = "bench"
	actionPass   string = "pass"
	actionFail   string = "fail"
	actionSkip   string = "skip"
	actionPause  string = "pause"
	actionCont   string = "cont"
	actionOutput string = "output"
)

// Result is a container for all Packages and their Tests
type Result struct {
	Packages []*Package
}

// Package which are contains Tests
type Package struct {
	Name    string
	Action  string
	Elapsed float64
	Tests   []*Test
}

func (p *Package) String() string {
	return fmt.Sprintf("%s -> %s - %v (%v)", p.Name, p.Action, len(p.Tests), p.Elapsed)
}

// Test with Action and their Result
type Test struct {
	Name   string
	Action string
	Result string
}

func (t Test) String() string {
	return fmt.Sprintf("%s: %s -> %s", t.Name, t.Action, t.Result)
}

// Handle the json-byte-stream and interprete it to a result
func Handle(b []byte) (r Result, err error) {
	return r, parseF(b, func(e TestEvent) {

		packg := r.getPackageByName(e.Package)

		// no Test means package
		if e.Test == "" {
			if packg.Name != e.Package {
				panic(fmt.Sprintf("Expected package: %s, got: %s", packg.Name, e.Package))
			}
			packg.Action = e.Action
			packg.Elapsed = e.Elapsed
		} else {
			// create a new Tests
			if e.Action == actionRun || e.Action == actionBench {
				test := &Test{Name: e.Test, Action: e.Action}
				packg.Tests = append(packg.Tests, test)
			} else if e.Action == actionPass || e.Action == actionFail || e.Action == actionSkip {
				test := packg.getTestByName(e.Test)
				if test == nil {
					test = &Test{Name: e.Test, Action: "--"}
					packg.Tests = append(packg.Tests, test)
				}
				// Test is to end, save the result
				test.Result = e.Action
			}
		}

	},
		func(e TestEvent) bool {
			// ignore all outputs
			return e.Action == actionOutput
		},
	)
}

func (r *Result) getPackageByName(name string) *Package {
	for _, p := range r.Packages {
		if name == p.Name {
			return p
		}
	}

	packg := &Package{Name: name}
	r.Packages = append(r.Packages, packg)
	return packg
}

func (p *Package) getTestByName(name string) *Test {
	for _, t := range p.Tests {
		if name == t.Name {
			return t
		}
	}
	return nil
}
