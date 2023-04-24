package rbarray

import (
	"reflect"
	"testing"
)

func TestArray_Pop(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected interface{}
	}{
		{
			name:     "Pop int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			expected: 3,
		},
		{
			name:     "Pop string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			expected: "baz",
		},
		{
			name:     "Empty array",
			array:    Array{},
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.array.Pop()
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestArray_Shift(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected interface{}
	}{
		{
			name:     "Shift int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			expected: 1,
		},
		{
			name:     "Shift string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			expected: "foo",
		},
		{
			name:     "Empty array",
			array:    Array{},
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.array.Shift()
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}
