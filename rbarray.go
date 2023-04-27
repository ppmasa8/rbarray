package rbarray

import (
	"fmt"
)

type IntArray []int
type StrArray []string

type Array struct {
	IntVals IntArray
	StrVals StrArray
}

// instance method Array#pop
// pop -> object | nil
func (a *Array) Pop() (interface{}, error) {
	if len(a.IntVals) > 0 {
		slice := a.IntVals
		last := slice[len(slice)-1]
		a.IntVals = slice[:len(slice)-1]
		return last, nil
	}
	if len(a.StrVals) > 0 {
		slice := a.StrVals
		last := slice[len(slice)-1]
		a.StrVals = slice[:len(slice)-1]
		return last, nil
	}
	return nil, fmt.Errorf("Array is empty")
}

// instance method Array#shift
// shift -> object | nil
func (a *Array) Shift() (interface{}, error) {
	if len(a.IntVals) > 0 {
		slice := a.IntVals
		first := slice[0]
		a.IntVals = slice[1:]
		return first, nil
	}
	if len(a.StrVals) > 0 {
		slice := a.StrVals
		first := slice[0]
		a.StrVals = slice[1:]
		return first, nil
	}
	return nil, fmt.Errorf("Array is empty")
}

// instance method Array#push
// push(*obj) -> self
func (a *Array) Push(obj interface{}) error {
	switch v := obj.(type) {
	case int:
		a.IntVals = append(a.IntVals, v)
	case string:
		a.StrVals = append(a.StrVals, v)
	default:
		return fmt.Errorf("Invalid type: %T", obj)
	}
	return nil
}

// instance method Array#unshift
// unshift(*obj) -> self
func (a *Array) Unshift(obj interface{}) error {
	switch v := obj.(type) {
	case int:
		a.IntVals = append(IntArray{v}, a.IntVals...)
	case string:
		a.StrVals = append(StrArray{v}, a.StrVals...)
	default:
		return fmt.Errorf("Invalid type: %T", obj)
	}
	return nil
}

// instance method Array#delete
// delete(val) { ... } -> object
func (a *Array) Delete(val interface{}) interface{} {
	switch v := val.(type) {
	case int:
		for i, n := range a.IntVals {
			if n == v {
				a.IntVals = append(a.IntVals[:i], a.IntVals[i+1:]...)
				return n
			}
		}
	case string:
		for i, s := range a.StrVals {
			if s == v {
				a.StrVals = append(a.StrVals[:i], a.StrVals[i+1:]...)
				return s
			}
		}
	}
	return nil
}
