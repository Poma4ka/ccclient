package packet

import (
	"encoding/binary"
	"log/slog"
)

// BuildingIDAssignedPacket - server->client: присваивает клиенту постоянный
// building_id для только что построенного здания. Приходит сразу после
// BuildBuildingAckPacket. Значение этого поля затем используется как
// BuildingID во всех последующих MoveBuildingPacket для этого объекта.
type BuildingIDAssignedPacket struct {
	BuildingID uint32
}

func (p *BuildingIDAssignedPacket) GetPacketID() ID { return 0x0518 }

func (p *BuildingIDAssignedPacket) Marshal() ([]byte, error) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data[0:4], p.BuildingID)
	return data, nil
}

func (p *BuildingIDAssignedPacket) Unmarshal(data []byte) error {
	if len(data) < 4 {
		return ErrDataTooShort
	}
	p.BuildingID = binary.LittleEndian.Uint32(data[0:4])
	return nil
}

func (p *BuildingIDAssignedPacket) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Uint64("building_id", uint64(p.BuildingID)),
	)
}
