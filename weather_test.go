package main

import "testing"

func MyFunction(i int) int {
	return i * 2;
}

func TestMyFunction(t *testing.T) {
    result := MyFunction(5)
    if result != 10 {
        t.Errorf("MyFunction(5) returned %d, but expected 10", result)
    }
}
