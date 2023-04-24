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
