package main

import (
	"bytes"
	"errors"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
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
	data := []struct {
		name     string
		input    []byte
		expected *Result
	}{
		{
			"basic multiplication",
			[]byte("CALC_5\n*\n2\n5"),
			&Result{
				Id:    "CALC_5",
				Value: 10,
			},
		},
		{
			"basic summation",
			[]byte("CALC_6\n+\n1\n2"),
			&Result{
				Id:    "CALC_6",
				Value: 3,
			},
		},
		{
			"basic substraction",
			[]byte("CALC_7\n-\n3\n1"),
			&Result{
				Id:    "CALC_7",
				Value: 2,
			},
		},
		{
			"zero division",
			[]byte("CALC_8\n/\n5\n0"),
			nil,
		},
		{
			"bad input",
			[]byte("foobar"),
			nil,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			in := make(chan []byte)
			out := make(chan Result)

			go DataProcessor(in, out)

			in <- d.input
			close(in)

			result, ok := <-out
			if !ok && d.expected != nil {
				t.Fatal("output channel was closed unexpectedly")
			}

			if d.expected == nil {
				return
			} else if diff := cmp.Diff(*d.expected, result); diff != "" {
				t.Error(diff)
			}
		})
	}
}

type notifyingWriter struct {
	buf *bytes.Buffer
	wg  *sync.WaitGroup
}

func (w *notifyingWriter) Write(p []byte) (int, error) {
	defer w.wg.Done()
	return w.buf.Write(p)
}

func TestIntegration(t *testing.T) {
	ch1 := make(chan []byte)
	ch2 := make(chan Result)

	var buf bytes.Buffer
	var wg sync.WaitGroup
	wg.Add(1)

	writer := &notifyingWriter{buf: &buf, wg: &wg}

	go DataProcessor(ch1, ch2)
	go WriteData(ch2, writer)

	ts := httptest.NewServer(NewController(ch1))
	defer ts.Close()

	body := `CALC_6
/
10
5`

	resp, err := http.Post(ts.URL, "text/plain", strings.NewReader(body))
	if err != nil {
		t.Fatalf("http post failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("unexpected status: %d", resp.StatusCode)
	}

	wg.Wait()

	got := buf.String()
	expected := "CALC_6:2\n"

	if got != expected {
		t.Errorf("was expecting `%s` got `%s`", expected, got)
	}
}
