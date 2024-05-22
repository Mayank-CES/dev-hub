package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v is already present", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	var err error
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil {
		err = fmt.Errorf("fromVertex is not found")
		fmt.Println(err.Error())
	} else if toVertex == nil {
		err = fmt.Errorf("toVertex is not found")
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err = fmt.Errorf("edge is already present")
		fmt.Println(err.Error())
	} else { // add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex witht a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	graph := &Graph{}
	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}
	graph.Print()
	/*
		Vertex 0:
		Vertex 1:
		Vertex 2:
		Vertex 3:
		Vertex 4:
	*/

	graph.AddEdge(1, 2)
	graph.Print()
	/*
		Vertex 0:
		Vertex 1:  2
		Vertex 2:
		Vertex 3:
		Vertex 4:
	*/

	graph.AddEdge(1, 7) // toVertex is not found

	graph.AddEdge(1, 2) // edge is already present
}
