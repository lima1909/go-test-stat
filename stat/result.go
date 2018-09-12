package stat

import "fmt"

// Result is a container for all Packages and their Tests
type Result struct {
	Packages []*Package
}

// Package which are contains Tests
type Package struct {
	Name  string
	Tests []*Test
}

func (p *Package) String() string {
	return fmt.Sprintf("%s - %v", p.Name, len(p.Tests))
}

// Test with Action and their Result
type Test struct {
	Name    string
	Action  string
	Result  string
	Elapsed float64
}

func (t Test) String() string {
	return fmt.Sprintf("%s: %s -> %s (%v)", t.Name, t.Action, t.Result, t.Elapsed)
}

// Handle the json-byte-stream and interprete it to a result
func Handle(b []byte) (r Result, err error) {
	return r, Parse(b, func(e TestEvent) {
		packg := r.getPackageByName(e.Package)
		// create new package
		if packg == nil {
			packg = &Package{Name: e.Package}
			r.Packages = append(r.Packages, packg)
		}

		test := packg.getTestByName(e.Test)
		// create a new Tests
		if e.Test != "" && test == nil {
			test := &Test{Name: e.Test, Action: e.Action, Elapsed: e.Elapsed}
			packg.Tests = append(packg.Tests, test)
		} else if e.Test != "" {
			// Test is to end, save the result
			test.Result = e.Action
		}
	})
}

func (r *Result) getPackageByName(name string) *Package {
	for _, p := range r.Packages {
		if name == p.Name {
			return p
		}
	}
	return nil
}

func (p *Package) getTestByName(name string) *Test {
	for _, t := range p.Tests {
		if name == t.Name {
			return t
		}
	}
	return nil
}
