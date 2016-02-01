package chain

import (
	"fmt"
)

const MAXSIZE = 5

type SElemType int
type SqStack struct {
	data [MAXSIZE]SElemType
	top  int
}

func NewSqStack(selem [MAXSIZE]SElemType, top int) SqStack {
	s := SqStack{selem, top}
	return s
}

func (S *SqStack) Push(e SElemType) bool {
	if S.top == MAXSIZE-1 {
		return false
	}
	S.top++
	S.data[S.top] = e
	return true
}

func (S *SqStack) Pop(e SElemType) bool {
	if S.top == -1 {
		return false
	}
	S.data[S.top] = e
	S.top--
	return true

}

func main() {
	fmt.Printf("%v\n", MAXSIZE)
	tmp := [MAXSIZE]SElemType{1, 2, 3}
	selem := tmp[:3]

	a := SqStack{tmp, 5}
	fmt.Printf("%v\n", a)
	// a.data = selem
	// a.top = cap(selem)

	fmt.Println(cap(selem))
	// fmt.Printf("%v\n", selem)
	// var b chain.SElemType
	// fmt.Printf("%v\n", a)
	// a.Pop(b)
	// fmt.Printf("%v\n", a)
}
