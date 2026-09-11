package packet

import (
	"fmt"
	"log/slog"
)

type DumbPacket struct {
	PacketID ID
	Payload  []byte
}

func (p *DumbPacket) GetPacketID() ID { return p.PacketID }

func (p *DumbPacket) Unmarshal(data []byte) error {
	p.Payload = data
	return nil
}

func (p *DumbPacket) Marshal() ([]byte, error) {
	return p.Payload, nil
}

func (p *DumbPacket) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("% x", p.Payload))
}
