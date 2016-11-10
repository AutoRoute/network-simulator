package netsim

import (
	"encoding/json"
	"io/ioutil"
)

// PacketJSON represents the JSON object of the packets to be
// simulated.
type PacketJSON struct {
	Packets []*Packet
}

// Packet represents a simulated AutoRoute packet. Each element
// in the "packets" JSON array will be unmarshaled into this.
type Packet struct {
	Source      NodeAddress
	Destination NodeAddress

	// T is the time-step when the packet will
	// be sent from the source node
	T int
}

// LoadPackets takes a string path to a JSON file containing
// the simulated packet data and returns a *PacketJSON. It also
// returns error if there was a file I/O or JSON decoding error.
func LoadPackets(filename string) (*PacketJSON, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var packets PacketJSON
	if err = json.Unmarshal(file, &packets); err != nil {
		return nil, err
	}

	return &packets, nil
}
