package string

import "testing"

func TestJoin(t *testing.T) {
	stringSlice := [] string{"Hello", "World"}
	joined := Join(stringSlice, ", ")
	expected := "Hello, World"

	if joined != expected {
		t.Errorf("expected %q, but got %q", expected, joined)
	}
}