package netsim

import (
	"encoding/json"
	"io/ioutil"
)

// Config represent the JSON object of the network configuration.
type Config struct {
	Graph *GraphConfig
}

// GraphConfig is a graph representation of the network defining
// nodes and edges. This what the "graph" JSON object will be
// unmarshaled into.
type GraphConfig struct {
	Nodes []*NodeConfig
	Edges []*EdgeConfig
}

// NodeConfig represents one node in the network. Each element in the
// "nodes" JSON array will be unmarshaled into this.
type NodeConfig struct {
	ID NodeAddress
}

// EdgeConfig represent an edge between two connected nodes in
// the network. Each element in the "edge" JSON object will be
// unmarshaled into this.
type EdgeConfig struct {
	First  NodeAddress
	Second NodeAddress
}

// LoadConfig takes a string path to a network configuration
// JSON file and returns a *Config object containing the network
// graph. It also returns an error if there was a file I/O or
// JSON decoding error.
func LoadConfig(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
