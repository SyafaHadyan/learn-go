package main

import "testing"

func TestHello(t *testing.T) {
	var got string = Hello("Name")
	var want string = "Hello, Name"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}