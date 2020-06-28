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
