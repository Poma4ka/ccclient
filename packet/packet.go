package packet

import "fmt"

// Packet generic packet interface
type Packet interface {
	// GetPacketID returns the packet identificator
	GetPacketID() ID
	// Marshal encodes the packet to bytes
	Marshal() ([]byte, error)
	// Unmarshal decodes the packet from bytes
	Unmarshal(data []byte) error
}

// SeqPacket is interface for sequensed packet
// Before send packet we call this method with actual seq_id
// Sequensed packets increment seqId for each client2server request (only if write was success)
type SeqPacket interface {
	// SetSeqId - add seqId to packet, seqId starts from 0
	SetSeqId(seqId SeqID)
}

type ID uint16

func (p ID) String() string {
	return fmt.Sprintf("0x%04x", p.Raw())
}

func (p ID) Raw() uint16 {
	return uint16(p)
}

type SeqID uint32

func (p SeqID) String() string {
	return fmt.Sprintf("0x%08x", p.Raw())
}

func (p SeqID) Raw() uint32 {
	return uint32(p)
}
