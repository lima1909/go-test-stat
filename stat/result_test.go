package stat

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestResult(t *testing.T) {
	testJSON := `
	{"Time":"2018-09-11T18:00:23.085685436+02:00","Action":"run","Package":"go-test-stats/stat","Test":"TestStat"}
	{"Time":"2018-09-11T18:00:23.085861883+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"=== RUN   TestStat\n"}
	{"Time":"2018-09-11T18:00:23.085899738+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"--- FAIL: TestStat (0.00s)\n"}
	{"Time":"2018-09-11T18:00:23.085909975+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"    stat_test.go:30: no err expected, but: err by parsing json: unexpected end of JSON input\n"}
	{"Time":"2018-09-11T18:00:23.085917956+02:00","Action":"fail","Package":"go-test-stats/stat","Test":"TestStat"}
	{"Time":"2018-09-11T18:00:23.085924746+02:00","Action":"output","Package":"go-test-stats/stat","Output":"FAIL\n"}
	{"Time":"2018-09-11T18:00:23.090100045+02:00","Action":"output","Package":"go-test-stats/stat","Output":"FAIL\tgo-test-stats/stat\t0.006s\n"}
	{"Time":"2018-09-11T18:00:23.090126626+02:00","Action":"fail","Package":"go-test-stats/stat","Elapsed":0.006}
	{"Time":"2018-09-11T18:00:23.097910794+02:00","Action":"run","Package":"go-test-stats","Test":"TestFoo"}
	{"Time":"2018-09-11T18:00:23.09793498+02:00","Action":"output","Package":"go-test-stats","Test":"TestFoo","Output":"=== RUN   TestFoo\n"}
	{"Time":"2018-09-11T18:00:23.097946758+02:00","Action":"output","Package":"go-test-stats","Test":"TestFoo","Output":"--- PASS: TestFoo (0.00s)\n"}
	{"Time":"2018-09-11T18:00:23.097956666+02:00","Action":"pass","Package":"go-test-stats","Test":"TestFoo"}
	{"Time":"2018-09-11T18:00:23.097962849+02:00","Action":"run","Package":"go-test-stats","Test":"TestBar"}
	{"Time":"2018-09-11T18:00:23.097967347+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"=== RUN   TestBar\n"}
	{"Time":"2018-09-11T18:00:23.097973321+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"--- FAIL: TestBar (0.00s)\n"}
	{"Time":"2018-09-11T18:00:23.097978502+02:00","Action":"output","Package":"go-test-stats","Test":"TestBar","Output":"    main_test.go:10: simulate err\n"}
	{"Time":"2018-09-11T18:00:23.097983218+02:00","Action":"fail","Package":"go-test-stats","Test":"TestBar"}
	{"Time":"2018-09-11T18:00:23.097988169+02:00","Action":"output","Package":"go-test-stats","Output":"FAIL\n"}
	{"Time":"2018-09-11T18:00:23.098268351+02:00","Action":"output","Package":"go-test-stats","Output":"FAIL\tgo-test-stats\t0.002s\n"}
	{"Time":"2018-09-11T18:00:23.098288569+02:00","Action":"fail","Package":"go-test-stats","Elapsed":0.002}
	{"Time":"2018-09-11T18:00:23.098849031+02:00","Action":"output","Package":"go-test-stats/ioread","Output":"?   \tgo-test-stats/ioread\t[no test files]\n"}
	{"Time":"2018-09-11T18:00:23.098862696+02:00","Action":"skip","Package":"go-test-stats/ioread","Elapsed":0}`

	r, err := Handle([]byte(testJSON))
	if err != nil {
		t.Errorf("no err expected, but: %v", err)
	}

	if len(r.Packages) != 3 {
		t.Errorf("expect 3 packages, got: %v", len(r.Packages))
	}

	expect := []struct {
		tests   int
		elapsed float64
	}{
		{tests: 1, elapsed: 0.006},
		{tests: 2, elapsed: 0.002},
		{tests: 0, elapsed: 0},
	}
	for i, exp := range expect {
		p := r.Packages[i]
		if len(p.Tests) != exp.tests {
			t.Errorf("expect %v Test, got: %v", exp.tests, len(p.Tests))
		}
		if p.Elapsed != exp.elapsed {
			t.Errorf("expect %v Elapsed, got: %v", exp.elapsed, p.Elapsed)
		}
	}

	test := r.Packages[0].Tests[0]
	if test.Action != "run" {
		t.Errorf("run != %s", test.Action)
	}
	if test.Result != "fail" {
		t.Errorf("pass != %s", test.Result)
	}

	test = r.Packages[1].Tests[0]
	if test.Action != "run" {
		t.Errorf("run != %s", test.Action)
	}
	if test.Result != "pass" {
		t.Errorf("pass != %s", test.Result)
	}
}

