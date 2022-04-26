package datastruct

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

type PriorityQueue []*Vertex

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
