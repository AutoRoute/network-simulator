package netsim

import (
	"errors"
	"fmt"
)

// NodeAddress is an AutoRoute node identifier
type NodeAddress string

// A Node is an AutoRoute packet router in the network.
// It is a aware of its neighbors and can send a packet
// to any Node it is directly connected to.
type Node struct {
	ID          NodeAddress
	connections map[NodeAddress]*Connection
	Packets     chan *Packet
	delivered   chan *Packet
}

// NewNode takes a NodeAddress and returns a *Node. The
// new Node is not connected to any other Node.
func NewNode(id NodeAddress, delivered chan *Packet) *Node {
	return &Node{id, make(map[NodeAddress]*Connection), make(chan *Packet), delivered}
}

// AddConnection takes a *Connection and adds it to this Node's
// connection map.
func (n Node) AddConnection(conn *Connection) error {
	neighbor := conn.GetNeighbor(n.ID)
	if n.ID == neighbor {
		msg := fmt.Sprintf("Node %s cannot add connection with itself", n.ID)
		return errors.New(msg)
	}
	if _, ok := n.connections[neighbor]; !ok {
		n.connections[neighbor] = conn
		return nil
	}

	msg := fmt.Sprintf("Node %s already has connection with neighbor %s", n.ID, neighbor)
	return errors.New(msg)
}

// SendPacket takes a *Packet and sends it
// to a neighbor determined by the Node's
// routing algorithm
func (n Node) SendPacket(p *Packet) error {
	if p.Source == n.ID {
		n.delivered <- p
		return nil
	}

	return nil
}
