package datastruct

import (
	"fmt"
	"math"
	"sync"
)

type ItemGraph struct {
	Nodes []*Node
	Edges map[Node][]*Edge
	lock  sync.RWMutex
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.Nodes = append(g.Nodes, n)
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node, weight int) {
	g.lock.Lock()
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Edge)
	}
	ed1 := Edge{
		Node:   n2,
		Weight: weight,
	}

	ed2 := Edge{
		Node:   n1,
		Weight: weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
	g.lock.Unlock()
}

// dijkstra implement
func getShortestPath(startNode *Node, endNode *Node, g *ItemGraph) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	//pq := make(PriorityQueue, 1)
	//heap.Init(&pq)
	q := NodeQueue{}
	pq := q.NewQ()
	start := Vertex{
		Node:     startNode,
		Distance: 0,
	}
	for _, nval := range g.Nodes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startNode.Value] = start.Distance
	pq.Enqueue(start)
	//im := 0
	for !pq.IsEmpty() {
		v := pq.Dequeue()
		if visited[v.Node.Value] {
			continue
		}
		visited[v.Node.Value] = true
		near := g.Edges[*v.Node]

		for _, val := range near {
			if !visited[val.Node.Value] {
				if dist[v.Node.Value]+val.Weight < dist[val.Node.Value] {
					store := Vertex{
						Node:     val.Node,
						Distance: dist[v.Node.Value] + val.Weight,
					}
					dist[val.Node.Value] = dist[v.Node.Value] + val.Weight
					//prev[val.Node.Value] = fmt.Sprintf("->%s", v.Node.Value)
					prev[val.Node.Value] = v.Node.Value
					pq.Enqueue(store)
				}
				//visited[val.Node.value] = true
			}
		}
	}
	fmt.Println(dist)
	fmt.Println(prev)
	pathval := prev[endNode.Value]
	var finalArr []string
	finalArr = append(finalArr, endNode.Value)
	for pathval != startNode.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	fmt.Println(finalArr)
	for i, j := 0, len(finalArr)-1; i < j; i, j = i+1, j-1 {
		finalArr[i], finalArr[j] = finalArr[j], finalArr[i]
	}
	return finalArr, dist[endNode.Value]

}

func CreateGraph(data InputGraph) *ItemGraph {
	var g ItemGraph
	nodes := make(map[string]*Node)
	for _, v := range data.Graph {
		if _, found := nodes[v.Source]; !found {
			nA := Node{v.Source}
			nodes[v.Source] = &nA
			g.AddNode(&nA)
		}
		if _, found := nodes[v.Destination]; !found {
			nA := Node{v.Destination}
			nodes[v.Destination] = &nA
			g.AddNode(&nA)
		}
		g.AddEdge(nodes[v.Source], nodes[v.Destination], v.Weight)
	}
	return &g
}

func GetShortestPath(from, to string, g *ItemGraph) *APIResponse {
	nA := &Node{from}
	nB := &Node{to}

	path, distance := getShortestPath(nA, nB, g)
	return &APIResponse{
		Path:     path,
		Distance: distance,
	}
}
