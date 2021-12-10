package server

import (
	"io"

	"github.com/r1soX/kissme/game/packets/decode"
)

type BonusServerPacket struct {
	CanCollect byte
	Day        byte
}

func (pack *BonusServerPacket) Parse(buffer io.Reader) error {
	return decode.Fill(pack, buffer)
}
