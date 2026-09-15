package packet

import (
	"encoding/binary"
	"log/slog"
)

type BuildBuildingPacket struct {
	SeqGenericPacket

	BuildingTypeID uint32
	Unknown        int32
	PosY           uint16
	PosX           uint16
}

func (p *BuildBuildingPacket) GetPacketID() ID { return 0x0514 }

func (p *BuildBuildingPacket) Marshal() ([]byte, error) {
	data := make([]byte, 20)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint32(data[4:8], p.BuildingTypeID)
	binary.LittleEndian.PutUint32(data[8:12], uint32(p.Unknown))
	binary.LittleEndian.PutUint16(data[12:14], p.PosY)
	binary.LittleEndian.PutUint16(data[14:16], p.PosX)
	// data[16:20] - 0x00000000

	return data, nil
}

func (p *BuildBuildingPacket) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.BuildingTypeID = binary.LittleEndian.Uint32(data[4:8])
	p.Unknown = int32(binary.LittleEndian.Uint32(data[8:12]))
	p.PosY = binary.LittleEndian.Uint16(data[12:14])
	p.PosX = binary.LittleEndian.Uint16(data[14:16])

	return nil
}

func (p *BuildBuildingPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("building_type_id", uint64(p.BuildingTypeID)),
		slog.Int("unknown", int(p.Unknown)),
		slog.Int("pos_y", int(p.PosY)),
		slog.Int("pos_x", int(p.PosX)),
	)
}
