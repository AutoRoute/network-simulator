package netsim

import (
	"testing"
)

func TestLoadPackets(t *testing.T) {
	packets, err := LoadPackets("examples/test_packets.json")
	if err != nil {
		t.Fatal(err)
	}

	if packets.Packets[0].Source != "A" ||
		packets.Packets[0].Destination != "B" ||
		packets.Packets[0].T != 0 ||
		packets.Packets[1].Source != "A" ||
		packets.Packets[1].Destination != "C" ||
		packets.Packets[1].T != 1 {
		t.Fatal("Packets not loaded correctly")
	}
}
