package main

import (
	. "data-algro/datastruct/chain"
	"fmt"
)

func main() {
	fmt.Printf("%v\n", MAXSIZE)
	tmp := [MAXSIZE]SElemType{1, 2, 3}
	// selem := tmp[:3]

	a := NewSqStack(tmp, 2)
	// a.data = selem
	// a.top = cap(selem)

	// fmt.Printf("%v\n", selem)
	var b SElemType = 4
	fmt.Printf("%v\n", a)
	a.Push(b)
	fmt.Printf("%v\n", a)
	var c SElemType
	a.Pop(c)
	fmt.Printf("%v\n", a)
}
