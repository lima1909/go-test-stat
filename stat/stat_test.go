package stat

import (
	"os"
	"testing"
	"time"
)

// for test of the Elapsed works
func TestSleep(t *testing.T) {
	time.Sleep(time.Millisecond * 123)
}

func TestFail(t *testing.T) {
	if len(os.Getenv("TESTFAIL")) > 0 {
		t.Errorf("Fail Example")
	}
}
