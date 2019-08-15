package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	err := Greet("Hello World")

	if err != nil {
		t.Error("Greet failed", err)
	}
}
