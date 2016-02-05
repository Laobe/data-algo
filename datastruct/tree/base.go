package tree

const MAX_TREE_SIZE = 100

type TElemType int

type PTNode struct {
	data   TElemType
	parent int
}

type PTree struct {
	nodes [MAX_TREE_SIZE]PTNode
	r, n  int
}

type CTNode struct {
	child int
	next  *CTNode
}

type ChildPtr *CTNode

type CTBox struct {
	data       TElemType
	firstchild ChildPtr
}

type CTree struct {
	nodes [MAX_TREE_SIZE]CTNode
	r, n  int
}

type CSNode struct {
	data                 TElemType
	firstchild, rightsib *CSNode
}

type CSTree *CSNode
