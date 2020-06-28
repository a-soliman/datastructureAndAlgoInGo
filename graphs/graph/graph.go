package graph

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
