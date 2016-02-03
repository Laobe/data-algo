package main

import (
	. "data-algro/datastruct/chain"
	"fmt"
)

func main() {
	use_SqStack_pop_push()
	use_SqDoubleStack_pop_push()
	use_String_index()
}

func use_SqStack_pop_push() {
	fmt.Printf("%v\n", MAXSIZE)
	tmp := [MAXSIZE]SElemType{1, 2, 3}

	a := NewSqStack(tmp, 2)

	var b SElemType = 4
	fmt.Printf("%v\n", a)
	a.Push(b)
	fmt.Printf("%v\n", a)
	var c SElemType
	a.Pop(c)
	fmt.Printf("%v\n", a)
}

func use_SqDoubleStack_pop_push() {
	fmt.Printf("%v\n", MAXSIZE)
	tmp := [MAXSIZE]SElemType{1, 2, 0, 0, 1}

	a := NewSqDoubleStack(tmp, 1, 4)

	var b SElemType = 4
	fmt.Printf("%v\n", a)
	a.Push(b, 1)
	fmt.Printf("%v\n", a)
	var c SElemType
	a.Pop(c, 4)
	fmt.Printf("%v\n", a)
}

func use_String_index() {
	var a String = "abcccabccc"
	var b String = "cc"
	z := a.Index(b, 2)
	fmt.Println(z)
}

func use_Queue_EnQueue_DeQueue() {
	tmp := [MAXSIZE]QElemType{1, 2, 3}
	fmt.Printf("%v\n", tmp)
}
