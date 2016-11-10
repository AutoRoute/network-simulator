package netsim

import (
	"errors"
	"fmt"
)

// Node identifier
type NodeAddress string

// A Node is an AutoRoute packet router in the network.
// It is a aware of its neighbors and can send a packet
// to any Node it is directly connected to.
type Node struct {
	ID        NodeAddress
	Neighbors map[NodeAddress]*Node
}

// NewNode takes a NodeAddress and returns a *Node. The
// new Node is not connected to any other Node.
func NewNode(id NodeAddress) *Node {
	return &Node{id, make(map[NodeAddress]*Node)}
}

// AddNeighbor takes a Node (other) and adds it to this Node's
// neighbor map.
// Note: other Node will not have this Node in its neighbor map.
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
