package tests

import (
	"testing"
)

func TestCheck1(t *testing.T) {

	if Check1(5, 5) != 10 {
		t.Error("Error during get web content.")
	}
}

func Check1(a int, b int) int {
	return a + b
}
