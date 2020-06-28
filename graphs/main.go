package main

import (
	"fmt"

	"github.com/a-soliman/datastructureAndAlgoInGo/graphs/graph"
)

func main() {
	graph := graph.NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 6)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(4, 5)
	fmt.Printf("GraphBFS: %v\n", graph.BFS(1))
	fmt.Printf("GraphDFS: %v\n", graph.DFS(1))
	fmt.Printf("HasPath: 1, 3: %v\n", graph.HasPath(1, 3))
}