func TestResult2(t *testing.T) {
	testJSON := `
	{"Time":"2018-09-28T16:02:55.9482323+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance","Output":"?   \tgithub.com/Nimsaja/DepotPerformance\t[no test files]\n"}
	{"Time":"2018-09-28T16:02:56.0012703+02:00","Action":"skip","Package":"github.com/Nimsaja/DepotPerformance","Elapsed":0.053}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestAddStock"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestAddStock","Output":"=== RUN   TestAddStock\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestAddStock","Output":"--- PASS: TestAddStock (0.00s)\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestAddStock","Elapsed":0}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestInitDefaultValues"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestInitDefaultValues","Output":"=== RUN   TestInitDefaultValues\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestInitDefaultValues","Output":"--- PASS: TestInitDefaultValues (0.00s)\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestInitDefaultValues","Elapsed":0}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestSum"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestSum","Output":"=== RUN   TestSum\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestSum","Output":"--- PASS: TestSum (0.00s)\n"}
	{"Time":"2018-09-28T16:02:56.4950206+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/depot","Test":"TestSum","Elapsed":0}
	{"Time":"2018-09-28T16:02:56.495991+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Output":"PASS\n"}
	{"Time":"2018-09-28T16:02:56.5041488+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/depot","Output":"ok  \tgithub.com/Nimsaja/DepotPerformance/depot\t0.326s\n"}
	{"Time":"2018-09-28T16:02:56.5041488+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/depot","Elapsed":0.326}
	{"Time":"2018-09-28T16:02:56.51322+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/store","Output":"?   \tgithub.com/Nimsaja/DepotPerformance/store\t[no test files]\n"}
	{"Time":"2018-09-28T16:02:56.51322+02:00","Action":"skip","Package":"github.com/Nimsaja/DepotPerformance/store","Elapsed":0}
	{"Time":"2018-09-28T16:02:57.1148653+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfResultArrayIsEmpty"}
	{"Time":"2018-09-28T16:02:57.1148653+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfResultArrayIsEmpty","Output":"=== RUN   TestConvertJSON2ResultIfResultArrayIsEmpty\n"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfResultArrayIsEmpty","Output":"--- PASS: TestConvertJSON2ResultIfResultArrayIsEmpty (0.00s)\n"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfResultArrayIsEmpty","Elapsed":0}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfStockIsFound"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfStockIsFound","Output":"=== RUN   TestConvertJSON2ResultIfStockIsFound\n"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfStockIsFound","Output":"--- PASS: TestConvertJSON2ResultIfStockIsFound (0.00s)\n"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestConvertJSON2ResultIfStockIsFound","Elapsed":0}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"run","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestGet"}
	{"Time":"2018-09-28T16:02:57.1158641+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestGet","Output":"=== RUN   TestGet\n"}
	{"Time":"2018-09-28T16:02:57.4069192+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestGet","Output":"--- PASS: TestGet (0.29s)\n"}
	{"Time":"2018-09-28T16:02:57.4069192+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Test":"TestGet","Elapsed":0.29}
	{"Time":"2018-09-28T16:02:57.4069192+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Output":"PASS\n"}
	{"Time":"2018-09-28T16:02:57.4175552+02:00","Action":"output","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Output":"ok  \tgithub.com/Nimsaja/DepotPerformance/yahoo\t1.021s\n"}
	{"Time":"2018-09-28T16:02:57.4175552+02:00","Action":"pass","Package":"github.com/Nimsaja/DepotPerformance/yahoo","Elapsed":1.021}	
`

	r, err := Handle([]byte(testJSON))
	if err != nil {
		t.Errorf("no err expected, but: %v", err)
	}

	if len(r.Packages) != 4 {
		t.Errorf("expect 4 packages, got: %v", len(r.Packages))
	}

	expect := []struct {
		tests   int
		elapsed float64
		action  string
	}{
		{tests: 0, action: actionSkip, elapsed: 0.053},
		{tests: 3, action: actionPass, elapsed: 0.326},
		{tests: 0, action: actionSkip, elapsed: 0},
		{tests: 3, action: actionPass, elapsed: 1.021},
	}
	for i, exp := range expect {
		p := r.Packages[i]
		if len(p.Tests) != exp.tests {
			t.Errorf("expect %v Test, got: %v", exp.tests, len(p.Tests))
		}
		if p.Elapsed != exp.elapsed {
			t.Errorf("expect %v Elapsed, got: %v", exp.elapsed, p.Elapsed)
		}
		if p.Action != exp.action {
			t.Errorf("expect %s Action, got: %s", exp.action, p.Action)
		}
	}
}

func TestNotJsonResult(t *testing.T) {
	r, err := Handle([]byte("NOT JSON"))
	if err == nil {
		t.Errorf("err expected, but is %v", err)
	}
	if len(r.Packages) > 0 {
		t.Errorf("no packages expected, got: %v", len(r.Packages))
	}
}

// only for see the skip in the Overview
func TestSkip(t *testing.T) {
	t.SkipNow()
}

// test all JSON files from GOROOT - test2json dir
func TestGoTest2JSON(t *testing.T) {
	goroot, found := os.LookupEnv("GOROOT")
	if found {
		testdatadir := filepath.Join(goroot, "src", "cmd", "internal", "test2json", "testdata")
		files, err := ioutil.ReadDir(testdatadir)
		if err != nil {
			t.Errorf("no err expected: %v", err)
		}
		for _, f := range files {
			t.Run(f.Name(), func(t *testing.T) {
				// test onl JSON-files
				if filepath.Ext(f.Name()) == ".json" {
					tf, err := os.Open(filepath.Join(testdatadir, f.Name()))
					if err != nil {
						t.Errorf("no err expected: %v", err)
					}
					defer tf.Close()

					b, err := ioutil.ReadAll(tf)
					if err != nil {
						t.Errorf("no err expected: %v", err)
					}
					r, err := Handle(b)
					if err != nil {
						t.Errorf("no err expected: %v", err)
					}
					if len(r.Packages) == 0 {
						t.Errorf("packaes expected: %v", len(r.Packages))
					}
					pckg := r.Packages[0]
					if len(pckg.Tests) == 0 && f.Name() != "benchshort.json" {
						t.Errorf("tests expected: %v", len(pckg.Tests))
					}
				}
			})

		}
	} else {
		t.Errorf("no env GOROOT set")
	}
}
