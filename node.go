package netsim

import (
	"errors"
	"fmt"
)

type NodeAddress string

type Node struct {
	ID        NodeAddress
	Neighbors map[NodeAddress]*Node
}

func NewNode(id NodeAddress) *Node {
	return &Node{id, make(map[NodeAddress]*Node)}
}

func (n Node) AddNeighbor(node *Node) error {
	if n.ID == node.ID {
		msg := fmt.Sprintf("Node %s cannot add itself as neighbor", n.ID)
		return errors.New(msg)
	}
	if _, ok := n.Neighbors[node.ID]; !ok {
		n.Neighbors[node.ID] = node
		return nil
	} else {
		msg := fmt.Sprintf("Node %s already has neighbor %s", n.ID, node.ID)
		return errors.New(msg)
	}
}
