package packet

import (
	"bytes"
	"encoding/binary"
	"log/slog"
)

type ConnectGameServerPacket struct {
	SeqGenericPacket

	UserID        uint64
	LoginKey      string
	ClientSign    uint32
	ClientVersion uint32
}

func (p *ConnectGameServerPacket) GetPacketID() ID { return 0x01f7 }

func (p *ConnectGameServerPacket) Marshal() ([]byte, error) {
	buf := make([]byte, 176)

	binary.LittleEndian.PutUint32(buf[0:4], uint32(p.SeqID))
	binary.LittleEndian.PutUint64(buf[4:12], p.UserID)
	copy(buf[12:168], p.LoginKey)
	binary.LittleEndian.PutUint32(buf[168:172], p.ClientSign)
	binary.LittleEndian.PutUint32(buf[172:176], p.ClientVersion)

	return buf, nil
}

func (p *ConnectGameServerPacket) Unmarshal(data []byte) error {
	if len(data) < 176 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.UserID = binary.LittleEndian.Uint64(data[4:12])
	p.LoginKey = string(bytes.TrimRight(data[12:168], "\x00"))
	p.ClientSign = binary.LittleEndian.Uint32(data[168:172])
	p.ClientVersion = binary.LittleEndian.Uint32(data[172:176])

	return nil
}

func (p *ConnectGameServerPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("user_id", p.UserID),
		slog.String("login_key", "***"),
		slog.Uint64("client_version", uint64(p.ClientVersion)),
	)
}
