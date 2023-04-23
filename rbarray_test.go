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
            name:     "Pop int from intVals",
            array:    Array{intVals: intArray{1, 2, 3}},
            expected: 3,
        },
        {
            name:     "Pop string from strVals",
            array:    Array{strVals: strArray{"foo", "bar", "baz"}},
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
