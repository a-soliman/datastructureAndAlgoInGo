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
