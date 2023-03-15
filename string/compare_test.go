package string

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	assertComparedResult := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
	t.Run("compare a < b lexicographically", func(t *testing.T) {
		got := Compare("a", "b")
		want := -1
		assertComparedResult(t, got, want)
	})
	t.Run("compare a > b lexicographically", func(t *testing.T) {
		got := Compare("b", "a")
		want := 1
		assertComparedResult(t, got, want)
	})
	t.Run("compare a == a lexicographically", func(t *testing.T) {
		got := Compare("a", "a")
		want := 0
		assertComparedResult(t, got, want)
	})
}

func ExampleCompare() {
	compared := Compare("a", "b")
	fmt.Println(compared)
	// Output: -1
}