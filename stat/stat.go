package stat

import (
	"fmt"
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
func Calculate(r Result) *Overview {
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

	return overv
}

// Print the overview (all results)
func (o *Overview) Print() {
	// fmt.Printf("Tests: %v (Elapsed: %v)\n", o.Tests, o.Elapsed)
	fmt.Printf("Tests: %v \n", o.Tests)
	fmt.Printf("Pass|Fail|Skip: %v | %v | %v\n", o.Pass(), o.Fail(), o.Skip())
	fmt.Printf("-> %s\n", o.result())
	fmt.Printf("Elapsed: %v\n", o.Elapsed)
	fmt.Printf("Packages without Tests: %v\n", o.emptyPackages())
}

const (
	red    = 31
	green  = 32
	yellow = 33
)

func color(str string, color int) string {
	return fmt.Sprintf("\x1b[1;%dm %s \x1b[0m", color, str)
}

// more than zero fails
func (o *Overview) result() string {

	if o.Fail() > 0 {
		return "\U0001f44e  \U0001f61f  \U0001f620  \U0000274C  " + color("✖", red)
	}
	return "\U0001f44d  \U0001f603  \U0001f917  \U000023E9  " + color("✓", green)
}

// Packages without Tests
func (o *Overview) emptyPackages() []string {
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
func (o *Overview) Pass() int {
	return count(o.TestsPerPackagePass)
}

// Fail is the sum of all failed test
func (o *Overview) Fail() int {
	return count(o.TestsPerPackageFail)
}

// Skip is the sum of all skiped test
func (o *Overview) Skip() int {
	return count(o.TestsPerPackageSkip)
}
