package stat

import "testing"

func TestParser(t *testing.T) {
	testJSONParse := `
{"Time":"2018-09-11T18:00:23.085685436+02:00","Action":"run","Package":"go-test-stats/stat","Test":"TestStat"}
{"Time":"2018-09-11T18:00:23.085861883+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"=== RUN   TestStat\n"}
{"Time":"2018-09-11T18:00:23.085899738+02:00","Action":"output","Package":"go-test-stats/stat","Test":"TestStat","Output":"--- FAIL: TestStat (0.00s)\n"}
`
	count := 0
	err := parse([]byte(testJSONParse), func(e TestEvent) {
		count++
	})
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}
	if count != 3 {
		t.Errorf("3 != %v", count)
	}
}
