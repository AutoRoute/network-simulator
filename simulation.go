package netsim

// Simulation keeps the state of the network
// and the packets sent/to be sent
type Simulation struct {
	network *Network
	packets []*Packet

	// Maximum number of
	// iterations
	maxItr int

	// Incoming channel of
	// delivered packets
	delivered <-chan *Packet
}

// NewSimulation takes a string path to a network configuration JSON
// file and a simulated packets JSON file and returns a *Simulation
// and error
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

// Run starts the simulation and blocks until all
// iterations are done or if there was an error.
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

			err := s.addPacket(p)
			if err != nil {
				return err
			}
		}

		t++
	}
	return nil
}

func (s Simulation) addPacket(p *Packet) error {
	n, err := s.network.getNode(p.Source)
	if err != nil {
		return err
	}

	n.Packets <- p
	return nil
}
