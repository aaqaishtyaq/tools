package utils

import (
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	t.Run("returns unique string slice from a string slice", func(t *testing.T) {
		data := []string{
			"branch-1",
			"branch-2",
			"branch-2",
			"branch-3",
			"branch-4",
			"branch-1",
			"branch-4",
			"branch-3",
		}

		want := []string{"branch-1", "branch-2", "branch-3", "branch-4"}
		got := UniqueStrings(data)

		if len(got) != len(want) {
			t.Errorf("got count %q, want count %q", len(got), len(want))
		}
	})
}

func TestGenerateRand(t *testing.T) {
	t.Run("returns unique string slice from a string slice", func(t *testing.T) {
		want := 10
		got := GenerateRand(want)

		if len(got) != want {
			t.Errorf("got count %q, want count %q", len(got), want)
		}
	})
}
