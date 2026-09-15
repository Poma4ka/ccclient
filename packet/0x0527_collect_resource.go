package packet

import (
	"encoding/binary"
	"log/slog"
)

// CollectResourcePacket - клиент -> сервер, сбор накопленного ресурса
// (золото/мана) со здания-сборщика.
//
// Формат подтверждён по game.log: серия сборов маны и серия сборов золота
// с разных зданий, во всех случаях один и тот же формат SeqID+BuildingID:
//
//	6f 00 00 00 03 00 00 00 -> building_id=3
//	70 00 00 00 2c 00 00 00 -> building_id=44
//	75 00 00 00 02 00 00 00 -> building_id=2
//	76 00 00 00 29 00 00 00 -> building_id=41
//
// Тип ресурса (золото/мана) в самом пакете не передаётся - сервер
// определяет его по типу здания с данным BuildingID.
type CollectResourcePacket struct {
	SeqGenericPacket

	BuildingID uint32
}

func (p *CollectResourcePacket) GetPacketID() ID { return 0x0527 }

func (p *CollectResourcePacket) Marshal() ([]byte, error) {
	data := make([]byte, 8)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint32(data[4:8], p.BuildingID)

	return data, nil
}

func (p *CollectResourcePacket) Unmarshal(data []byte) error {
	if len(data) < 8 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.BuildingID = binary.LittleEndian.Uint32(data[4:8])

	return nil
}

func (p *CollectResourcePacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("building_id", uint64(p.BuildingID)),
	)
}
