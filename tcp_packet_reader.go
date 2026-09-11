package ccclient

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"github.com/Poma4ka/ccclient/packet"
)

// ReadPacket read one CC packet
func ReadPacket(conn net.Conn) (packet.Packet, error) {
	pId, pData, err := ReadPacketRaw(conn)
	if err != nil {
		return nil, err
	}

	factory, ok := packet.Packets[packet.ID(pId)]
	if !ok {
		factory = func() packet.Packet {
			return &packet.DumbPacket{
				PacketID: packet.ID(pId),
			}
		}
	}

	p := factory()

	err = p.Unmarshal(pData)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ReadEncryptedPacket read one CC encrypted packet
func ReadEncryptedPacket(conn net.Conn, e Encrypter) (packet.Packet, error) {
	pId, pData, err := ReadPacketRaw(conn)
	if err != nil {
		return nil, err
	}

	factory, ok := packet.Packets[packet.ID(pId)]
	if !ok {
		factory = func() packet.Packet {
			return &packet.DumbPacket{
				PacketID: packet.ID(pId),
			}
		}
	}

	p := factory()

	err = p.Unmarshal(e.Decrypt(pData))
	if err != nil {
		return nil, err
	}

	return p, nil
}

// headerSize - default header size (packet size + packet ID)
const headerSize = 2 + 2

// ReadPacketRaw read one CC incoming packet and return raw bytes
func ReadPacketRaw(conn net.Conn) (uint16, []byte, error) {
	headerBuf := make([]byte, headerSize)

	if _, err := io.ReadFull(conn, headerBuf); err != nil {
		return 0, nil, err
	}

	packetSize := binary.LittleEndian.Uint16(headerBuf[0:2])
	packetId := binary.LittleEndian.Uint16(headerBuf[2:4])

	if packetSize < headerSize {
		return 0, nil, fmt.Errorf("received invalid packet size: %d", packetSize)
	}

	payloadSize := packetSize - headerSize
	if payloadSize <= 0 {
		return packetId, nil, nil
	}

	payloadBuf := make([]byte, payloadSize)

	if _, err := io.ReadFull(conn, payloadBuf); err != nil {
		return 0, nil, fmt.Errorf("failed to read complete payload for msg 0x%04x: %w", packetId, err)
	}

	return packetId, payloadBuf, nil
}

// ReadEncryptedPacketRaw read one CC encrypted packet and return raw bytes
func ReadEncryptedPacketRaw(conn net.Conn, e Encrypter) (uint16, []byte, error) {
	id, data, err := ReadPacketRaw(conn)
	if err != nil {
		return 0, nil, err
	}

	return id, e.Decrypt(data), nil
}
