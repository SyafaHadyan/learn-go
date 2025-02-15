package main

import "testing"

func TestHello(t *testing.T) {
	var got string = Hello()
	var want string = "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}