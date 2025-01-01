package graphs

import "fmt"

type Graph struct {
	AdjacentList map[string][]string
	NodeCount    int
}

func NewGraph() *Graph {
	return &Graph{
		AdjacentList: make(map[string][]string),
		NodeCount:    0,
	}
}

func (g *Graph) AddVertex(node string) {
	g.AdjacentList[node] = []string{}
	g.NodeCount++
}

func (g *Graph) AddEdge(node1 string, node2 string) {
	g.AdjacentList[node1] = append(g.AdjacentList[node1], node2)
	g.AdjacentList[node2] = append(g.AdjacentList[node2], node1)
}

func (g *Graph) PrintConnections() {
	for k, v := range g.AdjacentList {
		fmt.Printf("%s: %v\n", k, v)
	}
}
