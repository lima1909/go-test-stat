package stat

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

// for test of the Elapsed works
func TestSleep(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Int63n(999)
	d := time.Duration(r)
	time.Sleep(time.Millisecond * d)
}

func TestFail(t *testing.T) {
	if len(os.Getenv("TESTFAIL")) > 0 {
		t.Errorf("Fail Example")
	}
}

var (
	jsonTestTaskOrganizer = `
	{"Time":"2018-09-14T17:23:07.995803342+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestGetNextMonthForTask"}
	{"Time":"2018-09-14T17:23:07.996044948+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestGetNextMonthForTask","Output":"=== RUN   TestGetNextMonthForTask\n"}
	{"Time":"2018-09-14T17:23:07.99606333+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestGetNextMonthForTask","Output":"--- PASS: TestGetNextMonthForTask (0.00s)\n"}
	{"Time":"2018-09-14T17:23:07.996070562+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestGetNextMonthForTask","Elapsed":0}
	{"Time":"2018-09-14T17:23:07.996082212+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestRecalculateNextMonthProp"}
	{"Time":"2018-09-14T17:23:07.996088218+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestRecalculateNextMonthProp","Output":"=== RUN   TestRecalculateNextMonthProp\n"}
	{"Time":"2018-09-14T17:23:07.996094617+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestRecalculateNextMonthProp","Output":"--- PASS: TestRecalculateNextMonthProp (0.00s)\n"}
	{"Time":"2018-09-14T17:23:07.996103779+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestRecalculateNextMonthProp","Elapsed":0}
	{"Time":"2018-09-14T17:23:07.996109862+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestCalculateMonthList"}
	{"Time":"2018-09-14T17:23:07.996114996+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestCalculateMonthList","Output":"=== RUN   TestCalculateMonthList\n"}
	{"Time":"2018-09-14T17:23:07.996122926+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestCalculateMonthList","Output":"--- PASS: TestCalculateMonthList (0.00s)\n"}
	{"Time":"2018-09-14T17:23:07.996127805+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer/task","Test":"TestCalculateMonthList","Elapsed":0}
	{"Time":"2018-09-14T17:23:07.99613254+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Output":"PASS\n"}
	{"Time":"2018-09-14T17:23:07.996137424+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer/task","Output":"ok  \tgithub.com/Nimsaja/TaskOrganizer/task\t(cached)\n"}
	{"Time":"2018-09-14T17:23:07.996144633+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer/task","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.028004468+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestTaskList"}
	{"Time":"2018-09-14T17:23:08.028131979+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestTaskList","Output":"=== RUN   TestTaskList\n"}
	{"Time":"2018-09-14T17:23:08.028155109+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestTaskList","Output":"--- PASS: TestTaskList (0.00s)\n"}
	{"Time":"2018-09-14T17:23:08.028165521+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestTaskList","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.028174033+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheFirstTime"}
	{"Time":"2018-09-14T17:23:08.028182494+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheFirstTime","Output":"=== RUN   TestRecalcOfNextMonthIfRunTheFirstTime\n"}
	{"Time":"2018-09-14T17:23:08.028191264+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheFirstTime","Output":"--- PASS: TestRecalcOfNextMonthIfRunTheFirstTime (0.00s)\n"}
	{"Time":"2018-09-14T17:23:08.028202631+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheFirstTime","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.028208973+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheSecondTime"}
	{"Time":"2018-09-14T17:23:08.028244241+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheSecondTime","Output":"=== RUN   TestRecalcOfNextMonthIfRunTheSecondTime\n"}
	{"Time":"2018-09-14T17:23:08.028266206+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheSecondTime","Output":"--- PASS: TestRecalcOfNextMonthIfRunTheSecondTime (0.00s)\n"}
	{"Time":"2018-09-14T17:23:08.028276015+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfRunTheSecondTime","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.028282312+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfMonthChanged"}
	{"Time":"2018-09-14T17:23:08.028286908+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfMonthChanged","Output":"=== RUN   TestRecalcOfNextMonthIfMonthChanged\n"}
	{"Time":"2018-09-14T17:23:08.028295427+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfMonthChanged","Output":"--- PASS: TestRecalcOfNextMonthIfMonthChanged (0.00s)\n"}
	{"Time":"2018-09-14T17:23:08.028301822+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestRecalcOfNextMonthIfMonthChanged","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.02830922+02:00","Action":"run","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestMonthTasks"}
	{"Time":"2018-09-14T17:23:08.028314709+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestMonthTasks","Output":"=== RUN   TestMonthTasks\n"}
	{"Time":"2018-09-14T17:23:08.028320623+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestMonthTasks","Output":"--- PASS: TestMonthTasks (0.00s)\n"}
	{"Time":"2018-09-14T17:23:08.028331763+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Test":"TestMonthTasks","Elapsed":0}
	{"Time":"2018-09-14T17:23:08.028337643+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Output":"PASS\n"}
	{"Time":"2018-09-14T17:23:08.028343022+02:00","Action":"output","Package":"github.com/Nimsaja/TaskOrganizer","Output":"ok  \tgithub.com/Nimsaja/TaskOrganizer\t(cached)\n"}
	{"Time":"2018-09-14T17:23:08.028351134+02:00","Action":"pass","Package":"github.com/Nimsaja/TaskOrganizer","Elapsed":0}		
`
)

func TestTaskOrganizer(t *testing.T) {
	r, err := Handle([]byte(jsonTestTaskOrganizer))
	if err != nil {
		t.Errorf("no err expected: %v", err)
	}

	s := New(r)
	if s.Tests != 8 {
		t.Errorf("Tests 8 != %v", s.Tests)
	}
	if s.Pass != 8 {
		t.Errorf("Pass: 8 != %v", s.Pass)
	}
	if s.Fail != 0 {
		t.Errorf("Fail: 0 != %v", s.Fail)
	}
	if s.Skip != 0 {
		t.Errorf("Skip: 0 != %v", s.Skip)
	}
	if len(s.EmptyPackages) != 0 {
		t.Errorf("EmptyPackages: 0 != %v", len(s.EmptyPackages))
	}
}
