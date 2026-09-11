package packet

import "log/slog"

type PingPacket struct {
	SeqGenericPacket
}

func (p *PingPacket) GetPacketID() ID { return 0x03eb }

func (p *PingPacket) Marshal() ([]byte, error) {
	return p.SeqGenericPacket.Marshal()
}

func (p *PingPacket) Unmarshal(data []byte) error {
	return p.SeqGenericPacket.Unmarshal(data)
}

func (p *PingPacket) LogValue() slog.Value {
	return slog.Value{}
}
