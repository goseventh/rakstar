package player

import (
	"github.com/goseventh/rakstar/internal/samp"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

func (pb *PlayerBuilder) Life(life float32) *PlayerBuilder {
	samp.SetPlayerHealth(pb.ID, life)
	return pb
}

func (pb *PlayerBuilder) Armour(Armour float32) *PlayerBuilder {
	samp.SetPlayerArmour(pb.ID, Armour)
	return pb
}

func (pb *PlayerBuilder) Spawn() *PlayerBuilder {
	samp.SpawnPlayer(pb.ID)
	return pb
}

func (pb *PlayerBuilder) Nick(nick *string) *PlayerBuilder {
	samp.GetPlayerName(pb.ID, nick, playerConst.MaxPlayerName)
	return pb
}
