package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

type printableInt int

func (p printableInt) String() string {
	return fmt.Sprintf("printable-int:%d", int(p))
}

type printableFloat float64

func (p printableFloat) String() string {
	return fmt.Sprintf("printable-float:%.2f", float64(p))
}

func TestDouble(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		got  any
		want any
	}{
		{
			name: "int",
			got:  double(21),
			want: 42,
		},
		{
			name: "negative int",
			got:  double(-7),
			want: -14,
		},
		{
			name: "float64",
			got:  double(2.5),
			want: 5.0,
		},
		{
			name: "custom int type",
			got:  double(printableInt(11)),
			want: printableInt(22),
		},
		{
			name: "custom float64 type",
			got:  double(printableFloat(1.25)),
			want: printableFloat(2.5),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.got != tc.want {
				t.Fatalf("got %v, want %v", tc.got, tc.want)
			}
		})
	}
}

func TestPrintableTypesImplementConstraint(t *testing.T) {
	t.Parallel()

	assertPrintable(printableInt(1))
	assertPrintable(printableFloat(1.5))
}

func TestPrintNum(t *testing.T) {
	output := captureStdout(t, func() {
		printNum(printableInt(15))
		printNum(printableFloat(3.5))
	})

	want := "printable-int:15\nprintable-float:3.50\n"
	if output != want {
		t.Fatalf("got %q, want %q", output, want)
	}
}

func TestLinkedListAdd(t *testing.T) {
	element := &Element[int]{val: 1}

	list := LinkedList[int]{
		head: element,
		last: element,
	}

	list.Add(2)
	list.Add(3)

	if got := list.Index(1); got != 0 {
		t.Fatalf("Index(1) = %d, want 0", got)
	}
	if got := list.Index(2); got != 1 {
		t.Fatalf("Index(2) = %d, want 1", got)
	}
	if got := list.Index(3); got != 2 {
		t.Fatalf("Index(3) = %d, want 2", got)
	}
}

func TestLinkedListInsert(t *testing.T) {
	node3 := &Element[int]{val: 3}
	node2 := &Element[int]{val: 2, next: node3}
	head := &Element[int]{val: 1, next: node2}

	list := LinkedList[int]{
		head: head,
		last: node3,
	}

	list.Insert(99, 1)

	if got := list.Index(1); got != 0 {
		t.Fatalf("Index(1) = %d, want 0", got)
	}
	if got := list.Index(99); got != 1 {
		t.Fatalf("Index(99) = %d, want 1", got)
	}
	if got := list.Index(2); got != 2 {
		t.Fatalf("Index(2) = %d, want 2", got)
	}
	if got := list.Index(3); got != 3 {
		t.Fatalf("Index(3) = %d, want 3", got)
	}
}

func TestLinkedListIndexMissing(t *testing.T) {
	node3 := &Element[int]{val: 3}
	node2 := &Element[int]{val: 2, next: node3}
	head := &Element[int]{val: 1, next: node2}

	list := LinkedList[int]{
		head: head,
		last: node3,
	}

	if got := list.Index(42); got != -1 {
		t.Fatalf("Index(42) = %d, want -1", got)
	}
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	originalStdout := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}

	os.Stdout = writer

	outputCh := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, reader)
		outputCh <- buf.String()
	}()

	fn()

	_ = writer.Close()
	os.Stdout = originalStdout

	return <-outputCh
}

func assertPrintable[T Printable](T) {}
