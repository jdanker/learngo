package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jahred")
	want := "Hello, Jahred"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
