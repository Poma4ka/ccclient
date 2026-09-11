package packet

import (
	"bytes"
	"encoding/binary"
	"log/slog"
	"net"
	"strconv"
)

type ConnectLoginServerSuccessPacket struct {
	XGSPort    uint16
	UserID     uint64
	XGSIp      string
	LoginKey   string
	GSHostname string
	GSPort     uint16
}

func (p *ConnectLoginServerSuccessPacket) GetPacketID() ID { return 0x01f8 }

func (p *ConnectLoginServerSuccessPacket) Marshal() ([]byte, error) {
	buf := make([]byte, 266)

	binary.LittleEndian.PutUint16(buf[0:2], p.XGSPort)

	binary.LittleEndian.PutUint64(buf[4:12], p.UserID)

	copy(buf[12:44], p.XGSIp)
	copy(buf[44:133], p.LoginKey)
	copy(buf[133:262], p.GSHostname)

	binary.LittleEndian.PutUint16(buf[262:264], p.GSPort)

	return buf, nil
}

func (p *ConnectLoginServerSuccessPacket) Unmarshal(data []byte) error {
	if len(data) < 266 {
		return ErrDataTooShort
	}

	p.XGSPort = binary.LittleEndian.Uint16(data[0:2])
	p.UserID = binary.LittleEndian.Uint64(data[4:12])

	trimNull := func(b []byte) string {
		if idx := bytes.IndexByte(b, 0); idx != -1 {
			return string(b[:idx])
		}
		return string(b)
	}

	p.XGSIp = trimNull(data[12:44])
	p.LoginKey = trimNull(data[44:133])
	p.GSHostname = trimNull(data[133:262])

	p.GSPort = binary.LittleEndian.Uint16(data[262:264])

	return nil
}

func (p *ConnectLoginServerSuccessPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("user_id", p.UserID),
		slog.String("login_key", "***"),
		slog.String("game_server_host", net.JoinHostPort(p.GSHostname, strconv.Itoa(int(p.GSPort)))),
		slog.String("game_server_addr", net.JoinHostPort(p.XGSIp, strconv.Itoa(int(p.XGSPort)))),
	)
}
