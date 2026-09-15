package packet

import (
	"encoding/binary"
	"log/slog"

	"github.com/Poma4ka/ccclient/item"
)

// ItemBalancePacket - сервер -> клиент, уведомление о новом количестве
// предмета в инвентаре после операции, которая его расходует/тратит
// (продажа предмета - SellItemPacket 0x09c9, либо использование предмета -
// UseItemPacket 0x09c8, включая переброс таланта картой).
//
// Формат подтверждён по game.log:
//
//	продажа предметов (0x09c9):
//	  0c 00 30 00 00 00 -> item=0x000c count=48
//	  11 00 fa 02 00 00 -> item=0x0011 count=762
//	  3f 00 d0 1a 00 00 -> item=0x003f count=6864
//	  46 00 00 00 00 00 -> item=0x0046 count=0
//
//	использование карты таланта (0x09c8, action_type=0), 4 последовательных
//	ролла на разных героях:
//	  33 00 98 00 00 00 -> item=0x0033 count=152
//	  33 00 97 00 00 00 -> item=0x0033 count=151
//	  33 00 96 00 00 00 -> item=0x0033 count=150
//	  33 00 95 00 00 00 -> item=0x0033 count=149
//
// Count каждый раз уменьшается ровно на количество потраченных предметов -
// это остаток данного предмета в инвентаре после операции.
type ItemBalancePacket struct {
	ItemID item.Item
	Count  uint32
}

func (p *ItemBalancePacket) GetPacketID() ID { return 0x09c5 }

func (p *ItemBalancePacket) Marshal() ([]byte, error) {
	data := make([]byte, 6)

	binary.LittleEndian.PutUint16(data[0:2], uint16(p.ItemID))
	binary.LittleEndian.PutUint32(data[2:6], p.Count)

	return data, nil
}

func (p *ItemBalancePacket) Unmarshal(data []byte) error {
	if len(data) < 6 {
		return ErrDataTooShort
	}

	p.ItemID = item.Item(binary.LittleEndian.Uint16(data[0:2]))
	p.Count = binary.LittleEndian.Uint32(data[2:6])

	return nil
}

func (p *ItemBalancePacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("item_id", p.ItemID.String()),
		slog.Uint64("count", uint64(p.Count)),
	)
}
