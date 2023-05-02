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
		err      error
	}{
		{
			name:     "Pop int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      3,
			expected: Array{IntVals: IntArray{1, 2}},
			err:      nil,
		},
		{
			name:     "Pop string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "baz",
			expected: Array{StrVals: StrArray{"foo", "bar"}},
			err:      nil,
		},
		{
			name:     "Empty array",
			array:    Array{},
			obj:      nil,
			expected: Array{},
			err:      errors.New("Array is empty"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.array.Pop()
			if !reflect.DeepEqual(got, test.obj) {
				t.Errorf("Expected %v but got %v", test.obj, got)
			}
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
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
		err      error
	}{
		{
			name:     "Shift int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      1,
			expected: Array{IntVals: IntArray{2, 3}},
			err:      nil,
		},
		{
			name:     "Shift string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "foo",
			expected: Array{StrVals: StrArray{"bar", "baz"}},
			err:      nil,
		},
		{
			name:     "Empty array",
			array:    Array{},
			obj:      nil,
			expected: Array{},
			err:      errors.New("Array is empty"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.array.Shift()
			if !reflect.DeepEqual(got, test.obj) {
				t.Errorf("Expected %v but got %v", test.obj, got)
			}
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
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
		err      error
	}{
		{
			name:     "Push int to IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      4,
			expected: Array{IntVals: IntArray{1, 2, 3, 4}},
			err:      nil,
		},
		{
			name:     "Push string to StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "qux",
			expected: Array{StrVals: StrArray{"foo", "bar", "baz", "qux"}},
			err:      nil,
		},
		{
			name:     "Push invalid type",
			array:    Array{},
			obj:      4.0,
			expected: Array{},
			err:      errors.New("invalid type: float64"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.array.Push(test.obj)
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, got)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
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
		err      error
	}{
		{
			name:     "Unshift int to IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      4,
			expected: Array{IntVals: IntArray{4, 1, 2, 3}},
			err:      nil,
		},
		{
			name:     "Unshift string to StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "qux",
			expected: Array{StrVals: StrArray{"qux", "foo", "bar", "baz"}},
			err:      nil,
		},
		{
			name:     "Unshift invalid type",
			array:    Array{},
			obj:      4.0,
			expected: Array{},
			err:      errors.New("invalid type: float64"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.array.Unshift(test.obj)
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, got)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
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
		err      error
	}{
		{
			name:     "Delete int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3}},
			obj:      2,
			expected: Array{IntVals: IntArray{1, 3}},
			err:      nil,
		},
		{
			name:     "Delete string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz"}},
			obj:      "bar",
			expected: Array{StrVals: StrArray{"foo", "baz"}},
			err:      nil,
		},
		{
			name:     "Delete invalid type",
			array:    Array{},
			obj:      4.0,
			expected: Array{},
			err:      errors.New("invalid type: float64"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.array.Delete(test.obj)
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
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
			got := test.array.Uniq()
			if !reflect.DeepEqual(test.array, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, test.array)
			}
			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("Expected %v but got %v", test.expected, got)
			}
		})
	}
}

func TestArray_Sum(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected interface{}
		err      error
	}{
		{
			name:     "Sum int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 23,
			err:      nil,
		},
		{
			name:     "Sum empty int from IntVals",
			array:    Array{},
			expected: nil,
			err:      errors.New("cannot sum empty IntArray"),
		},
		{
			name:     "Sum string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: nil,
			err:      errors.New("cannot sum empty IntArray"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum, err := test.array.Sum()
			if sum != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, sum)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
			}
		})
	}
}

func TestArray_Max(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		expected interface{}
		err      error
	}{
		{
			name:     "Max int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 5,
			err:      nil,
		},
		{
			name:     "Max empty int from IntVals",
			array:    Array{},
			expected: nil,
			err:      errors.New("cannot max empty IntArray"),
		},
		{
			name:     "Max string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: nil,
			err:      errors.New("cannot max empty IntArray"),
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
		expected interface{}
		err      error
	}{
		{
			name:     "Min int from IntVals",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expected: 1,
			err:      nil,
		},
		{
			name:     "Min empty int from IntVals",
			array:    Array{},
			expected: nil,
			err:      errors.New("cannot min empty IntArray"),
		},
		{
			name:     "Min string from StrVals",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expected: nil,
			err:      errors.New("cannot min empty IntArray"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			min, err := test.array.Min()
			if min != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, min)
			}
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected %v but got %v", test.err, err)
			}
		})
	}
}

func TestArray_Size(t *testing.T) {
	tests := []struct {
		name            string
		array           Array
		expectedIntSize int
		expectedStrSize int
	}{
		{
			name:            "IntVals with duplicates",
			array:           Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			expectedIntSize: 7,
			expectedStrSize: 0,
		},
		{
			name:            "Empty Array",
			array:           Array{},
			expectedIntSize: 0,
			expectedStrSize: 0,
		},
		{
			name:            "StrVals with duplicates",
			array:           Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			expectedIntSize: 0,
			expectedStrSize: 6,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			intSize, strSize := test.array.Size()
			if intSize != test.expectedIntSize {
				t.Errorf("For test %q, expected int size %d but got %d", test.name, test.expectedIntSize, intSize)
			}
			if strSize != test.expectedStrSize {
				t.Errorf("For test %q, expected str size %d but got %d", test.name, test.expectedStrSize, strSize)
			}
		})
	}
}

func TestArray_Include(t *testing.T) {
	tests := []struct {
		name     string
		array    Array
		obj      interface{}
		expected bool
	}{
		{
			name:     "IntVals with duplicates",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			obj:      3,
			expected: true,
		},
		{
			name:     "IntVals with duplicates",
			array:    Array{IntVals: IntArray{1, 2, 3, 3, 4, 5, 5}},
			obj:      6,
			expected: false,
		},
		{
			name:     "StrVals with duplicates",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			obj:      "baz",
			expected: true,
		},
		{
			name:     "StrVals with duplicates",
			array:    Array{StrVals: StrArray{"foo", "bar", "baz", "baz", "qux", "qux"}},
			obj:      "quux",
			expected: false,
		},
		{
			name:     "Empty Array",
			array:    Array{},
			obj:      "quux",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.array.Include(test.obj) != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, test.array.Include(test.obj))
			}
		})
	}
}
