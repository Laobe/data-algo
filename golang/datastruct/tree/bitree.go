package tree

import (
	"fmt"
)

type BiTNode struct {
	data           TElemType
	lchild, rchild *BiTNode
}
type BiTree *BiTNode

func CreateBiTree(T *BiTree) {
	var ch TElemType
	fmt.Scanf("%d", &ch)
	if ch == "#" {
		*T == nil
	} else {
		*T = malloc(sizeof(BiTree))
	}
	CreateBiTree(&(*T).lchild)
	CreateBiTree(&(*T).rchild)
}
