package packet

import (
	"encoding/binary"
	"log/slog"
)

type ConnectGameServerSuccessPacket struct {
	UserID    uint64
	CryptoKey []byte
}

func (p *ConnectGameServerSuccessPacket) GetPacketID() ID { return 0x01f9 }

func (p *ConnectGameServerSuccessPacket) Marshal() ([]byte, error) {
	data := make([]byte, 28)

	binary.LittleEndian.PutUint64(data[4:12], p.UserID)

	if len(p.CryptoKey) > 0 {
		copy(data[12:28], p.CryptoKey)
	}

	return data, nil
}

func (p *ConnectGameServerSuccessPacket) Unmarshal(data []byte) error {
	if len(data) < 28 {
		return ErrDataTooShort
	}

	p.UserID = binary.LittleEndian.Uint64(data[4:12])

	p.CryptoKey = make([]byte, 16)
	copy(p.CryptoKey, data[12:28])

	return nil
}

func (p *ConnectGameServerSuccessPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("user_id", p.UserID),
		slog.String("crypto_key", "***"),
	)
}
