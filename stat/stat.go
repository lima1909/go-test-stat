package stat

import (
	"fmt"
)

// Overview collect all satistics
type Overview struct {
	Packages            int
	Tests               int
	TestsPerPackage     map[string]int
	TestsPerPackagePass map[string]int
	TestsPerPackageFail map[string]int
	TestsPerPackageSkip map[string]int
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

// Print the overview (all results)
func (o *Overview) Print() {
	fmt.Printf("Packages: %v\n", o.Packages)
	fmt.Printf("Tests: %v\n", o.Tests)
	fmt.Printf("Pass/Fail/Skip: %v / %v / %v\n", o.Pass(), o.Fail(), o.Skip())
}

// Calculate stats based on the result
func Calculate(r Result) *Overview {
	overv := &Overview{
		Packages:            len(r.Packages),
		TestsPerPackage:     make(map[string]int),
		TestsPerPackagePass: make(map[string]int),
		TestsPerPackageFail: make(map[string]int),
		TestsPerPackageSkip: make(map[string]int),
	}
	for _, pckg := range r.Packages {
		overv.TestsPerPackage[pckg.Name] = 0
		overv.TestsPerPackagePass[pckg.Name] = 0
		overv.TestsPerPackageFail[pckg.Name] = 0
		overv.TestsPerPackageSkip[pckg.Name] = 0

		for _, t := range pckg.Tests {
			overv.Tests++
			val := overv.TestsPerPackage[pckg.Name]
			overv.TestsPerPackage[pckg.Name] = val + 1

			if t.Result == "pass" {
				val := overv.TestsPerPackagePass[pckg.Name]
				overv.TestsPerPackagePass[pckg.Name] = val + 1
			} else if t.Result == "fail" {
				val := overv.TestsPerPackageFail[pckg.Name]
				overv.TestsPerPackageFail[pckg.Name] = val + 1
			} else if t.Result == "skip" {
				val := overv.TestsPerPackageSkip[pckg.Name]
				overv.TestsPerPackageSkip[pckg.Name] = val + 1
			}
		}
	}

	return overv
}
