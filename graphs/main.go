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

/*
Knight's Tour On A Chess Board
You are given a rows * cols chessboard and a knight that moves like in normal chess. Currently knight is at starting position denoted by start_row th row and start_col th col, and want to reach at ending position denoted by end_row th row and end_col th col.  The goal is to calculate the minimum number of moves that the knight needs to take to get from starting position to ending position.
start_row, start_col, end_row and end_col are 0-indexed.

Example
Input:
rows = 5
cols = 5
start_row = 0
start_col = 0
end_row = 4
end_col = 1

Output: 3
3 moves to reach from (0, 0) to (4, 1):
(0, 0) -> (1, 2) -> (3, 3) -> (4, 1).
*/

func findMinimumNumberOfMoves(rows int32, cols int32, startRow int32, startCol int32, endRow int32, endCol int32) int32 {
	visited := buildBoard(rows, cols)
	queue := [][]int32{
		{startRow, startCol},
	}
	var current []int32
	distance := int32(0)
	internalCount := 1

	for len(queue) > 0 {
		if internalCount == 0 {
			internalCount = len(queue)
			distance++
		}
		current = queue[0]
		queue = queue[1:]
		internalCount--
		r, c := current[0], current[1]
		if r == endRow && c == endCol {
			return distance
		}
		potentialNextMoves := getValidNextMoves(rows, cols, r, c)
		for _, move := range potentialNextMoves {
			if visited[move[0]][move[1]] == -1 {
				visited[move[0]][move[1]] = 1
				queue = append(queue, move)
			}
		}
	}
	return -1
}

func buildBoard(rows, cols int32) [][]int32 {
	board := make([][]int32, rows)

	for i := int32(0); i < rows; i++ {
		row := make([]int32, cols)
		for j := int32(0); j < cols; j++ {
			row[j] = -1
		}
		board[i] = row
	}
	return board
}

func getValidNextMoves(rows, cols, startRow, startCol int32) [][]int32 {
	res := [][]int32{}
	moves := [][]int32{
		{startRow + 2, startCol - 1},
		{startRow + 2, startCol + 1},
		{startRow + 1, startCol - 2},
		{startRow + 1, startCol + 2},
		{startRow - 2, startCol - 1},
		{startRow - 2, startCol + 1},
		{startRow - 1, startCol - 2},
		{startRow - 1, startCol + 2},
	}

	for _, move := range moves {
		r, c := move[0], move[1]

		if r < rows && r >= 0 && c < cols && c >= 0 {
			res = append(res, move)
		}
	}
	return res
}

/*
String Transformation Using Given Dictionary Of Words
You are given a dictionary of words and two strings, start and stop. All given strings have equal length.
Transform string start to string stop one character per step using words from the dictionary. For example, "abc" -> "abd" is a valid transformation step because only one character is changed (c->d) while "abc" -> "axy" is not a valid step transformation because two characters are changed (c->x and c->y).
You need to find the shortest possible sequence of strings (two or more) such that:

First string is start.
Last string is stop.
Every string (except the first one) differs from the previous one by exactly one character.
Every string (except, possibly, first and last ones) are in the dictionary of words.
i.e. output = [start, <strings from the given dictionary>, stop] and len(output) >= 2.
If two or more such sequences exist, any one of them is a correct answer.
If no such sequence is there to be found, [“-1”] (a sequence of one string, “-1”) is the correct answer.

Example One
Input:
words = ["cat", "hat", "bad", "had"]
start = "bat"
stop = "had"
Output:
["bat", "bad", "had"]
or
["bat", "hat", "had"]
From "bat" change character 't' to 'd', so new string will be "bad".
From "bad" change character 'b' to 'h', so new string will be "had".
or
From "bat" change character 'b' to 'h', so new string will be "hat".
From "hat" change character 't' to 'd', so new string will be "had".

Example Two
Input:
words = []
start = bbb
stop = bbc
Output: ["bbb", "bbc"]
From "bbb" change the last character 'b' to 'c', so new string will be "bbc".

Example Three
Input:
words = []
start = "zzzzzz"
stop = "zzzzzz"
Output: [-1]

Function must return an array of strings of length >= 2, where the first string is start and the last string is stop, if the transformation is possible. Else return an array of strings containing only one string "-1", i.e. return ["-1"].
Here, the words dictionary is empty and ["zzzzzz", "zzzzzz"] is not a valid transformation hence return ["-1"].

Example Four
Input:
words = ["cccw", "accc", "accw"]
start = "cccc"
stop = "cccc"
Output:
["cccc", "cccw", "cccc"]
Or:
["cccc", "accc", "cccc"]
*/
func stringTransformation(words []string, start string, stop string) []string {
	visited := make(map[string]string)
	queue := []string{start}
	var current string
	res := []string{}

	if oneStepAway(start, stop) {
		return []string{start, stop}
	}

	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		if oneStepAway(current, stop) {
			// here should be magic
			//recBuildRes(&visited, current, start, &res)
			res = append(res, stop)
			parent, child := visited[current], current
			for child != "" {
				res = append(res, child)
				child = parent
				parent = visited[child]
			}
			reverse(&res)
			return res
		}
		for _, word := range words {
			if oneStepAway(current, word) && visited[word] == "" {
				visited[word] = current
				queue = append(queue, word)
			}
		}
	}
	return []string{"-1"}
}

