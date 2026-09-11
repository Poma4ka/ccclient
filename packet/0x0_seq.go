package packet

import (
	"encoding/binary"
)

type SeqGenericPacket struct {
	SeqID SeqID
}

func (p *SeqGenericPacket) SetSeqId(seqId SeqID) {
	p.SeqID = seqId
}

func (p *SeqGenericPacket) Unmarshal(data []byte) error {
	if len(data) < 4 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[:4]))
	return nil
}

func (p *SeqGenericPacket) Marshal() ([]byte, error) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(p.SeqID))
	return buf, nil
}
