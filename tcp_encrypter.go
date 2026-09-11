package ccclient

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

type Encrypter interface {
	Encrypt(data []byte) []byte
	Decrypt(encrypted []byte) []byte
}

func NewEncrypter(key []byte) (Encrypter, error) {
	if len(key) < 8 {
		return nil, fmt.Errorf("DES key must be at least 8 bytes, got %d", len(key))
	}

	block, err := des.NewCipher(key[:8])
	if err != nil {
		return nil, fmt.Errorf("failed to create DES cipher: %w", err)
	}

	return &encrypter{
		block: block,
	}, nil
}

type encrypter struct {
	block cipher.Block
}

func (e *encrypter) Encrypt(data []byte) []byte {
	ciphertext := make([]byte, len(data))
	blockSize := e.block.BlockSize()

	alignedLen := (len(data) / blockSize) * blockSize

	for i := 0; i < alignedLen; i += blockSize {
		e.block.Encrypt(ciphertext[i:i+blockSize], data[i:i+blockSize])
	}

	if alignedLen < len(data) {
		copy(ciphertext[alignedLen:], data[alignedLen:])
	}

	return ciphertext
}

func (e *encrypter) Decrypt(encrypted []byte) []byte {
	plaintext := make([]byte, len(encrypted))
	blockSize := e.block.BlockSize()

	alignedLen := (len(encrypted) / blockSize) * blockSize

	for i := 0; i < alignedLen; i += blockSize {
		e.block.Decrypt(plaintext[i:i+blockSize], encrypted[i:i+blockSize])
	}

	if alignedLen < len(encrypted) {
		copy(plaintext[alignedLen:], encrypted[alignedLen:])
	}

	return plaintext
}