func reverse(input *[]string) {
	i, j := 0, len(*input)-1
	for i < j {
		(*input)[i], (*input)[j] = (*input)[j], (*input)[i]
		i++
		j--
	}
}

func oneStepAway(word, otherWord string) bool {
	diff := 0
	i, j := 0, 0
	for i < len(word) && j < len(otherWord) {
		if word[i] != otherWord[j] {
			diff++
		}
		i++
		j++
	}
	for i < len(word) {
		diff++
		i++
	}
	for j < len(otherWord) {
		diff++
		j++
	}
	return diff == 1
}

func recBuildRes(dict *map[string]string, word string, start string, res *[]string) {
	parent, found := (*dict)[word]
	if found {
		recBuildRes(dict, parent, start, res)
	}
	if word == start {
		*res = append(*res, word)
		return
	}
	*res = append(*res, word)
}

/*
	HasCycle
	given n nodes, and a list of edges
	determine of the nodes have a cycle
*/
func hasCycleUndirected(n int, edges [][]int) bool {
	adjList := buildAdjList(n, edges)
	visited := make([]bool, n)
	for i := range visited {
		if visited[i] == false {
			incluedesCycle := hasCycleUtil(i, -1, &adjList, &visited)
			if incluedesCycle {
				return true
			}
		}
	}
	return false
}

func hasCycleUtil(idx int, parent int, adjList *[][]int, visited *[]bool) bool {
	(*visited)[idx] = true
	for _, neighbor := range (*adjList)[idx] {
		if (*visited)[neighbor] == true && neighbor != parent {
			return true
		} else if (*visited)[neighbor] == false {
			hasCycle := hasCycleUtil(neighbor, idx, adjList, visited)
			if hasCycle {
				return true
			}
		}
	}
	return false
}

func hasCycleDirected(n int, edges [][]int) bool {
	adjList := buildDirectedAdjList(n, edges)
	visited := make([]bool, n)
	for i, isVisited := range visited {
		if !isVisited {
			hasCycle := hasCycleDirectedUtil(i, &adjList, &visited)
			if hasCycle {
				return true
			}
		}
	}
	return false
}

func buildDirectedAdjList(n int, edges [][]int) [][]int {
	res := make([][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		res[from] = append(res[from], to)
	}
	return res
}

func hasCycleDirectedUtil(idx int, adjList *[][]int, visited *[]bool) bool {
	(*visited)[idx] = true
	for _, neighbor := range (*adjList)[idx] {
		if (*visited)[neighbor] == true {
			return true
		}
		hasCycle := hasCycleDirectedUtil(neighbor, adjList, visited)
		if hasCycle {
			return true
		}
	}
	return false
}

/*
TransposeGraph
Transpose of a Graph G is a graph G' that has the same set of vertices, but the direction of edges is reveres.
*/
type graphNode struct {
	value     int
	neighbors []*graphNode
}

func transposeGraph(n int, edges [][]int) *graphNode {
	nodesMap := make(map[int]*graphNode)
	for i := 0; i < n; i++ {
		nodesMap[i] = &graphNode{i, []*graphNode{}}
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		nodesMap[to].neighbors = append(nodesMap[to].neighbors, nodesMap[from])
	}
	return nodesMap[edges[0][1]]
}

/*
	StronglyConnectedComponents
	return the strongly connected components of a graph
*/
func stronglyConnectedComponents(n int, edges [][]int) [][]int {
	adjList := buildDirectedAdjList(n, edges)
	visited := make([]bool, n)
	stack := []int{}
	res := [][]int{}

	for i := 0; i < n; i++ {
		stronglyConnectedComponentsDfsUtil(i, &adjList, &visited, &stack)
	}

	reversedEdges := reverseEdges(edges)
	reversedAdjList := buildDirectedAdjList(n, reversedEdges)
	visited = make([]bool, n)
	stack2 := []int{}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		if visited[current] != true {
			stronglyConnectedComponentsDfsUtil(current, &reversedAdjList, &visited, &stack2)
			copiedComponents := make([]int, len(stack2))
			copy(copiedComponents, stack2)
			res = append(res, copiedComponents)
			stack2 = []int{}
		}
	}
	return res
}

