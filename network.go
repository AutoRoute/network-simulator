package netsim

import (
	"errors"
	"fmt"
)

// A Network represents an AutoRoute network of connected nodes.
type Network struct {
	nodes map[NodeAddress]*Node
}

// NewNetwork takes a *Config object and returns a *Network and error.
// Based off the config, it creates Node objects from the graph's
// NodeConfig and adds neighbors to the nodes from the graph's EdgeConfig
func NewNetwork(config *Config, delivered chan *Packet) (*Network, error) {
	network := &Network{make(map[NodeAddress]*Node)}
	for _, n := range config.Graph.Nodes {
		if err := network.addNode(n, delivered); err != nil {
			return nil, err
		}
	}

	for _, e := range config.Graph.Edges {
		if err := network.addEdge(e); err != nil {
			return nil, err
		}
	}
	return network, nil
}

func (n Network) addNode(node *NodeConfig, delivered chan *Packet) error {
	if _, ok := n.nodes[node.ID]; !ok {
		n.nodes[node.ID] = NewNode(node.ID, delivered)
		return nil
	}

	msg := fmt.Sprintf("Node %s already in network", node.ID)
	return errors.New(msg)
}

func (n Network) addEdge(edge *EdgeConfig) error {
	node1, ok := n.nodes[edge.First]
	if !ok {
		msg := fmt.Sprintf("Cannot add edge %s-%s: Node %s not in network",
			edge.First, edge.Second, edge.First)
		return errors.New(msg)
	}
	node2, ok := n.nodes[edge.Second]
	if !ok {
		msg := fmt.Sprintf("Cannot add edge %s-%s: Node %s not in network",
			edge.First, edge.Second, edge.Second)
		return errors.New(msg)
	}

	conn := NewConnection(node1.ID, node2.ID, node1.Packets, node2.Packets)
	if err := node1.AddConnection(conn); err != nil {
		return err
	}
	if err := node2.AddConnection(conn); err != nil {
		// This should never be reached since neighborship should
		// be recorded on both nodes
		return err
	}
	return nil
}

func (n Network) getNode(id NodeAddress) (*Node, error) {
	node, ok := n.nodes[id]
	if !ok {
		msg := fmt.Sprintf("Cannot find node %s in network", id)
		return nil, errors.New(msg)
	}

	return node, nil
}
