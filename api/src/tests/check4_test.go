package tests

import (
	"testing"
)

func TestCheck4(t *testing.T) {

	if Check4(5, 5) != 10 {
		t.Error("Error during get web content.")
	}
}

func Check4(a int, b int) int {
	return a + b
}
