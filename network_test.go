package netsim

import (
	"testing"
)

func TestNewNetwork(t *testing.T) {
	config, err := LoadConfig("examples/test_config.json")
	if err != nil {
		t.Fatal(err)
	}

	network, err := NewNetwork(config)
	if err != nil {
		t.Fatal(err)
	}

	nodeA, ok := network.Nodes["A"]
	if !ok {
		t.Fatal("Node A not in network")
	}

	nodeB, ok := network.Nodes["B"]
	if !ok {
		t.Fatal("Node B not in network")
	}

	if _, ok = nodeA.Neighbors["B"]; !ok {
		t.Fatal("Node A doesn't have neighbor B")
	}

	if _, ok = nodeB.Neighbors["A"]; !ok {
		t.Fatal("Node B doesn't have neighbor A")
	}
}

func TestDuplicateNodes(t *testing.T) {
	config := &Config{
		&GraphConfig{
			[]*NodeConfig{
				&NodeConfig{"A"},
				&NodeConfig{"A"},
			},
			[]*EdgeConfig{},
		},
	}

	_, err := NewNetwork(config)
	if err == nil {
		t.Fatal("Didn't fail on duplicate nodes")
	}
}

func TestAddEdgeNotInNetwork(t *testing.T) {
	config := &Config{
		&GraphConfig{
			[]*NodeConfig{},
			[]*EdgeConfig{
				&EdgeConfig{"A", "B"},
			},
		},
	}

	_, err := NewNetwork(config)
	if err == nil {
		t.Fatal("Didn't fail on edge with nodes not in network")
	}

	config = &Config{
		&GraphConfig{
			[]*NodeConfig{
				&NodeConfig{"A"},
			},
			[]*EdgeConfig{
				&EdgeConfig{"A", "B"},
			},
		},
	}

	_, err = NewNetwork(config)
	if err == nil {
		t.Fatal("Didn't fail on edge with node not in network")
	}
}

func TestDuplicateEdge(t *testing.T) {
	config := &Config{
		&GraphConfig{
			[]*NodeConfig{
				&NodeConfig{"A"},
				&NodeConfig{"B"},
			},
			[]*EdgeConfig{
				&EdgeConfig{"A", "B"},
				&EdgeConfig{"A", "B"},
			},
		},
	}

	_, err := NewNetwork(config)
	if err == nil {
		t.Fatal("Didn't fail on duplicate edge")
	}
}
