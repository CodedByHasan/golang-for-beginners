package main

import "testing"
import "example.com/test/process"

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "true"
	res := process.CheckEven(i)

	if expected != res {
		t.Errorf("expected: %v, got:%v", expected, res)
	}
}
