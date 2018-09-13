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

func TestPercent(t *testing.T) {
	o := New()
	o.Tests = 10
	r := o.Percent(2)
	if r != 2 {
		t.Errorf("2 != %v", r)
	}
}

func TestPercentWithZeroTests(t *testing.T) {
	o := New()
	o.Tests = 0
	r := o.Percent(2)
	if r != 0 {
		t.Errorf("0 != %v", r)
	}
}
