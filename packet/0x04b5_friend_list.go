package packet

import (
	"bytes"
	"encoding/binary"
	"log/slog"
)

type FriendListPacket struct {
	FriendCount uint32
	Friends     []FriendInfo
}

type FriendInfo struct {
	PlayerID   uint64
	Might      uint32
	PlayerName string
	GuildName  string
}

func (p *FriendListPacket) GetPacketID() ID { return 0x04b5 }

func (p *FriendListPacket) Marshal() ([]byte, error) {
	return nil, nil
}

func (p *FriendListPacket) Unmarshal(data []byte) error {
	if len(data) < 12 {
		return ErrDataTooShort
	}

	totalPayloadSize := len(data) - 12
	p.FriendCount = uint32(totalPayloadSize / 88)

	if p.FriendCount > 200 || p.FriendCount == 0 {
		p.FriendCount = 10 // Защита
	}

	p.Friends = make([]FriendInfo, 0, p.FriendCount)

	offset := 12
	itemSize := 88

	trimNull := func(b []byte) string {
		if idx := bytes.IndexByte(b, 0); idx != -1 {
			return string(b[:idx])
		}
		return string(b)
	}

	for i := uint32(0); i < p.FriendCount; i++ {
		if offset+itemSize > len(data) {
			break
		}

		chunk := data[offset : offset+itemSize]

		friend := FriendInfo{
			PlayerName: trimNull(chunk[0:32]),
			GuildName:  trimNull(chunk[32:64]),
			PlayerID:   binary.LittleEndian.Uint64(chunk[64:72]),
			Might:      binary.LittleEndian.Uint32(chunk[72:76]),
		}

		p.Friends = append(p.Friends, friend)
		offset += itemSize
	}

	return nil
}

func (p *FriendListPacket) LogValue() slog.Value {
	return slog.Value{}
}
