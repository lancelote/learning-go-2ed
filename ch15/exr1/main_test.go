package main

import (
	"errors"
	"github.com/google/go-cmp/cmp"
	"strconv"
	"testing"
)

func TestParser(t *testing.T) {
	data := []struct {
		name     string
		input    []byte
		expected Input
		err      error
	}{
		{
			"basic summation",
			[]byte("CALC_1\n+\n3\n2"),
			Input{Id: "CALC_1", Op: "+", Val1: 3, Val2: 2},
			nil,
		},
		{
			"basic multiplication",
			[]byte("CALC_2\n*\n100\n3000"),
			Input{Id: "CALC_2", Op: "*", Val1: 100, Val2: 3000},
			nil,
		},
		{
			"invalid first value",
			[]byte("CALC_3\n+\na\n3"),
			Input{},
			strconv.ErrSyntax,
		},
		{
			"invalid second value",
			[]byte("CALC_4\n+\n1\nb"),
			Input{},
			strconv.ErrSyntax,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := parser(d.input)
			if diff := cmp.Diff(d.expected, result); diff != "" {
				t.Error(diff)
			}
			if !errors.Is(err, d.err) {
				t.Errorf("expected conversion syntax error, got `%v`", err)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	in := make(chan []byte)
	out := make(chan Result)

	go DataProcessor(in, out)

	data := []byte("CALC_5\n*\n2\n5")
	in <- data
	close(in)

	result, ok := <-out
	if !ok {
		t.Fatal("output channel was closed unexpectedly")
	}

	expected := Result{
		Id:    "CALC_5",
		Value: 10,
	}

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
