package stat

import (
	"log"
)

const (
	actionRun    string = "run"
	actionPass   string = "pass"
	actionFail   string = "fail"
	actionSkip   string = "skip"
	actionPause  string = "pause"
	actionCont   string = "cont"
	actionBench  string = "bench"
	actionOutput string = "output"
)

// Overview collect all satistics
type Overview struct {
	Packages            int
	Tests               int
	Elapsed             float64
	TestsPerPackage     map[string]int
	TestsPerPackagePass map[string]int
	TestsPerPackageFail map[string]int
	TestsPerPackageSkip map[string]int
}

// New instance of Overview
func New() *Overview {
	return &Overview{
		TestsPerPackage:     make(map[string]int),
		TestsPerPackagePass: make(map[string]int),
		TestsPerPackageFail: make(map[string]int),
		TestsPerPackageSkip: make(map[string]int),
	}
}

// Calculate stats based on the result
func Calculate(r Result) Overview {
	overv := New()
	overv.Packages = len(r.Packages)

	for _, pckg := range r.Packages {
		overv.TestsPerPackage[pckg.Name] = 0
		overv.TestsPerPackagePass[pckg.Name] = 0
		overv.TestsPerPackageFail[pckg.Name] = 0
		overv.TestsPerPackageSkip[pckg.Name] = 0

		overv.Tests = len(pckg.Tests)
		for _, t := range pckg.Tests {
			overv.Elapsed = overv.Elapsed + t.Elapsed

			val := overv.TestsPerPackage[pckg.Name]
			overv.TestsPerPackage[pckg.Name] = val + 1

			switch t.Result {
			case actionPass:
				val := overv.TestsPerPackagePass[pckg.Name]
				overv.TestsPerPackagePass[pckg.Name] = val + 1
			case actionFail:
				val := overv.TestsPerPackageFail[pckg.Name]
				overv.TestsPerPackageFail[pckg.Name] = val + 1
			case actionSkip:
				val := overv.TestsPerPackageSkip[pckg.Name]
				overv.TestsPerPackageSkip[pckg.Name] = val + 1
			default:
				log.Fatalf("Invalid Action in Result: %s", t.Result)
			}
		}
	}

	return *overv
}

// EmptyPackages Packages without Tests
func (o Overview) EmptyPackages() []string {
	names := make([]string, 0)
	for k, v := range o.TestsPerPackage {
		if v == 0 {
			names = append(names, k)
		}
	}
	return names
}

func count(m map[string]int) int {
	val := 0
	for _, v := range m {
		val += v
	}
	return val
}

// Pass is the sum of all succesfull test
func (o Overview) Pass() int {
	return count(o.TestsPerPackagePass)
}

// Fail is the sum of all failed test
func (o Overview) Fail() int {
	return count(o.TestsPerPackageFail)
}

// Skip is the sum of all skiped test
func (o Overview) Skip() int {
	return count(o.TestsPerPackageSkip)
}

// Percent calculate percent of all Tests (e.g. percent from Pass)
func (o Overview) Percent(v int) int {
	r := v * 100 / o.Tests
	if r > 0 && r < 10 {
		r = 10
	}
	return int(r) / 10
}
