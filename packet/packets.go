package packet

// Packets - map of all known packets constructors
var Packets = map[ID]func() Packet{
	(*ConnectLoginServerPacket)(nil).GetPacketID():        func() Packet { return &ConnectLoginServerPacket{} },
	(*ConnectLoginServerSuccessPacket)(nil).GetPacketID(): func() Packet { return &ConnectLoginServerSuccessPacket{} },
	(*ConnectGameServerPacket)(nil).GetPacketID():         func() Packet { return &ConnectGameServerPacket{} },
	(*ConnectGameServerSuccessPacket)(nil).GetPacketID():  func() Packet { return &ConnectGameServerSuccessPacket{} },
	(*AckEncryptionPacket)(nil).GetPacketID():             func() Packet { return &AckEncryptionPacket{} },
	(*PingPacket)(nil).GetPacketID():                      func() Packet { return &PingPacket{} },
	(*PongPacket)(nil).GetPacketID():                      func() Packet { return &PongPacket{} },
	(*FriendListPacket)(nil).GetPacketID():                func() Packet { return &FriendListPacket{} },
}
