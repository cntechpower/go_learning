package main

import (
	"flag"
	"sync"
	"testing"
)

var a string
var b = flag.String("b", "b", "")
var once sync.Once

func init() {
	flag.StringVar(&a, "a", "default", "")
}

func parseFlag() {
	once.Do(flag.Parse)
}

func TestA(t *testing.T) {
	parseFlag()
	t.Log(a)
	t.Log(*b)
}

func TestB(t *testing.T) {
	parseFlag()
	t.Log(a)
	t.Log(*b)
}

// go test -v unit_test.go -args -a test -b test2
