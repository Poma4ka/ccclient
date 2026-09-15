package packet

import (
	"encoding/binary"
	"log/slog"
)

type PongPacket struct {
	Field1    uint32
	Field2    uint32
	Field3    uint32
	Timestamp uint32
	Field5    uint32
}

func (p *PongPacket) GetPacketID() ID { return 0x03ec }

func (p *PongPacket) Marshal() ([]byte, error) {
	buf := make([]byte, 20)

	binary.LittleEndian.PutUint32(buf[0:4], p.Field1)
	binary.LittleEndian.PutUint32(buf[4:8], p.Field2)
	binary.LittleEndian.PutUint32(buf[8:12], p.Field3)
	binary.LittleEndian.PutUint32(buf[12:16], p.Timestamp)
	binary.LittleEndian.PutUint32(buf[16:20], p.Field5)

	return buf, nil
}

func (p *PongPacket) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return ErrDataTooShort
	}

	p.Field1 = binary.LittleEndian.Uint32(data[0:4])
	p.Field2 = binary.LittleEndian.Uint32(data[4:8])
	p.Field3 = binary.LittleEndian.Uint32(data[8:12])
	p.Timestamp = binary.LittleEndian.Uint32(data[12:16])
	p.Field5 = binary.LittleEndian.Uint32(data[16:20])

	return nil
}

func (p *PongPacket) LogValue() slog.Value {
	return slog.Value{}
}
