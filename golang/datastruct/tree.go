package main

import (
	. "data-algro/datastruct/tree"
	// "fmt"
)

func main() {
}

func use_BiTree() {
	var data TElemType
	var lchild, rchild *BiTNode
	T := *BiTree{data, lchild, rchild}
	CreateBiTree(T)
}
