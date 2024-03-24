package main

import (
	"fmt"
	"sort"
)

type Graph struct {
	Vertices []string
	Edges    map[string]map[string]int
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: make([]string, 0),
		Edges:    make(map[string]map[string]int),
	}
}

func (g *Graph) AddVertex(vertex string) {
	if _, exists := g.Edges[vertex]; !exists {
		g.Vertices = append(g.Vertices, vertex)
		g.Edges[vertex] = make(map[string]int)
	}
}

func (g *Graph) AddEdge(source, destination string, weight int) {
	if _, exists := g.Edges[source]; !exists {
		g.AddVertex(source)
	}
	if _, exists := g.Edges[destination]; !exists {
		g.AddVertex(destination)
	}
	g.Edges[source][destination] = weight
}

func (g *Graph) DependencyIndex(M, N []string) int {
	count := 0
	for _, v := range M {
		for _, u := range N {
			if _, exists := g.Edges[v][u]; exists {
				count++
			}
		}
	}
	return count
}

func main() {
	graph := NewGraph()

	// Задаємо множини вершин графа
	vertexSets := map[string][]string{
		"M": {"A", "B", "C"},
		"N": {"X", "Y", "Z"},
	}

	// Задаємо зв'язки між вершинами графа
	edges := []struct {
		source, destination string
		weight              int
	}{
		{"A", "X", 1},
		{"B", "Y", 2},
		{"C", "Z", 3},
		{"A", "Y", 4},
		{"B", "Z", 5},
	}

	// Додаємо вершини та зв'язки у граф
	for _, edge := range edges {
		graph.AddEdge(edge.source, edge.destination, edge.weight)
	}

	// Сортуємо множини вершин графа
	M := vertexSets["M"]
	N := vertexSets["N"]
	sort.Strings(M)
	sort.Strings(N)

	// Визначаємо індекс залежності
	index := graph.DependencyIndex(M, N)

	// Виводимо результат
	fmt.Printf("Множина M: %v\n", M)
	fmt.Printf("Множина N: %v\n", N)
	fmt.Printf("Індекс залежності: %d\n", index)
}
