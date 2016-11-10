package netsim

type Simulation struct {
	network *Network
	packets []*Packet
}

func NewSimulation(config string, packets string) (*Simulation, error) {
	c, err := LoadConfig(config)
	if err != nil {
		return nil, err
	}

	p, err := LoadPackets(packets)
	if err != nil {
		return nil, err
	}

	n, err := NewNetwork(c)
	if err != nil {
		return nil, err
	}

	return &Simulation{n, p}, nil
}
