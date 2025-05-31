package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestParser(t *testing.T) {
	data := []struct {
		name     string
		input    []byte
		expected Input
	}{
		{
			"basic summation",
			[]byte("CALC_1\n+\n3\n2"),
			Input{Id: "CALC_1", Op: "+", Val1: 3, Val2: 2},
		},
		{
			"basic multiplication",
			[]byte("CALC_2\n*\n100\n3000"),
			Input{Id: "CALC_2", Op: "*", Val1: 100, Val2: 3000},
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, _ := parser(d.input)
			if diff := cmp.Diff(d.expected, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}
