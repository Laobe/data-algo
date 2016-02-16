package image

import (
	"fmt"
)

type MGraph struct {
	vexs                  [MAXVEX]VertexType
	arc                   [MAXVEX][MAXVEX]EdgeType
	numVertexes, numEdges int
}

func CreateMGraph(G *MGraph) {
	// var i, j, k, w int
	var i, j int
	var edge, w EdgeType
	fmt.Printf("输入顶点数和边数:\n")
	fmt.Scanf("%d,%d", &G.numVertexes, &G.numEdges)
	for i := 0; i < G.numVertexes; i++ {
		fmt.Scanf("%s", &G.vexs)
	}
	for i := 0; i < G.numVertexes; i++ {
		for j := 0; j < G.numEdges; j++ {
			G.arc[i][j] = edge
		}
	}
	for k := 0; k < G.numEdges; k++ {
		fmt.Printf("输入边(vi, vj)上的下表i, 下表j和权w:\n")
		fmt.Scanf("%d,%d,%d", &i, &j, &w)
		G.arc[i][j] = w
		G.arc[j][i] = G.arc[i][j]
	}
}
