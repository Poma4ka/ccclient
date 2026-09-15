package packet

import (
	"encoding/binary"
	"encoding/hex"
	"log/slog"
)

type BuildBuildingAckPacket struct {
	Unknown    [8]byte
	EntityGUID [8]byte
	PosY       uint16
	PosX       uint16
}

func (p *BuildBuildingAckPacket) GetPacketID() ID { return 0x0515 }

func (p *BuildBuildingAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 20)

	copy(data[0:8], p.Unknown[:])
	copy(data[8:16], p.EntityGUID[:])
	binary.LittleEndian.PutUint16(data[16:18], p.PosY)
	binary.LittleEndian.PutUint16(data[18:20], p.PosX)

	return data, nil
}

func (p *BuildBuildingAckPacket) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return ErrDataTooShort
	}

	copy(p.Unknown[:], data[0:8])
	copy(p.EntityGUID[:], data[8:16])
	p.PosY = binary.LittleEndian.Uint16(data[16:18])
	p.PosX = binary.LittleEndian.Uint16(data[18:20])

	return nil
}

func (p *BuildBuildingAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("transaction_hash", hex.EncodeToString(p.Unknown[:])),
		slog.String("entity_guid", hex.EncodeToString(p.EntityGUID[:])),
		slog.Int("pos_y", int(p.PosY)),
		slog.Int("pos_x", int(p.PosX)),
	)
}
