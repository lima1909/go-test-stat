package iowrite

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/lima1909/go-test-stat/stat"
)

const (
	red    = 31
	green  = 32
	yellow = 33
)

var (
	resultFail = color("✖", red)
	resultPass = color("✓", green)
	resultSkip = color("↷", yellow)

	allTestPass        = "\U0001f603   "
	oneOrMoreTestsFail = "\U0001f61f   "
)

func color(str string, color int) string {
	return fmt.Sprintf("\x1b[1;%dm %s \x1b[0m", color, str)
}

// more than zero fails
func resultPercent(o stat.Overview) string {
	f := o.Percent(o.Fail())
	s := o.Percent(o.Skip())
	p := 10 - f - s

	return strings.Repeat(resultPass, p) + strings.Repeat(resultFail, f) + strings.Repeat(resultSkip, s)
}

// emoji fail: \U0001f44e  \U0001f61f  \U0001f620  \U0000274C
// emoji pass: \U0001f44d  \U0001f603  \U0001f917  \U000023E9
func resultTotal(o stat.Overview) string {
	if o.Fail() > 0 {
		return oneOrMoreTestsFail
	}
	return allTestPass
}

// Print the overview (all results)
func Print(o stat.Overview) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Total\tTests\tPass\tFail\tSkip\t")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t\n", resultTotal(o), o.Tests, o.Pass(), o.Fail(), o.Skip())
	w.Flush()

	fmt.Printf("%s\n", resultPercent(o))
	fmt.Printf("Elapsed: %v\n", o.Elapsed)
	fmt.Printf("Packages without Tests: %v\n", o.EmptyPackages())
}
