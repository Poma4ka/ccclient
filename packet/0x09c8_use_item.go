package packet

import (
	"encoding/binary"
	"log/slog"

	"github.com/Poma4ka/ccclient/item"
)

// UseItemPacket - клиент -> сервер, использование предмета.
//
// Изначально считался пакетом переброса таланта героя (RollTalentPacket),
// но по game.log подтверждено, что это более общий пакет "использовать
// предмет", где ActionType определяет конкретное действие:
//
//	ActionType=0 - переброс таланта героя картой (TargetID = HeroID):
//	  22 00 00 00 b3 00 00 00 00 00 00 00 01 00 33 00 01 00 [16 x 00]
//	  -> target_id=0xb3(179) action_type=0 unknown2=1 item=0x0033 count=1
//	  Ответ сервера: 0x03ee, 0x078d (полное состояние талантов), 0x09c5,
//	  0x05f1 (хэш результата ролла).
//
//	ActionType=3 - использование предмета со склада, без привязки к герою:
//	  09 00 01 00 03 00 -> item_action_ack: item=0x0009 count=1 action_type=3
//	  target_id в этих случаях всегда совпадал с item_count (1, 5, 48) -
//	  назначение поля в этом режиме не установлено, возможно клиент просто
//	  дублирует количество при отсутствии цели.
//	  Ответ сервера: 0x03ee, 0x09d8 (ItemActionAckPacket), 0x09c5 -
//	  БЕЗ 0x078d и 0x05f1 (эти два специфичны именно для реролла таланта).
//
// Unknown2 во всех наблюдениях равен 1 - назначение не установлено.
type UseItemPacket struct {
	SeqGenericPacket

	TargetID   uint32
	ActionType uint32
	Unknown2   uint16
	ItemID     item.Item
	ItemCount  uint16
}

func (p *UseItemPacket) GetPacketID() ID { return 0x09c8 }

func (p *UseItemPacket) Marshal() ([]byte, error) {
	data := make([]byte, 34)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint32(data[4:8], p.TargetID)
	binary.LittleEndian.PutUint32(data[8:12], p.ActionType)
	binary.LittleEndian.PutUint16(data[12:14], p.Unknown2)
	binary.LittleEndian.PutUint16(data[14:16], uint16(p.ItemID))
	binary.LittleEndian.PutUint16(data[16:18], p.ItemCount)
	// data[18:34] - 16 нулевых байт (резерв, назначение не известно)

	return data, nil
}

func (p *UseItemPacket) Unmarshal(data []byte) error {
	if len(data) < 34 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.TargetID = binary.LittleEndian.Uint32(data[4:8])
	p.ActionType = binary.LittleEndian.Uint32(data[8:12])
	p.Unknown2 = binary.LittleEndian.Uint16(data[12:14])
	p.ItemID = item.Item(binary.LittleEndian.Uint16(data[14:16]))
	p.ItemCount = binary.LittleEndian.Uint16(data[16:18])

	return nil
}

func (p *UseItemPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("target_id", uint64(p.TargetID)),
		slog.Uint64("action_type", uint64(p.ActionType)),
		slog.Uint64("unknown2", uint64(p.Unknown2)),
		slog.String("item_id", p.ItemID.String()),
		slog.Int("item_count", int(p.ItemCount)),
	)
}
