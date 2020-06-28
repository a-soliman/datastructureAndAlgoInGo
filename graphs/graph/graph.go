package graph

// Vertex type
type Vertex struct {
	Value int
	Edges map[int]*Vertex
}

// Graph type
type Graph struct {
	Vertices map[int]*Vertex
	V        int
	E        int
}
