package graph

import "errors"

// Vertex type
type Vertex struct {
	Value int
	Edges map[int]*Vertex
}

func (v *Vertex) getEdges() []int {
	keys := make([]int, 0, len(v.Edges))
	for k := range v.Edges {
		keys = append(keys, k)
	}
	return keys
}

func (v *Vertex) addEdge(toVertex *Vertex) {
	v.Edges[toVertex.Value] = toVertex
}

func (v *Vertex) removeEdge(to int) {
	delete(v.Edges, to)
}

// Graph type
type Graph struct {
	Vertices map[int]*Vertex
	V        int
	E        int
}

// New returns a new graph
func New() *Graph {
	verticesMap := make(map[int]*Vertex)
	return &Graph{verticesMap, 0, 0}
}

// NewFromSlice returns a new Graph with the given values as vertices
func NewFromSlice(input []int) *Graph {
	graph := New()
	for _, item := range input {
		graph.AddVertex(item)
	}
	return graph
}

// AddVertex adds a new vertex
func (g *Graph) AddVertex(value int) {
	vertex := &Vertex{value, make(map[int]*Vertex)}
	g.Vertices[value] = vertex
	g.V++
}

// RemoveVertex removes a vertex
func (g *Graph) RemoveVertex(value int) error {
	vertex, found := g.Vertices[value]
	if !found {
		return errors.New("Vertex was not found")
	}
	neighbors := vertex.getEdges()
	for _, neighbor := range neighbors {
		vertex := g.Vertices[neighbor]
		vertex.removeEdge(value)
		g.E--
	}
	delete(g.Vertices, value)
	g.V--
	return nil
}

// AddEdge adds an edge between two given vertices
func (g *Graph) AddEdge(from, to int) error {
	fromVertex, found := g.Vertices[from]
	if !found {
		return errors.New("From node was not found")
	}
	toVertex, found := g.Vertices[to]
	if !found {
		return errors.New("To node was not found")
	}
	fromVertex.addEdge(toVertex)
	toVertex.addEdge(fromVertex)
	g.E++
	return nil
}

// GetEdges given a value of a vertex it returns the vertex neighbors
func (g *Graph) GetEdges(value int) ([]int, error) {
	vertex, found := g.Vertices[value]
	if !found {
		return []int{}, errors.New("Vertex was not found")
	}
	return vertex.getEdges(), nil
}

// BFS given a root value, it returns []int representing the BFS starting from that root
func (g *Graph) BFS(rootValue int) []int {
	root := g.Vertices[rootValue]
	res := []int{}
	queue := []*Vertex{root}
	visited := make(map[int]bool)
	visited[rootValue] = true
	var current *Vertex = nil

	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		res = append(res, current.Value)
		for _, item := range current.Edges {
			if !visited[item.Value] {
				queue = append(queue, item)
				visited[item.Value] = true
			}
		}
	}
	return res
}

// DFS given a root value, it returns []int representing the DFS starting from that root
func (g *Graph) DFS(rootValue int) []int {
	root := g.Vertices[rootValue]
	res := []int{}
	visited := make(map[int]bool)
	visited[rootValue] = true
	dfsUtil(root, &visited, &res)
	return res
}

func dfsUtil(vertex *Vertex, visited *map[int]bool, res *[]int) {
	*res = append(*res, vertex.Value)
	for _, item := range vertex.Edges {
		if !(*visited)[item.Value] {
			(*visited)[item.Value] = true
			dfsUtil(item, visited, res)
		}
	}
}
