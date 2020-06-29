package main

import (
	"fmt"

	"github.com/a-soliman/datastructureAndAlgoInGo/graphs/graph"
)

/*
CountTheComponents
given an n int as vertices start from 0 to n-1, and given an arrays represents the edges between these components
find how many separate components does the graph contain
*/
func countComponents(n int, edges [][]int) int {
	// create a graph
	adjList := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adjList[from] = append(adjList[from], to)
		adjList[to] = append(adjList[to], from)
	}
	visited := make([]int, n)
	components := 0
	for i := range visited {
		visited[i] = 0
	}
	for i, item := range visited {
		if item == 0 {
			components++
			dfs(&adjList, i, &visited)
		}
	}
	return components
}

func dfs(list *[][]int, idx int, visited *[]int) {
	(*visited)[idx] = 1
	neighbors := (*list)[idx]
	for _, neighbor := range neighbors {
		if (*visited)[neighbor] == 0 {
			dfs(list, neighbor, visited)
		}
	}
}

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

	countComponentsOutput := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})
	fmt.Printf("CountComponents: %d\n", countComponentsOutput)
}
