package image

type EdgeNode struct {
	adjvex int
	weigth EdgeType
	next   *EdgeNode
}

type VertexNode struct {
	data      VertexType
	firstedge *EdgeNode
}

type GraphAdjList struct {
	adjList               AdjList
	numVertexes, numEdges int
}

func CreateAlGraph(G *GraphAdjList) {

}
