package packet

import (
	"encoding/binary"
	"encoding/hex"
	"log/slog"
)

// CollectResourceAckPacket - сервер -> клиент, подтверждение сбора ресурса
// со здания (ответ на CollectResourcePacket, 0x0527).
//
// Формат подтверждён по game.log (12 байт: 8 байт хэш/nonce + uint32):
//
//	1c b2 03 01 d0 6e 86 b4 8c 25 05 00 -> amount=338316
//	48 1d d8 7c f0 3f 40 19 6f 84 05 00 -> amount=361071
//	28 c8 4c 26 63 ed 7e a3 fa d4 05 00 -> amount=381178
//
// Amount похоже на текущий общий запас ресурса (золото/мана) в кошельке
// игрока после сбора: значение растёт от пакета к пакету, но при сборе
// сразу нескольких зданий подряд иногда повторяется у двух ответов подряд -
// возможно, сервер отдаёт актуальный на момент ответа общий баланс, а не
// именно "добавлено с этого здания". Первые 8 байт - хэш/nonce, как и в
// других *AckPacket (см. QuickAckPacket, MoveBuildingAckPacket),
// назначение не установлено.
type CollectResourceAckPacket struct {
	Hash   [8]byte
	Amount uint32
}

func (p *CollectResourceAckPacket) GetPacketID() ID { return 0x0528 }

func (p *CollectResourceAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 12)

	copy(data[0:8], p.Hash[:])
	binary.LittleEndian.PutUint32(data[8:12], p.Amount)

	return data, nil
}

func (p *CollectResourceAckPacket) Unmarshal(data []byte) error {
	if len(data) < 12 {
		return ErrDataTooShort
	}

	copy(p.Hash[:], data[0:8])
	p.Amount = binary.LittleEndian.Uint32(data[8:12])

	return nil
}

func (p *CollectResourceAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("hash", hex.EncodeToString(p.Hash[:])),
		slog.Uint64("amount", uint64(p.Amount)),
	)
}
