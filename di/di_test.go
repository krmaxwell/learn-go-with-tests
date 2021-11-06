package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kyle")

	got := buffer.String()
	want := "Hello, Kyle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
