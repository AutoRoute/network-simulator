package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type NodeAddress string

type Config struct {
	Network []*NodeConfig
}

type NodeConfig struct {
	ID        NodeAddress
	Neighbors []NodeAddress
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

	err = verifyConfig(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func verifyConfig(config *Config) error {
	links := make(map[string]int)
	for _, i := range config.Network {
		for _, j := range i.Neighbors {
			if i.ID == j {
				msg := fmt.Sprintf("Node %s cannot be it's own neighbor", i.ID)
				return errors.New(msg)
			}

			var pair string
			if i.ID < j {
				pair = strings.Join([]string{string(i.ID), string(j)}, ",")
			} else {
				pair = strings.Join([]string{string(j), string(i.ID)}, ",")
			}

			if _, ok := links[pair]; ok {
				links[pair]++
			} else {
				links[pair] = 1
			}
		}
	}

	for k, v := range links {
		if v == 1 {
			msg := fmt.Sprintf("Link %s only defined on one node", k)
			return errors.New(msg)
		} else if v > 2 {
			msg := fmt.Sprintf("Link %s only defined on one node", k)
			return errors.New(msg)
		} else if v != 2 {
			return errors.New("Link defined not exactly twice?")
		}
	}
	return nil
}
