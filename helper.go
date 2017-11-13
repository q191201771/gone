package main

import (
	"testing"
)

func assert(t *testing.T, a interface{}, b interface{}, msg string) {
	if a == b {
		return
	}
	t.Errorf("%s %v != %v", msg, a, b)
}
