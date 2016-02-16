package main

import (
	. "data-algro/datastruct/image"
	// "fmt"
)

func main() {
	// use_SqStack_pop_push()
	// use_SqDoubleStack_pop_push()
	// use_String_index()
}

func use_MGraph() {
	var maxvex = 10
	vexs := make([]VertexType, maxvex)
	src := make([][]EdgeType, maxvex)
	G := MGraph{vexs, src, maxvex, maxvex}
	CreateMGraph(*G)
	// fmt.Printf("%v\n", MAXSIZE)
	// tmp := [MAXSIZE]SElemType{1, 2, 3}

	// a := NewSqStack(tmp, 2)

	// var b SElemType = 4
	// fmt.Printf("%v\n", a)
	// a.Push(b)
	// fmt.Printf("%v\n", a)
	// var c SElemType
	// a.Pop(c)
	// fmt.Printf("%v\n", a)
}
