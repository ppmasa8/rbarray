package rbarray

import "log"

type intArray []int
type strArray []string

type Array struct {
	intVals intArray
	strVals strArray
}

func (a *Array) Pop() interface{} {
	if len(a.intVals) > 0 {
		slice := a.intVals
		last := slice[len(slice)-1]
		a.intVals = slice[:len(slice)-1]
		return last
	}
	if len(a.strVals) > 0 {
		slice := a.strVals
		last := slice[len(slice)-1]
		a.strVals = slice[:len(slice)-1]
		return last
	}
	log.Println("Array is empty")
	return nil
}
