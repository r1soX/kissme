package client

import "github.com/r1soX/kissme/game/packets/encode"

// BonusClientPacket ...
type BonusClientPacket struct {
	Type       int16
	DeviceType byte
}

// NewBonusClientPacket ...
func NewBonusClientPacket() encode.ClientPacket {
	return &BonusClientPacket{
		Type:       61,
		DeviceType: 0x04,
	}
}

// Bytes ...
func (p *BonusClientPacket) Bytes() []byte {
	return encode.Load(p)
}
