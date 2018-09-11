package stat

import (
	"testing"
)

var (
	testJSON = `{"Time":"2018-09-11T18:00:23.085685436+02:00","Action":"run","Package":"go-test-stats/stat","Test":"TestStat"}
{"Time":"2018-09-11T18:00:23.085861883+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"=== RUN   TestStat\n"}
{"Time":"2018-09-11T18:00:23.085899738+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"--- FAIL: TestStat (0.00s)\n"}
{"Time":"2018-09-11T18:00:23.085909975+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"    stat_test.go:30: no err expected, but: err by parsing json: unexpected end of JSON input\n"}
{"Time":"2018-09-11T18:00:23.085917956+02:00","Action":"fail","Package":"go-test-stats/stat","Test":"TestStat","Elapsed":0}
{"Time":"2018-09-11T18:00:23.085924746+02:00","Action":"output","Package":"go-test-stats/stat","Output":"FAIL\n"}
{"Time":"2018-09-11T18:00:23.090100045+02:00","Action":"output","Package":"go-test-stats/stat","Output":"FAIL\tgo-test-stats/stat\t0.006s\n"}
{"Time":"2018-09-11T18:00:23.090126626+02:00","Action":"fail","Package":"go-test-stats/stat","Elapsed":0.006}
{"Time":"2018-09-11T18:00:23.097910794+02:00","Action":"run","Package":"go-test-stats","Test":"TestFoo"}
{"Time":"2018-09-11T18:00:23.09793498+02:00","Action":"output","Package":"go-test-stats","Test":"TestFoo","Output":"=== RUN   TestFoo\n"}
{"Time":"2018-09-11T18:00:23.097946758+02:00","Action":"output","Package":"go-test-stats","Test":"TestFoo","Output":"--- PASS: TestFoo (0.00s)\n"}
{"Time":"2018-09-11T18:00:23.097956666+02:00","Action":"pass","Package":"go-test-stats","Test":"TestFoo","Elapsed":0}
{"Time":"2018-09-11T18:00:23.097962849+02:00","Action":"run","Package":"go-test-stats","Test":"TestBar"}
{"Time":"2018-09-11T18:00:23.097967347+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"=== RUN   TestBar\n"}
{"Time":"2018-09-11T18:00:23.097973321+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"--- FAIL: TestBar (0.00s)\n"}
{"Time":"2018-09-11T18:00:23.097978502+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"    main_test.go:10: simulate err\n"}
{"Time":"2018-09-11T18:00:23.097983218+02:00","Action":"fail","Package":"go-test-stats","Test":"TestBar","Elapsed":0}
{"Time":"2018-09-11T18:00:23.097988169+02:00","Action":"output","Package":"go-test-stats","Output":"FAIL\n"}
{"Time":"2018-09-11T18:00:23.098268351+02:00","Action":"output","Package":"go-test-stats","Output":"FAIL\tgo-test-stats\t0.002s\n"}
{"Time":"2018-09-11T18:00:23.098288569+02:00","Action":"fail","Package":"go-test-stats","Elapsed":0.002}
{"Time":"2018-09-11T18:00:23.098849031+02:00","Action":"output","Package":"go-test-stats/ioread","Output":"?   \tgo-test-stats/ioread\t[no test files]\n"}
{"Time":"2018-09-11T18:00:23.098862696+02:00","Action":"skip","Package":"go-test-stats/ioread","Elapsed":0}`
)

func TestResult(t *testing.T) {
	r, err := Handle([]byte(testJSON))
	if err != nil {
		t.Errorf("no err expected, but: %v", err)
	}

	if len(r.Packages) != 3 {
		t.Errorf("expect 3 packages, got: %v", len(r.Packages))
	}

	expect := []int{1, 2, 0}
	for i, exp := range expect {
		p := r.Packages[i]
		if len(p.Tests) != exp {
			t.Errorf("expect %v Test, got: %v", exp, len(p.Tests))
		}
	}
}

func TestSkip(t *testing.T) {
	t.SkipNow()
}
