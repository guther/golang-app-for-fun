package tests

import (
	"testing"
)

func TestCheck2(t *testing.T) {

	if Check2(5, 5) != 10 {
		t.Error("Error during get web content.")
	}
}

func Check2(a int, b int) int {
	return a + b
}
