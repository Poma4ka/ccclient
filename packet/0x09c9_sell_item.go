package packet

import (
	"encoding/binary"
	"log/slog"

	"github.com/Poma4ka/ccclient/item"
)

// SellItemPacket - клиент -> сервер, продажа предмета из инвентаря.
//
// Формат подтверждён по game.log (продажа 4 разных предметов, разное количество):
//
//	1a 00 00 00 0c 00 01 00 -> item=0x000c (ItemHonorBadgePack3) count=1
//	1b 00 00 00 11 00 01 00 -> item=0x0011 (ItemWorkHammer3)     count=1
//	1c 00 00 00 3f 00 16 00 -> item=0x003f (ItemManaPack3)       count=22
//	1d 00 00 00 46 00 09 00 -> item=0x0046 (ItemPileOfGems)      count=9
type SellItemPacket struct {
	SeqGenericPacket

	ItemID item.Item
	Count  uint16
}

func (p *SellItemPacket) GetPacketID() ID { return 0x09c9 }

func (p *SellItemPacket) Marshal() ([]byte, error) {
	data := make([]byte, 8)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint16(data[4:6], uint16(p.ItemID))
	binary.LittleEndian.PutUint16(data[6:8], p.Count)

	return data, nil
}

func (p *SellItemPacket) Unmarshal(data []byte) error {
	if len(data) < 8 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.ItemID = item.Item(binary.LittleEndian.Uint16(data[4:6]))
	p.Count = binary.LittleEndian.Uint16(data[6:8])

	return nil
}

func (p *SellItemPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("item_id", p.ItemID.String()),
		slog.Int("count", int(p.Count)),
	)
}
