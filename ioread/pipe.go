package ioread

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Pipe read from stdin (pipe) and convert it to string
func Pipe() ([]byte, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, fmt.Errorf("err by read from stdin-stat: %v", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return nil, fmt.Errorf("no value per pipe")
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("can not read from stdin: %v", err)
	}
	return b, nil
}
