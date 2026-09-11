package ccclient

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/Poma4ka/ccclient/packet"
)

// WritePacket write packet to connection
func WritePacket(conn net.Conn, p packet.Packet) error {
	data, err := p.Marshal()
	if err != nil {
		return err
	}

	return WritePacketRaw(conn, uint16(p.GetPacketID()), data)
}

// WriteEncryptedPacket write encrypted packet to connection
func WriteEncryptedPacket(conn net.Conn, p packet.Packet, e Encrypter) error {
	data, err := p.Marshal()
	if err != nil {
		return err
	}

	return WritePacketRaw(conn, uint16(p.GetPacketID()), e.Encrypt(data))
}

// WriteEncryptedPacketRaw write raw packet to connection with encryption
func WriteEncryptedPacketRaw(conn net.Conn, msgID uint16, data []byte, e Encrypter) error {
	return WritePacketRaw(conn, msgID, e.Encrypt(data))
}

// WritePacketRaw write raw packet to connection
func WritePacketRaw(conn net.Conn, packetId uint16, data []byte) error {
	packetSize := uint16(len(data) + headerSize)
	buffer := make([]byte, packetSize)

	binary.LittleEndian.PutUint16(buffer[0:2], packetSize)
	binary.LittleEndian.PutUint16(buffer[2:4], packetId)
	copy(buffer[4:], data)

	_, err := conn.Write(buffer)
	if err != nil {
		return fmt.Errorf("tcp write failed: %w", err)
	}

	return nil
}
