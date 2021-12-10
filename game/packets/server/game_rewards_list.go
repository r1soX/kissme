package server

import (
	"io"

	"github.com/r1soX/kissme/game/packets/decode"
)

// GAME_REWARDS(13);
//[id:I, count:I, json:S]

// GameRewardsListServerPacket ...
type GameRewardsListServerPacket struct {
	El    int16
	ID    int32
	Count int32
	Json  string
}

// Parse ...
func (pack *GameRewardsListServerPacket) Parse(buffer io.Reader) error {
	return decode.Fill(pack, buffer)
}
