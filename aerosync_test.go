package main

import "testing"

func TestFoo(t *testing.T) {
	const foo, bar = 1, 4
	temp := foo + bar
	if temp != 3 {
		t.Errorf("Expected 3, actual %d", temp)
	}
}
