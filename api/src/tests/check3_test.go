package tests

import (
	"testing"
)

func TestCheck3(t *testing.T) {

	if Check3(5, 5) != 10 {
		t.Error("Error during get web content.")
	}
}

func Check3(a int, b int) int {
	return a + b
}
