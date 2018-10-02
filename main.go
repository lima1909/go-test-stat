package main

import (
	"github.com/lima1909/go-test-stat/ioread"
	"github.com/lima1909/go-test-stat/iowrite"
	"github.com/lima1909/go-test-stat/stat"
)

func main() {
	b, err := ioread.Pipe()
	if err != nil {
		panic(err)
	}
	r, err := stat.Handle(b)
	if err != nil {
		panic(err)
	}
	s := stat.New(r)
	iowrite.Print(s)
}