func reverseEdges(adjList [][]int) [][]int {
	res := make([][]int, len(adjList))
	for i, item := range adjList {
		from, to := item[0], item[1]
		res[i] = []int{to, from}
	}
	return res
}

func stronglyConnectedComponentsDfsUtil(idx int, adjList *[][]int, visited *[]bool, stack *[]int) {
	if (*visited)[idx] == true {
		return
	}
	(*visited)[idx] = true
	neighbors := (*adjList)[idx]
	for _, neighbor := range neighbors {
		if (*visited)[neighbor] == false {
			stronglyConnectedComponentsDfsUtil(neighbor, adjList, visited, stack)
		}
	}
	*stack = append(*stack, idx)
}

/*
Zombie Clusters
There are zombies in Seattle. Liv and Ravi are trying to track them down to find out who is creating new zombies in an effort to prevent an apocalypse. Other than the patient-zero zombies (who became so by mixing MaxRager and tainted Utopium), new people only become zombies after being scratched by an existing zombie. Zombiism is transitive. This means that if zombie 0 knows zombie 1 and zombie 1 knows zombie 2, then zombie 0 is connected to zombie 2 by way of knowing zombie 1. A zombie cluster is a group of zombies who are directly or indirectly linked through the other zombies they know, such as the one who scratched them or supplies who them with brains.
We have a two-dimensional array with n rows and n columns where each cell, zombies[A][B], denotes whether zombie A knows zombie B. The diagram showing connectedness will be made up of a number of binary strings, characters 0 or 1. Each of the characters in the string represents whether the zombie associated with a row element is connected to the zombie at that character's index. For instance, a zombie 0 with a connectedness string '110' is connected to zombies 0 (itself) and zombie 1, but not to zombie 2. The complete matrix of zombie connectedness is:

110
110
001

Zombies 0 and 1 are connected. Zombie 2 is not.
Your task is to determine the number of connected groups of zombies, or clusters, in a given matrix.

Example One
Input:
[“1100”,
 “1110”,
 “0110”,
 “0001”]
Output: 2
*/
func zombieCluster(zombies []string) int32 {
	size := len(zombies)
	visited := make([]bool, size)
	var res int32
	for i := 0; i < size; i++ {
		if visited[i] == false {
			res++
			dfsUtil(i, zombies, &visited)
		}
	}
	return res
}

func dfsUtil(idx int, zombies []string, visited *[]bool) {
	(*visited)[idx] = true
	neighbors := zombies[idx]
	for neighborIdx := 0; neighborIdx < len(neighbors); neighborIdx++ {
		if string(neighbors[neighborIdx]) == "1" && (*visited)[neighborIdx] == false {
			(*visited)[neighborIdx] = true
			dfsUtil(neighborIdx, zombies, visited)
		}
	}
}

func findCriticalConnections(noOfServers int32, noOfConnections int32, connections [][]int32) [][]int32 {
	adjList := buildConnectionsAdjList(noOfServers, connections)
	ids := make(map[int32]int32)
	visited := make([]bool, noOfServers)
	parentsIds := make([]int32, noOfServers)
	minAccessable := make([]int32, noOfServers)
	res := [][]int32{}
	// assign ids to servers
	var id int32 = 0
	for i := 0; i < int(noOfServers); i++ {
		ids[int32(i)] = id
		minAccessable[i] = id
		id++
	}
	visited[0] = true
	parentsIds[0] = -1
	connectionDfsUtil(0, adjList, &ids, &visited, &parentsIds, &minAccessable, &res)
	if len(res) == 0 {
		res = append(res, []int32{-1, -1})
	}
	return res
}

func buildConnectionsAdjList(n int32, edges [][]int32) [][]int32 {
	res := make([][]int32, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		res[from] = append(res[from], to)
		res[to] = append(res[to], from)
	}
	return res
}

func connectionDfsUtil(nodeIdx int32, adjList [][]int32, ids *map[int32]int32, visited *[]bool, parentsIds *[]int32, minAccessable *[]int32, res *[][]int32) int32 {
	serverID := (*ids)[nodeIdx]
	minAccessibleServer := (*minAccessable)[nodeIdx]

	parentID := (*parentsIds)[nodeIdx]
	neighbors := adjList[nodeIdx]
	for _, neighbor := range neighbors {
		neighborID := (*ids)[neighbor]
		if (*visited)[neighbor] {
			if neighborID != parentID && (*minAccessable)[neighbor] < minAccessibleServer {
				(*minAccessable)[nodeIdx] = (*minAccessable)[neighbor]
				minAccessibleServer = (*minAccessable)[neighbor]
			}
		} else {
			// it wanst visited
			(*visited)[neighbor] = true
			(*parentsIds)[neighbor] = serverID
			minFound := connectionDfsUtil(neighbor, adjList, ids, visited, parentsIds, minAccessable, res)
			if minFound > minAccessibleServer {
				*res = append(*res, []int32{nodeIdx, neighbor})
			} else if minFound < minAccessibleServer {
				(*minAccessable)[nodeIdx] = minFound
				minAccessibleServer = minFound
			}
		}
	}
	return int32(minAccessibleServer)
}

