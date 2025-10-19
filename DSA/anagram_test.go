package languagebasics

import (
	"slices"
	"testing"
)

func TestAnagramIndices(t *testing.T) {
	tests := []struct {
		a      string
		b      string
		output []int
	}{
		{"", "", []int{}},
		{"aaa", "", []int{}},
		{"aaa", "a", []int{0, 1, 2}},
	}
	for _, tt := range tests {
		result := AnagramIndices(tt.a, tt.b)
		if !slices.Equal(result, tt.output) {
			t.Errorf("test case with string (%q) and (%q) failed. ", tt.a, tt.b)
		}
	}
}
