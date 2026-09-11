package packet

// Packets - map of all known packets constructors
var Packets = map[ID]func() Packet{
	(*PingPacket)(nil).GetPacketID(): func() Packet { return &PingPacket{} },
	(*PongPacket)(nil).GetPacketID(): func() Packet { return &PongPacket{} },
}
