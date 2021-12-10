package server

import (
	"io"

	"github.com/r1soX/kissme/game/packets/decode"
)

// PLAYERS_KISSES(218); [player_id:I, kisses:I]

type KissesServerPacket struct {
	InnerID int32
	Kisses  int32
}

func (pack *KissesServerPacket) Parse(buffer io.Reader) error {
	return decode.Fill(pack, buffer)
}
