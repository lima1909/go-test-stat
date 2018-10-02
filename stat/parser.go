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
type filter func(e TestEvent) bool

func noFilter(_ TestEvent) bool {
	return false
}

// parse json bytes and unmarshal to TestEvents (fire func occured by every TestEvent)
func parse(b []byte, occ occurred) error {
	return parseF(b, occ, noFilter)
}

// parseF parse with Filter
func parseF(b []byte, occ occurred, fltr filter) error {
	jsonstr := strings.TrimSpace(string(b))
	scanner := bufio.NewScanner(strings.NewReader(jsonstr))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var event TestEvent
		err := json.Unmarshal(scanner.Bytes(), &event)
		if err != nil {
			return fmt.Errorf("err by parsing json: %v", err)
		}
		if !fltr(event) {
			occ(event)
		}
	}
	return nil
}
