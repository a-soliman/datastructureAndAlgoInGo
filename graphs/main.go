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
}
