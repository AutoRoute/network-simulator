package netsim

import (
	"errors"
	"fmt"
)

type Network struct {
	Nodes map[NodeAddress]*Node
}

func NewNetwork(config *Config) (*Network, error) {
	network := &Network{make(map[NodeAddress]*Node)}
	for _, n := range config.Graph.Nodes {
		if err := network.addNode(n); err != nil {
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

func (n Network) addNode(node *NodeConfig) error {
	if _, ok := n.Nodes[node.ID]; !ok {
		n.Nodes[node.ID] = NewNode(node.ID)
		return nil
	} else {
		msg := fmt.Sprintf("Node %s already in network", node.ID)
		return errors.New(msg)
	}
}

func (n Network) addEdge(edge *EdgeConfig) error {
	node1, ok := n.Nodes[edge.First]
	if !ok {
		msg := fmt.Sprintf("Cannot add edge %s-%s: Node %s not in network",
			edge.First, edge.Second, edge.First)
		return errors.New(msg)
	}
	node2, ok := n.Nodes[edge.Second]
	if !ok {
		msg := fmt.Sprintf("Cannot add edge %s-%s: Node %s not in network",
			edge.First, edge.Second, edge.Second)
		return errors.New(msg)
	}

	if err := node1.AddNeighbor(node2); err != nil {
		return err
	}
	if err := node2.AddNeighbor(node1); err != nil {
		// This should never be reached since neighborship should
		// be recorded on both nodes
		return err
	}
	return nil
}
