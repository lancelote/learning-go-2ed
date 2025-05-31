package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParser(t *testing.T) {
	expected := Input{Id: "CALC_1", Op: "+", Val1: 3, Val2: 2}
	result, _ := parser([]byte("CALC_1\n+\n3\n2"))

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
