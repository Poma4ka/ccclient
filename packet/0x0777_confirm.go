package packet

import (
	"encoding/binary"
	"log/slog"
)

type ConfirmPacket struct {
	SeqGenericPacket

	ActionCode uint32
}

func (p *ConfirmPacket) GetPacketID() ID { return 0x0777 }

func (p *ConfirmPacket) Marshal() ([]byte, error) {
	data := make([]byte, 8)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint32(data[4:8], p.ActionCode)

	return data, nil
}

func (p *ConfirmPacket) Unmarshal(data []byte) error {
	if len(data) < 8 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.ActionCode = binary.LittleEndian.Uint32(data[4:8])

	return nil
}

func (p *ConfirmPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("action_code", uint64(p.ActionCode)),
	)
}
