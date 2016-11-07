package netsim

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("examples/test_config.json")
	if err != nil {
		t.Fatal(err)
	}

	if config.Graph.Nodes[0].ID != "A" ||
		config.Graph.Nodes[1].ID != "B" ||
		config.Graph.Edges[0].First != "A" ||
		config.Graph.Edges[0].Second != "B" {
		t.Fatal("Config not parsed correctly")
	}
}

func TestBadFile(t *testing.T) {
	_, err := LoadConfig("fake/path")
	if err == nil {
		t.Fatal("Didn't fail on non-existent file")
	}
}

func TestBadJSON(t *testing.T) {
	_, err := LoadConfig("examples/bad.json")
	if err == nil {
		t.Fatal("Didn't fail on bad JSON")
	}
}
