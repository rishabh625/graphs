package datastruct

import "sync"

type Node struct {
	Value string
}

type Edge struct {
	Node   *Node
	Weight int
}

type Vertex struct {
	Node     *Node
	Distance int
}

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	Lock  sync.RWMutex
}

type PriorityQueue []*Vertex

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

type InputGraph struct {
	Graph []InputData `json:"graph"`
	From  string      `json:"from"`
	To    string      `json:"to"`
}

type InputData struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
}

type APIResponse struct {
	Path     []string `json:"path"`
	Distance int      `json:"distance"`
}
