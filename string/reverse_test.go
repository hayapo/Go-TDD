package string

import (
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("reverse 'dowango' to 'ognawod'", func(t *testing.T) {
		got := StringReverse("dowango")
		want := "ognawod"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}