package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := LoadConfig("examples/test_config.json")
	if err != nil {
		t.Fatal(err)
	}

	if config.Network[0].ID != "A" ||
		config.Network[0].Neighbors[0] != "B" ||
		config.Network[0].Neighbors[1] != "C" ||
		config.Network[1].ID != "B" ||
		config.Network[1].Neighbors[0] != "A" ||
		config.Network[1].Neighbors[1] != "C" ||
		config.Network[2].ID != "C" ||
		config.Network[2].Neighbors[0] != "A" ||
		config.Network[2].Neighbors[1] != "B" {
		t.Fatal("Configs are not the same")
	}
}

func TestInvalidConfig(t *testing.T) {
	_, err := LoadConfig("examples/bad_config.json")
	if err == nil {
		t.Fatal("Invalid config was validated")
	}
}
