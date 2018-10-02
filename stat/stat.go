package stat

import "log"

// Stat is the statistics from tes result
type Stat struct {
	Packages      int
	Tests         int
	Elapsed       float64
	Pass          int
	Fail          int
	Skip          int
	EmptyPackages []string
	result        Result
}

// New stats based on the result
func New(r Result) (stat Stat) {
	stat.result = r
	stat.Packages = len(r.Packages)
	for _, p := range r.Packages {
		stat.Elapsed += p.Elapsed
		lenTests := len(p.Tests)
		stat.Tests += lenTests
		if lenTests == 0 {
			stat.EmptyPackages = append(stat.EmptyPackages, p.Name)
		}

		for _, t := range p.Tests {
			switch t.Result {
			case actionPass:
				stat.Pass++
			case actionFail:
				stat.Fail++
			case actionSkip:
				stat.Skip++
			default:
				log.Fatalf("Invalid Action in Result: %s", t.Result)
			}

		}
	}
	return
}
