package main

import "fmt"

type Address struct {
	Addr int
}

type Host struct {
	Address         Address
	PhysicalAddress string // TODO use this
	Neighbours      []*Host
}

func (h *Host) Receive(m Message) {
	fmt.Printf("[%d] receiveved message from [%d]: %s\n", h.Address.Addr, m.Source.Address.Addr, m.Content)
}

type Message struct {
	Source      Host
	Destination Host
	Content     string
}

type Graph struct { // a network is just a graph
	Hosts map[Address]*Host
}

func NewGraph() *Graph {
	return &Graph{
		Hosts: make(map[Address]*Host),
	}
}

func (g *Graph) AddNode(h Host) {
	if _, exists := g.Hosts[h.Address]; !exists {
		newNode := &Host{
			Address:    h.Address,
			Neighbours: []*Host{},
		}
		g.Hosts[h.Address] = newNode
	}
}

func (g *Graph) AddEdge(n1, n2 Address) {
	node1 := g.Hosts[n1]
	node2 := g.Hosts[n2]

	node1.Neighbours = append(node1.Neighbours, node2)
	node2.Neighbours = append(node2.Neighbours, node1)
}

func main() {

	net := NewGraph()

	for i := 0; i < 10; i++ {
		// TODO (nick): is this useless wrapping of Address?
		addr := Address{Addr: i}
		h := Host{Address: addr}
		net.AddNode(h)
	}

	// adjacency list represeting the network (which is a graph)
	net.AddEdge(Address{1}, Address{2})
	net.AddEdge(Address{1}, Address{3})
	net.AddEdge(Address{1}, Address{4})
	net.AddEdge(Address{1}, Address{5})
	net.AddEdge(Address{1}, Address{6})
	net.AddEdge(Address{6}, Address{3})
	net.AddEdge(Address{3}, Address{2})
	net.AddEdge(Address{1}, Address{7})
	net.AddEdge(Address{1}, Address{8})
	net.AddEdge(Address{8}, Address{6})
	net.AddEdge(Address{1}, Address{9})

	sourceHost := net.Hosts[Address{2}]
	destHost := net.Hosts[Address{8}]

	msg := Message{Source: *sourceHost, Destination: *destHost, Content: "Hello, World!"}

	bfsDeliverMessage(msg, sourceHost)
	// at node {2}
	// at node {1}
	// at node {3}
	// at node {4}
	// at node {5}
	// at node {6}
	// at node {7}
	// at node {8}
	// [8] receiveved message from [2]: Hello, World!
	//
	// Although there were other valid routes from 8 to 2, (such as 2, 3, 6, 8),
	// BFS took the shortest path (2, 1 , 8) like it is supposed to.
}

// we can do BFS to find nodes. If the node has a route to our destination then we can
// jump through it to continue closer to the destination
func bfsDeliverMessage(m Message, root *Host) bool {
	visited := make(map[*Host]bool)
	route := make(map[*Host]*Host)

	q := []*Host{}
	q = append(q, root)
	visited[root] = true
	route[root] = nil

	for len(q) > 0 {
		layerSize := len(q)

		for i := 0; i < layerSize; i++ {
			node := q[0]
			q = q[1:]
			fmt.Printf("at node %d\n", node.Address)
			if node.Address == m.Destination.Address {
				node.Receive(m)
				traceRoute(node, route)

				return true
			}

			for _, neighbor := range node.Neighbours {
				if !visited[neighbor] {
					q = append(q, neighbor)
					visited[neighbor] = true
					route[neighbor] = node
				}
			}
		}
	}
	fmt.Printf("No route found from %d to %d", m.Source.Address, m.Destination.Address)
	return false
}

func traceRoute(node *Host, parent map[*Host]*Host) {
	stack := []*Host{}
	for node != nil {
		stack = append(stack, node)
		node = parent[node]
	}

	fmt.Printf("Path taken aka route trace: ")
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stack[i].Address.Addr)
	}
	fmt.Println()
}
