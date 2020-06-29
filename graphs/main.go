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

/*
GraphIsValidTree
Given n nodes labels from 0 to n-1 and a list of undirected edges(each edge is a pair of nodes), write a function to check whether
these edges make up a valid
Example1:
Input: n =5, and edges = [[0,1], [0,2], [0,3], [1,4]]
Output: true

Example2:
Input: n = 5, and edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
Output: false
*/
func isValidTree(n int, edges [][]int) bool {
	adjList := buildAdjList(n, edges)
	visited, parents := make([]int, n), make([]int, n)
	for i := range visited {
		visited[i] = 0
	}
	for i := range parents {
		parents[i] = -1
	}
	noCircle := validTreeDFSUtil(&adjList, &visited, &parents, 0)
	oneComponent := true
	for _, item := range visited {
		if item != 1 {
			oneComponent = false
			break
		}
	}
	return noCircle && oneComponent
}

func buildAdjList(n int, edges [][]int) [][]int {
	adjList := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adjList[from] = append(adjList[from], to)
		adjList[to] = append(adjList[to], from)
	}
	return adjList
}

func validTreeDFSUtil(adjList *[][]int, visited *[]int, parents *[]int, idx int) bool {
	(*visited)[idx] = 1
	neighbors := (*adjList)[idx]
	for _, neighbor := range neighbors {
		if (*visited)[neighbor] == 0 {
			if (*parents)[neighbor] != -1 {
				return false
			}
			(*parents)[neighbor] = idx
			isValid := validTreeDFSUtil(adjList, visited, parents, neighbor)
			if !isValid {
				return false
			}
		} else {
			if neighbor != (*parents)[idx] {
				return false
			}
		}
	}
	return true
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

	validTreeOutput := isValidTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}})
	fmt.Printf("validTree (VAlid): %v\n", validTreeOutput)

	validTreeOutput = isValidTree(5, [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 3}})
	fmt.Printf("validTree (Invalid): %v\n", validTreeOutput)
}
