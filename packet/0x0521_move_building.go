package packet

import (
	"encoding/binary"
	"log/slog"
)

type MoveBuildingPacket struct {
	SeqGenericPacket

	BuildingID uint32
	PosY       uint16
	PosX       uint16
}

func (p *MoveBuildingPacket) GetPacketID() ID { return 0x0521 }

func (p *MoveBuildingPacket) Marshal() ([]byte, error) {
	data := make([]byte, 24)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	// data[4:8]  - 0x00000000
	// data[8:12] - 0x00000000
	binary.LittleEndian.PutUint32(data[12:16], p.BuildingID)
	binary.LittleEndian.PutUint16(data[16:18], p.PosY)
	binary.LittleEndian.PutUint16(data[18:20], p.PosX)
	// data[20:24] - 0x00000000

	return data, nil
}

func (p *MoveBuildingPacket) Unmarshal(data []byte) error {
	if len(data) < 24 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.BuildingID = binary.LittleEndian.Uint32(data[12:16])
	p.PosY = binary.LittleEndian.Uint16(data[16:18])
	p.PosX = binary.LittleEndian.Uint16(data[18:20])

	return nil
}

func (p *MoveBuildingPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("building_id", uint64(p.BuildingID)),
		slog.Int("pos_y", int(p.PosY)),
		slog.Int("pos_x", int(p.PosX)),
	)
}
