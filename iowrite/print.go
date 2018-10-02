package iowrite

import (
	"fmt"
	"os"
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

// emoji fail: \U0001f44e  \U0001f61f  \U0001f620  \U0000274C
// emoji pass: \U0001f44d  \U0001f603  \U0001f917  \U000023E9
func resultTotal(s stat.Stat) string {
	if s.Fail > 0 {
		return oneOrMoreTestsFail
	}
	return allTestPass
}

// Print the overview (all results)
func Print(s stat.Stat) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Total\tTests\tPass\tFail\tSkip\t")
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t\n", resultTotal(s), s.Tests, s.Pass, s.Fail, s.Skip)
	w.Flush()

	fmt.Printf("Elapsed: %v\n", s.Elapsed)
	fmt.Printf("Packages without Tests (%v/%v):\n%v \n", s.Packages, len(s.EmptyPackages), s.EmptyPackages)
}
