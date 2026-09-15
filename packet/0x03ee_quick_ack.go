package packet

import (
	"encoding/hex"
	"log/slog"
)

// QuickAckPacket - сервер -> клиент, универсальное лёгкое подтверждение
// клиентского действия. Изначально считался специфичным для построения
// зданий, но по game.log подтверждено, что он приходит после множества
// разных действий: постройка (0x0514), перемещение здания, сбор ресурса
// (0x0527), продажа предмета (0x09c9), использование предмета (0x09c8) -
// во всех случаях сразу после запроса приходит один и тот же 8-байтный
// хэш/nonce, за которым уже следует специфичный для действия пакет
// (0x09d8, 0x09c5, 0x0528 и т.д.). Похоже на общий чек-сумма/anti-cheat
// токен, не зависящий от типа действия.
type QuickAckPacket struct {
	Unknown [8]byte
}

func (p *QuickAckPacket) GetPacketID() ID { return 0x03ee }

func (p *QuickAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 8)
	copy(data, p.Unknown[:])
	return data, nil
}

func (p *QuickAckPacket) Unmarshal(data []byte) error {
	if len(data) < 8 {
		return ErrDataTooShort
	}
	copy(p.Unknown[:], data[0:8])
	return nil
}

func (p *QuickAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("data", hex.EncodeToString(p.Unknown[:])),
	)
}
