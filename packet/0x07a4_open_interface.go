package packet

import (
	"encoding/binary"
	"log/slog"
)

// OpenInterfacePacket - клиент -> сервер, уведомление об открытии экрана
// интерфейса.
//
// Формат подтверждён по game.log (2 сэмпла, оба перед открытием экрана
// талантов героя):
//
//	20 00 00 00 00 00 00 00 0f
//	5b 00 00 00 00 00 00 00 0f
//
// В обоих случаях Unknown=0, InterfaceID=0x0f. Других значений InterfaceID
// пока не зафиксировано - нужно повторить захват для других экранов
// (герои, инвентарь, гильдия, почта, арена, магазин и т.д.), чтобы
// сопоставить конкретные значения с конкретными интерфейсами.
type OpenInterfacePacket struct {
	SeqGenericPacket

	Unknown     uint32
	InterfaceID uint8
}

func (p *OpenInterfacePacket) GetPacketID() ID {
	return 0x07a4
}

func (p *OpenInterfacePacket) Marshal() ([]byte, error) {
	data := make([]byte, 9)

	binary.LittleEndian.PutUint32(data[0:4], p.SeqID.Raw())
	binary.LittleEndian.PutUint32(data[4:8], p.Unknown)
	data[8] = p.InterfaceID

	return data, nil
}

func (p *OpenInterfacePacket) Unmarshal(data []byte) error {
	if len(data) < 9 {
		return ErrDataTooShort
	}

	p.SeqID = SeqID(binary.LittleEndian.Uint32(data[0:4]))
	p.Unknown = binary.LittleEndian.Uint32(data[4:8])
	p.InterfaceID = data[8]

	return nil
}

func (p *OpenInterfacePacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("unknown", uint64(p.Unknown)),
		slog.Uint64("interface_id", uint64(p.InterfaceID)),
	)
}
