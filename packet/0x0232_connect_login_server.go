package packet

import (
	"bytes"
	"encoding/binary"
	"log/slog"
)

type ConnectLoginServerPacket struct {
	ClientVersion uint32
	UserID        uint64
	AuthKey       string
	GameID        uint32
}

func (p *ConnectLoginServerPacket) GetPacketID() ID { return 0x0232 }

func (p *ConnectLoginServerPacket) Marshal() ([]byte, error) {
	buf := make([]byte, 536)

	binary.LittleEndian.PutUint32(buf[0:4], p.ClientVersion)
	binary.LittleEndian.PutUint64(buf[4:12], p.UserID)
	copy(buf[12:524], p.AuthKey)
	binary.LittleEndian.PutUint32(buf[524:528], p.GameID)

	return buf, nil
}

func (p *ConnectLoginServerPacket) Unmarshal(data []byte) error {
	if len(data) < 536 {
		return ErrDataTooShort
	}

	p.ClientVersion = binary.LittleEndian.Uint32(data[0:4])
	p.UserID = binary.LittleEndian.Uint64(data[4:12])
	p.AuthKey = string(bytes.TrimRight(data[12:524], "\x00"))
	p.GameID = binary.LittleEndian.Uint32(data[524:528])

	return nil
}

func (p *ConnectLoginServerPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("user_id", p.UserID),
		slog.Uint64("game_id", uint64(p.GameID)),
		slog.String("auth_key", "***"),
		slog.Uint64("client_version", uint64(p.ClientVersion)),
	)
}
