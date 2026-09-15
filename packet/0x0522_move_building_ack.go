package packet

import (
	"encoding/binary"
	"encoding/hex"
	"log/slog"
)

type MoveBuildingAckPacket struct {
	EntityGUID [8]byte
	Nonce      [8]byte
	Status     uint32
}

func (p *MoveBuildingAckPacket) GetPacketID() ID { return 0x0522 }

func (p *MoveBuildingAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 20)

	copy(data[0:8], p.EntityGUID[:])
	copy(data[8:16], p.Nonce[:])
	binary.LittleEndian.PutUint32(data[16:20], p.Status)

	return data, nil
}

func (p *MoveBuildingAckPacket) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return ErrDataTooShort
	}

	copy(p.EntityGUID[:], data[0:8])
	copy(p.Nonce[:], data[8:16])
	p.Status = binary.LittleEndian.Uint32(data[16:20])

	return nil
}

func (p *MoveBuildingAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("entity_guid", hex.EncodeToString(p.EntityGUID[:])),
		slog.String("nonce", hex.EncodeToString(p.Nonce[:])),
		slog.Uint64("status", uint64(p.Status)),
	)
}
