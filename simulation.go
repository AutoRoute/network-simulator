package netsim

type Simulation struct {
	network *Network
	packets []*Packet

	// Maximum number of
	// iterations
	max_itr int

	// Incoming channel of
	// delivered packets
	delivered <-chan *Packet
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

	d := make(chan *Packet)

	n, err := NewNetwork(c, d)
	if err != nil {
		return nil, err
	}

	return &Simulation{n, p.Packets, 0, d}, nil
}

func (s Simulation) Run() error {
	t := 0
	for {
		if t > s.packets[len(s.packets)-1].T {
			break
		}

		for _, p := range s.packets {
			if p.T != t {
				break
			}

			err := s.sendPacket(p)
			if err != nil {
				return err
			}
		}

		t++
	}
	return nil
}

func (s Simulation) sendPacket(p *Packet) error {
	n, err := s.network.getNode(p.Source)
	if err != nil {
		return err
	}

	return n.SendPacket(p)
}