func main() {
	graph := graph.NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 6)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(3, 5)
	graph.AddEdge(3, 4)
	fmt.Printf("GraphBFS: %v\n", graph.BFS(1))
	fmt.Printf("GraphDFS: %v\n", graph.DFS(1))
	fmt.Printf("HasPath: 1, 3: %v\n", graph.HasPath(1, 3))
	fmt.Printf("CountPaths: From :1, To :5 => %d\n", graph.CountAllPaths(1, 5))
	fmt.Printf("FindPaths: From :1, To :5 => %v\n", graph.FindAllPaths(1, 5))
	fmt.Printf("Distance: From :1, To :5 => %d\n", graph.Distance(1, 5))

	countComponentsOutput := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})
	fmt.Printf("CountComponents: %d\n", countComponentsOutput)

	validTreeOutput := isValidTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}})
	fmt.Printf("validTree (VAlid): %v\n", validTreeOutput)

	validTreeOutput = isValidTree(5, [][]int{{0, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 3}})
	fmt.Printf("validTree (Invalid): %v\n", validTreeOutput)

	findMinimumNumberOfMovesOutput := findMinimumNumberOfMoves(4, 24975, 3, 21841, 1, 13)
	fmt.Printf("Knight Chess Board: %d\n", findMinimumNumberOfMovesOutput)

	fmt.Println("\nStringTransformation:")
	stringTransformationInput, start, stop := []string{"cat", "caz", "hat", "bad", "had"}, "bat", "had"
	stringTransformationOutput := stringTransformation(stringTransformationInput, start, stop)
	fmt.Printf("Input: %v\nStart: %s\nStop: %s\nOutput: %v\n", stringTransformationInput, start, stop, stringTransformationOutput)

	stringTransformationInput, start, stop = []string{"cccw", "accc", "accw"}, "cccc", "cccc"
	stringTransformationOutput = stringTransformation(stringTransformationInput, start, stop)
	fmt.Printf("Input: %v\nStart: %s\nStop: %s\nOutput: %v\n", stringTransformationInput, start, stop, stringTransformationOutput)

	stringTransformationInput, start, stop = []string{}, "bbb", "bbc"
	stringTransformationOutput = stringTransformation(stringTransformationInput, start, stop)
	fmt.Printf("Input: %v\nStart: %s\nStop: %s\nOutput: %v\n", stringTransformationInput, start, stop, stringTransformationOutput)

	stringTransformationInput, start, stop = []string{}, "zzzzz", "zzzzz"
	stringTransformationOutput = stringTransformation(stringTransformationInput, start, stop)
	fmt.Printf("Input: %v\nStart: %s\nStop: %s\nOutput: %v\n", stringTransformationInput, start, stop, stringTransformationOutput)

	hasCycleInput := [][]int{{0, 1}, {0, 2}, {3, 4}, {3, 5}, {4, 5}}
	hasCycleOutput := hasCycleUndirected(6, hasCycleInput)
	fmt.Printf("Has Cycle <<undirected graph>>: %v\n", hasCycleOutput)

	hasCycleInput = [][]int{{0, 1}, {0, 2}, {2, 3}, {1, 3}, {3, 4}, {4, 1}}
	hasCycleOutput = hasCycleDirected(5, hasCycleInput)
	fmt.Printf("Has Cycle <<directed graph>>: %v\n", hasCycleOutput)

	transposeGraphOutput := transposeGraph(4, [][]int{{0, 1}, {0, 2}, {1, 2}, {2, 3}})
	fmt.Printf("TransposeGraph: %v\n", transposeGraphOutput.neighbors)

	stronglyConnectedComponentsOutput := stronglyConnectedComponents(7, [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 3},
		{5, 6},
	})
	fmt.Printf("StronglyConnectedComponents: %v\n", stronglyConnectedComponentsOutput)

	zombieOutput := zombieCluster([]string{"1100", "1110", "0110", "0001"})
	fmt.Printf("Zombie: %d\n", zombieOutput)

	criticalConnectionsOutput := findCriticalConnections(7, 8, [][]int32{
		{0, 1},
		{1, 2},
		{2, 0},
		{1, 3},
		{1, 4},
		{1, 6},
		{3, 5},
		{4, 5},
	})
	fmt.Printf("CriticalConnections: %v\n", criticalConnectionsOutput)
}
