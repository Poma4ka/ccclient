package packet

import (
	"encoding/hex"
	"log/slog"
)

// TalentRollAckPacket - сервер -> клиент, подтверждение результата
// переброса таланта героя. Приходит сразу после ItemBalancePacket (0x09c5)
// в ответ на UseItemPacket (0x09c8) с ActionType=0 (переброс таланта картой,
// TargetID=HeroID), а также на бесплатную ежедневную попытку (0x05f0).
// Для UseItemPacket с другими ActionType (например ActionType=3 -
// использование обычного предмета со склада) этот пакет не приходит -
// вместо него сервер шлёт ItemActionAckPacket (0x09d8).
//
// Все 16 байт устойчиво выглядят как непрозрачный хэш/идентификатор
// результата ролла (полностью совпадающие значения наблюдались у разных
// героев в разных сессиях), а не как читаемые числовые поля - разложить
// их на осмысленные под-поля по имеющимся данным не удалось. Payload
// хранится как есть.
type TalentRollAckPacket struct {
	Data [16]byte
}

func (p *TalentRollAckPacket) GetPacketID() ID { return 0x05f1 }

func (p *TalentRollAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 16)
	copy(data, p.Data[:])
	return data, nil
}

func (p *TalentRollAckPacket) Unmarshal(data []byte) error {
	if len(data) < 16 {
		return ErrDataTooShort
	}
	copy(p.Data[:], data[0:16])
	return nil
}

func (p *TalentRollAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("data", hex.EncodeToString(p.Data[:])),
	)
}
