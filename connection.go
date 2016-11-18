package netsim

// Connection is a proxy for sending packets
// between two Nodes
type Connection struct {
	node1    NodeAddress
	node2    NodeAddress
	packets1 <-chan *Packet
	packets2 <-chan *Packet
}

// NewConnection takes two NodeAddresses and two *Packet channels
// and creates a *Connection object for the two nodes.
func NewConnection(id1 NodeAddress, id2 NodeAddress, p1 chan *Packet, p2 chan *Packet) *Connection {
	return &Connection(id1, id2, p1, p2)
}

// GetNeighbor takes the NodeAddress of one of the
// connected Nodes and returns the NodeAddress of
// its neighbor
func (c Connection) GetNeighbor(me NodeAddress) NodeAddress {
	if me == c.node1 {
		return c.node2
	}
	return c.node1
}
