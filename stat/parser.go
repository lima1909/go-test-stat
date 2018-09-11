package stat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// TestEvent is the JSON struct which is occurred
// on Test has many Events
type TestEvent struct {
	Time    time.Time `json:",omitempty"`
	Action  string
	Package string  `json:",omitempty"`
	Test    string  `json:",omitempty"`
	Elapsed float64 `json:",omitempty"`
	Output  string  `json:",omitempty"`
}

type occurred func(e TestEvent)

// Parse json bytes and unmarshal to TestEvents (fire func occured by every TestEvent)
func Parse(b []byte, occ occurred) error {
	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var event TestEvent
		err := json.Unmarshal(scanner.Bytes(), &event)
		if err != nil {
			return fmt.Errorf("err by parsing json: %v", err)
		}
		occ(event)
	}
	return nil
}
