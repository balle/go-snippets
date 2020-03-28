package main

import "testing"

func TestCompare(t *testing.T) {
	result := compare(1, 1)

	if result != true {
		t.Fatal("oh nooo")
	}
}
