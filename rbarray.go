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
// pop() -> object | nil
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
// shift() -> object | nil
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
func (a *Array) Delete(val interface{}) (interface{}, error) {
	switch v := val.(type) {
	case int:
		for i, n := range a.IntVals {
			if n == v {
				a.IntVals = append(a.IntVals[:i], a.IntVals[i+1:]...)
				return n, nil
			}
		}
	case string:
		for i, s := range a.StrVals {
			if s == v {
				a.StrVals = append(a.StrVals[:i], a.StrVals[i+1:]...)
				return s, nil
			}
		}
	}
	return nil, fmt.Errorf("Value not found")
}

// instance method Array#uniq
// uniq() -> Array
func (a *Array) Uniq() {
	seen := make(map[interface{}]bool)
	var uniqInts IntArray
	var uniqStrs StrArray
	for _, n := range a.IntVals {
		if _, ok := seen[n]; !ok {
			seen[n] = true
			uniqInts = append(uniqInts, n)
		}
	}
	for _, s := range a.StrVals {
		if _, ok := seen[s]; !ok {
			seen[s] = true
			uniqStrs = append(uniqStrs, s)
		}
	}
	a.IntVals = uniqInts
	a.StrVals = uniqStrs
}

// Only works for IntArray
// instance method Enumerable#sum
// sum() -> object
func (a *Array) Sum() int {
	var sum int
	for _, n := range a.IntVals {
		sum += n
	}
	return sum
}

// Only works for IntArray
// instance method Enumerable#max
// max() -> object
func (a *Array) Max() int {
	max := a.IntVals[0]
	for _, n := range a.IntVals {
		if n > max {
			max = n
		}
	}
	return max
}

// Only works for IntArray
// instance method Enumerable#min
// min() -> object
func (a *Array) Min() int {
	min := a.IntVals[0]
	for _, n := range a.IntVals {
		if n < min {
			min = n
		}
	}
	return min
}
