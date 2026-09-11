package ccclient

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/Poma4ka/ccclient/packet"
)

type TcpConn interface {
	Open() error
	Close() error

	EnableEncryption(e Encrypter)
	DisableEncryption()

	WritePacket(p packet.Packet) error
	ReadPacket() (packet.Packet, error)
}

func OpenTcp(address string) TcpConn {
	return &tcpConn{
		address: address,
	}
}

type tcpConn struct {
	conn net.Conn
	mu   sync.Mutex

	enc   Encrypter
	encMu sync.RWMutex

	wMu sync.Mutex

	address string
}

func (c *tcpConn) Open() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		_ = c.conn.Close()
	}

	var err error

	c.conn, err = net.Dial("tcp", c.address)
	if err != nil {
		return fmt.Errorf("tcp connect failed to %s: %w", c.address, err)
	}
	return nil
}

func (c *tcpConn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		err := c.conn.Close()
		c.conn = nil
		return err
	}
	return nil
}

func (c *tcpConn) EnableEncryption(e Encrypter) {
	c.encMu.Lock()
	defer c.encMu.Unlock()

	c.enc = e
}

func (c *tcpConn) DisableEncryption() {
	c.encMu.Lock()
	defer c.encMu.Unlock()

	c.enc = nil
}

func (c *tcpConn) WritePacket(p packet.Packet) error {
	c.wMu.Lock()
	defer c.wMu.Unlock()

	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()

	if conn == nil {
		return io.EOF
	}

	c.encMu.Lock()
	enc := c.enc
	c.encMu.Unlock()

	if enc == nil {
		return WritePacket(conn, p)
	}

	return WriteEncryptedPacket(conn, p, enc)
}

func (c *tcpConn) ReadPacket() (packet.Packet, error) {
	c.mu.Lock()
	conn := c.conn
	c.mu.Unlock()

	if conn == nil {
		return nil, io.EOF
	}

	c.encMu.Lock()
	enc := c.enc
	c.encMu.Unlock()

	if enc == nil {
		return ReadPacket(conn)
	}

	return ReadEncryptedPacket(conn, enc)
}
