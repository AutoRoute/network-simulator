package netsim

import (
	"testing"
)

func TestAddNeighbor(t *testing.T) {
	node1 := NewNode("A", nil)
	node2 := NewNode("B", nil)

	if err := node1.AddNeighbor(node1); err == nil {
		t.Fatalf("Node %s added itself as neighbor", node1.ID)
	}

	if err := node1.AddNeighbor(node2); err != nil {
		t.Fatal(err)
	}

	if err := node1.AddNeighbor(node2); err == nil {
		t.Fatalf("Node %s added neighbor %s twice", node1.ID, node2.ID)
	}
}
