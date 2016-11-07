package netsim

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Graph *GraphConfig
}

type GraphConfig struct {
	Nodes []*NodeConfig
	Edges []*EdgeConfig
}

type NodeConfig struct {
	ID NodeAddress
}

type EdgeConfig struct {
	First  NodeAddress
	Second NodeAddress
}

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
