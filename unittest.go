package main

import "testing"

func TestAdd(t *testing.T) {
	correct := 10

	r := 7 + 3

	if r != correct {
		t.Errorf("Got %d expected %d", r, correct)
	}
}
