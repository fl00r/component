package component

import (
	"container/list"
	"errors"
)

// Graph represented by string vertices and linked list of strings for edges
type Graph struct {
	vertices map[string]map[string]bool
}

// NewGraph returns empty Graph type
func NewGraph() *Graph {
	return &Graph{make(map[string]map[string]bool)}
}

// AddEdge connects two vertices by adding vertice2 to vetice1's list of edges
func (g *Graph) AddEdge(verticeFrom, verticeTo string) *Graph {
	edgesFrom := g.vertices[verticeFrom]
	if edgesFrom == nil {
		g.vertices[verticeFrom] = make(map[string]bool)
		edgesFrom = g.vertices[verticeFrom]
	}
	edgesFrom[verticeTo] = true

	edgesTo := g.vertices[verticeTo]
	if edgesTo == nil {
		g.vertices[verticeTo] = make(map[string]bool)
	}

	return g
}

// TopologicalSort not-generic ^^ topological sort function for our Graph type
//
// Kahn's algorithm
func (g Graph) TopologicalSort() ([]string, error) {
	verticesCount := len(g.vertices)
	sortedVertices := make([]string, 0, verticesCount)

	startingVertices := list.New()
	for k, v := range g.vertices {
		if len(v) == 0 {
			startingVertices.PushBack(k)
		}
	}

	if startingVertices.Len() == 0 {
		return nil, errors.New("Graph is not acycle")
	}

	for v := startingVertices.Front(); v != nil; v = v.Next() {
		val := v.Value.(string)
		delete(g.vertices, val)
		sortedVertices = append(sortedVertices, val)
		for k, w := range g.vertices {
			if w[val] {
				delete(w, val)
				if len(w) == 0 {
					startingVertices.PushBack(k)
				}
			}
		}
	}

	if len(g.vertices) != 0 {
		return nil, errors.New("Graph is not acycle")
	}

	return sortedVertices, nil
}
