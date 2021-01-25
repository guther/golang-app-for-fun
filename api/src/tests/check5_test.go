package tests

import (
	"testing"
)

func TestCheck5(t *testing.T) {

	if Check5(5, 5) != 10 {
		t.Error("Error during get web content.")
	}
}

func Check5(a int, b int) int {
	return a + b
}
