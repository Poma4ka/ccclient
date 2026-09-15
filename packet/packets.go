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
	(*MoveBuildingPacket)(nil).GetPacketID():              func() Packet { return &MoveBuildingPacket{} },
	(*ConfirmPacket)(nil).GetPacketID():                   func() Packet { return &ConfirmPacket{} },
	(*MoveBuildingAckPacket)(nil).GetPacketID():           func() Packet { return &MoveBuildingAckPacket{} },
	(*BuildBuildingPacket)(nil).GetPacketID():             func() Packet { return &BuildBuildingPacket{} },
	(*QuickAckPacket)(nil).GetPacketID():                  func() Packet { return &QuickAckPacket{} },
	(*BuildBuildingAckPacket)(nil).GetPacketID():          func() Packet { return &BuildBuildingAckPacket{} },
	(*BuildingIDAssignedPacket)(nil).GetPacketID():        func() Packet { return &BuildingIDAssignedPacket{} },
	(*OpenInterfacePacket)(nil).GetPacketID():             func() Packet { return &OpenInterfacePacket{} },
	(*UseItemPacket)(nil).GetPacketID():                   func() Packet { return &UseItemPacket{} },
	(*SellItemPacket)(nil).GetPacketID():                  func() Packet { return &SellItemPacket{} },
	(*ItemBalancePacket)(nil).GetPacketID():               func() Packet { return &ItemBalancePacket{} },
	(*ItemActionAckPacket)(nil).GetPacketID():             func() Packet { return &ItemActionAckPacket{} },
	(*TalentRollAckPacket)(nil).GetPacketID():             func() Packet { return &TalentRollAckPacket{} },
	(*CollectResourcePacket)(nil).GetPacketID():           func() Packet { return &CollectResourcePacket{} },
	(*CollectResourceAckPacket)(nil).GetPacketID():        func() Packet { return &CollectResourceAckPacket{} },
}
