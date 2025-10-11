package languagebasics

import (
	"reflect"
	"testing"
)

func TestFlattenSlice(t *testing.T) {
	tests := []struct {
		input  []interface{}
		output []int
	}{
		{[]interface{}{[]interface{}{1, 2}, []interface{}{3, 4}}, []int{1, 2, 3, 4}},
		{[]interface{}{"ram"}, []int{}}, // non-int ignored
		{[]interface{}{}, []int{}},      // empty slice
	}

	for _, tt := range tests {
		got := flattenSlice(tt.input)
		if !reflect.DeepEqual(got, tt.output) {
			t.Errorf("flattenSlice(%v) = %v; want %v", tt.input, got, tt.output)
		}
	}
}
