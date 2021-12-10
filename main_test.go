package main

import (
	"testing"
)

func TestHello(t *testing.T){
	expected:="Hello World"
	current:=Hello()

	if expected !=current {
		t.Fatalf("Wanted %s got %s", expected, current)

	}
}
