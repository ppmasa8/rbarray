package rbarray

import (
	"errors"
	"reflect"
	"testing"
)

func TestArray_Pop(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		obj      interface{}
		expected interface{}
	}{
		{
			name:     "Pop int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      3,
			expected: Array{IntVals: IntArray{1, 2}},
		},
		{
			name:     "Pop string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "baz",
			expected: Array{StrVals: StrArray{"foo", "bar"}},
		},
		{
			name:     "Empty array",
			array:    Array{},
			obj:      nil,
			expected: Array{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := test.array.Pop()
			// if err != nil {
			// 	t.Errorf("%v", err)
			// }
			if !reflect.DeepEqual(got, test.obj) {
				t.Errorf("Expected %v but got %v", test.obj, got)
			}
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
		})
	}
}

func TestArray_Shift(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		obj      interface{}
		expected interface{}
	}{
		{
			name:     "Shift int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      1,
			expected: Array{IntVals: IntArray{2, 3}},
		},
		{
			name:     "Shift string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "foo",
			expected: Array{StrVals: StrArray{"bar", "baz"}},
		},
		{
			name:     "Empty array",
			array:    Array{},
			obj:      nil,
			expected: Array{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := test.array.Shift()
			// if err != nil {
			// 	t.Errorf("%v", err)
			// }
			if !reflect.DeepEqual(got, test.obj) {
				t.Errorf("Expected %v but got %v", test.obj, got)
			}
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
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

func TestArray_Unshift(t *testing.T) {
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

func TestArray_Delete(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		obj      interface{}
		expected interface{}
	}{
		{
			name:     "Delete int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      2,
			expected: Array{IntVals: IntArray{1, 3}},
		},
		{
			name:     "Delete string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "bar",
			expected: Array{StrVals: StrArray{"foo", "baz"}},
		},
		{
			name:     "Delete invalid type",
			array:    Array{},
			obj:      4.0,
			expected: Array{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.array.Delete(test.obj)
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
		})
	}
}

func TestArray_Uniq(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected Array
	}{
		{
			name:     "Uniq int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: Array{IntVals: IntArray{1, 2, 3, 4, 5}},
		},
		{
			name:     "Uniq string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: Array{StrVals: StrArray{"foo", "bar", "baz", "qux"}},
		},
		{
			name:     "Uniq invalid type",
			array:    Array{},
			expected: Array{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.array.Uniq()
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
		})
	}
}

func TestArray_Sum(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected int
	}{
		{
			name:     "Sum int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 23,
		},
		{
			name:     "Sum empty int from IntVals",
			array:    Array{},
			expected: 0,
		},
		{
			name:     "Sum string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := test.array.Sum()
			if sum != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, sum)
			}
		})
	}
}

func TestArray_Max(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected int
		err error
	}{
		{
			name:     "Max int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 5,
			err: nil,
		},
		{
			name:     "Max empty int from IntVals",
			array:    Array{},
			expected: 0,
			err: errors.New("IntArray is empty"),
		},
		{
			name:     "Max string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: 0,
			err: errors.New("IntArray is empty"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			max, err := test.array.Max()
			if max != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, max)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
			}
		})
	}
}

func TestArray_Min(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected int
	}{
		{
			name:     "Min int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 1,
		},
		{
			name:     "Min empty int from IntVals",
			array:    Array{},
			expected: 0,
		},
		{
			name:     "Min string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			min, _ := test.array.Min()
			if min != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, min)
			}
		})
	}
}
