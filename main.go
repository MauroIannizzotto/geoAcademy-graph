package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
	adjacent map[int][]int
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex add a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the graph
// agrega una linea que va desde un nodo a otro, entonces le indicamos de donde hasta donde, pero para eso
// primero nos traemos la direccion en mem del nodo desde donde sale y el nodo hasta donde llega
func (g *Graph) AddEdge(from, to int) {
	//get vertex (obtenemos la direccion del nodo de donde sale o a donde llega)
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error (revisamos si existen los nodos o hay algun error)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge (agregamos la linea del grafo)
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
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

// HasEulerianPath verify if the graph has an eulerian path
func (g *Graph) HasEulerianPath() bool {
	//para saber si tiene un camino euleriano debe tener o dos o ningun vertice de grado impar, por lo que creo un contador
	oddDegreeCount := 0

	// Contar el número de vértices con grado impar
	for _, v := range g.vertices {
		vertexCounter := len(v.adjacent)
		if vertexCounter%2 != 0 {
			oddDegreeCount++
		}
	}

	// Verificar si el número de vértices con grado impar es 0 o 2
	return oddDegreeCount == 0 || oddDegreeCount == 2
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddVertex(4)
	test.AddEdge(1, 2)
	test.AddEdge(1, 2)
	test.AddEdge(3, 4)
	test.AddEdge(6, 2)
	test.Print()

	if test.HasEulerianPath() {
		fmt.Println("El grafo tiene un camino euleriano.")
	} else {
		fmt.Println("El grafo no tiene un camino euleriano.")
	}
}
