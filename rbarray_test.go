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

func TestArray_Push(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		obj      interface{}
		expected interface{}
	}{
		{
			name:     "Push int to IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      4,
			expected: Array{IntVals: IntArray{1, 2, 3, 4}},
		},
		{
			name:     "Push string to StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "qux",
			expected: Array{StrVals: StrArray{"foo", "bar", "baz", "qux"}},
		},
		{
			name:     "Push invalid type",
			array:    Array{},
			obj:      4.0,
			expected: Array{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.array.Push(test.obj)
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
		})
	}
}
