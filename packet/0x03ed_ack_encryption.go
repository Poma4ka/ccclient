package packet

import (
	"encoding/binary"
	"log/slog"
)

type AckEncryptionPacket struct {
	SeqGenericPacket

	UserID         uint64
	PlatformLangID uint32
}

func (p *AckEncryptionPacket) GetPacketID() ID { return 0x03ed }

func (p *AckEncryptionPacket) Marshal() ([]byte, error) {
	data := make([]byte, 20)

	binary.LittleEndian.PutUint32(data[0:4], uint32(p.SeqID))
	binary.LittleEndian.PutUint64(data[4:12], p.UserID)
	// Unknown 4 bytes
	binary.LittleEndian.PutUint32(data[16:20], p.PlatformLangID)

	return data, nil
}

func (p *AckEncryptionPacket) Unmarshal(data []byte) error {
	if len(data) < 20 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.UserID = binary.LittleEndian.Uint64(data[4:12])
	// Unknown 4 bytes
	p.PlatformLangID = binary.LittleEndian.Uint32(data[16:20])

	return nil
}

func (p *AckEncryptionPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("user_id", p.UserID),
		slog.Uint64("lang_id", uint64(p.PlatformLangID)),
	)
}
