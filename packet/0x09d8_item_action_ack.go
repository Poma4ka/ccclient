package packet

import (
	"encoding/binary"
	"log/slog"

	"github.com/Poma4ka/ccclient/item"
)

// ItemActionAckPacket - сервер -> клиент, подтверждение действия над
// предметом (продажа или использование). Приходит после QuickAckPacket
// (0x03ee) и перед ItemBalancePacket (0x09c5).
//
// Формат подтверждён по game.log, 6 байт: ItemID(2) + Count(2) +
// ActionType(2):
//
//	продажа предмета (SellItemPacket, 0x09c9):
//	  3c 00 01 00 01 00 -> item=0x003c count=1  action_type=1
//	  06 00 0c 00 01 00 -> item=0x0006 count=12 action_type=1
//
//	использование предмета (UseItemPacket, 0x09c8, action_type=3):
//	  09 00 01 00 03 00 -> item=0x0009 count=1  action_type=3
//	  09 00 05 00 03 00 -> item=0x0009 count=5  action_type=3
//	  0c 00 30 00 03 00 -> item=0x000c count=48 action_type=3
//
// ActionType здесь совпадает с типом операции: 1 для продажи, 3 для
// использования предмета (то же значение, что в UseItemPacket.ActionType).
// Count равен количеству предметов, затронутых операцией (продано либо
// использовано), не путать с остатком в ItemBalancePacket.
type ItemActionAckPacket struct {
	ItemID     item.Item
	Count      uint16
	ActionType uint16
}

func (p *ItemActionAckPacket) GetPacketID() ID { return 0x09d8 }

func (p *ItemActionAckPacket) Marshal() ([]byte, error) {
	data := make([]byte, 6)

	binary.LittleEndian.PutUint16(data[0:2], uint16(p.ItemID))
	binary.LittleEndian.PutUint16(data[2:4], p.Count)
	binary.LittleEndian.PutUint16(data[4:6], p.ActionType)

	return data, nil
}

func (p *ItemActionAckPacket) Unmarshal(data []byte) error {
	if len(data) < 6 {
		return ErrDataTooShort
	}

	p.ItemID = item.Item(binary.LittleEndian.Uint16(data[0:2]))
	p.Count = binary.LittleEndian.Uint16(data[2:4])
	p.ActionType = binary.LittleEndian.Uint16(data[4:6])

	return nil
}

func (p *ItemActionAckPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("item_id", p.ItemID.String()),
		slog.Int("count", int(p.Count)),
		slog.Uint64("action_type", uint64(p.ActionType)),
	)
}
