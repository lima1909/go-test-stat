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
	Name   string
	Action string
	Result string
}

func (t Test) String() string {
	return fmt.Sprintf("%s: %s -> %s", t.Name, t.Action, t.Result)
}

// Handle ...
func Handle(b []byte) (Result, error) {
	var r Result
	var (
		currentPackage string
		currentTest    string
	)

	err := Parse(b, func(e TestEvent) {

		if e.Package != currentPackage {
			currentPackage = e.Package
			r.Packages = append(r.Packages, &Package{Name: currentPackage})
		}
		// Tests
		if e.Test != "" && e.Test != currentTest {
			currentTest = e.Test
			idx := len(r.Packages) - 1
			packg := r.Packages[idx]
			packg.Tests = append(packg.Tests, &Test{Name: currentTest, Action: e.Action})
		} else if e.Test != "" {
			idx := len(r.Packages) - 1
			packg := r.Packages[idx]
			idx = len(packg.Tests) - 1
			tst := packg.Tests[idx]
			tst.Result = e.Action
		}
	})

	if err != nil {
		return r, err
	}
	return r, nil
}
