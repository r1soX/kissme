package client

import "github.com/r1soX/kissme/game/packets/encode"

// "IIIIBI,SS",
// BUY(6); good_id:I, cost:I, target_id:I, data:I, price_type:B, count:I, hash:S, params: S

// PrizeClientPacket ...
type PrizeClientPacket struct {
	Type        int16
	DeviceType  byte
	Good_id     int
	Cost        int
	Target_id   int
	Data        int
	Price_type  byte
	Count       int
	Screen_type int
}

// NewPrizeClientPacket ...
func NewPrizeClientPacket(good_id, cost, target_id, data int, price_type byte, count int, screen_type int) encode.ClientPacket {
	return &PrizeClientPacket{
		Type:        6,
		DeviceType:  0x04,
		Good_id:     good_id,
		Cost:        cost,
		Target_id:   target_id,
		Data:        data,
		Price_type:  price_type,
		Count:       count,
		Screen_type: screen_type,
	}
}

// Bytes ...
func (p *PrizeClientPacket) Bytes() []byte {
	return encode.Load(p)
}
