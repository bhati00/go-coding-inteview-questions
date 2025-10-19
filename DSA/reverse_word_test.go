package languagebasics

import (
	"testing"
)

func TestReverseWord(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello", "olleh"},
		{"GoLang", "gnaLoG"},
		{"世界", "界世"},
		{"", ""},
	}
	for _, tt := range tests {
		got := ReverseString(tt.input)
		if got != tt.output {
			t.Errorf("ReverseWords(%q) = %q; want %q", tt.input, got, tt.output)
		}
	}
}
