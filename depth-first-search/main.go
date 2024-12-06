package main

import "fmt"

type (
	Node struct {
		Name     string
		Children []*Node
	}

	Graph struct {
		nodes []*Node
	}
)

func NewGraph() *Graph {
	return &Graph{
		nodes: make([]*Node, 0),
	}
}

func (g *Graph) AddNode(name string, children ...*Node) {
	node := &Node{
		Name:     name,
		Children: children,
	}

	g.nodes = append(g.nodes, node)
}

func (n *Node) DFS(name string) *Node {
	stack := []*Node{n}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		fmt.Println("current: ", current)
		stack = stack[:len(stack)-1]
		if current.Name == name {
			return current
		}
		for _, child := range current.Children {
			stack = append(stack, child)
		}
	}
	return nil
}

func main() {
	fmt.Println("Depth First Search")

	graph := NewGraph()
	graph.AddNode("A", &Node{
		Name: "1", Children: []*Node{
			{Name: "X", Children: []*Node{}},
			{Name: "Y", Children: []*Node{}},
			{Name: "Z", Children: []*Node{}},
		}})

	fmt.Println("graph: ", graph)

	found := graph.nodes[0].DFS("Z")

	fmt.Println("found: ", found)
}
