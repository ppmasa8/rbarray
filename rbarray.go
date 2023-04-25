package rbarray

import "log"

type IntArray []int
type StrArray []string

type Array struct {
	IntVals IntArray
	StrVals StrArray
}

// instance method Array#pop
// pop -> object | nil
func (a *Array) Pop() interface{} {
	if len(a.IntVals) > 0 {
		slice := a.IntVals
		last := slice[len(slice)-1]
		a.IntVals = slice[:len(slice)-1]
		return last
	}
	if len(a.StrVals) > 0 {
		slice := a.StrVals
		last := slice[len(slice)-1]
		a.StrVals = slice[:len(slice)-1]
		return last
	}
	log.Println("Array is empty")
	return nil
}

// instance method Array#shift
// shift -> object | nil
func (a *Array) Shift() interface{} {
	if len(a.IntVals) > 0 {
		slice := a.IntVals
		first := slice[0]
		a.IntVals = slice[1:]
		return first
	}
	if len(a.StrVals) > 0 {
		slice := a.StrVals
		first := slice[0]
		a.StrVals = slice[1:]
		return first
	}
	log.Println("Array is empty")
	return nil
}

// instance method Array#push
// push(*obj) -> self
func (a *Array) Push(obj interface{}) {
	switch obj.(type) {
	case int:
		a.IntVals = append(a.IntVals, obj.(int))
	case string:
		a.StrVals = append(a.StrVals, obj.(string))
	default:
		log.Println("Invalid type")
	}
}
