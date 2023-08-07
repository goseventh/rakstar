package player

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

func (pb *PlayerBuilder) Life(life float32) *PlayerBuilder {
	natives.SetPlayerHealth(pb.ID, life)
	return pb
}

func (pb *PlayerBuilder) Armour(Armour float32) *PlayerBuilder {
	natives.SetPlayerArmour(pb.ID, Armour)
	return pb
}

func (pb *PlayerBuilder) Spawn() *PlayerBuilder {
	natives.SpawnPlayer(pb.ID)
	return pb
}

func (pb *PlayerBuilder) Nick(nick *string) *PlayerBuilder {
	natives.GetPlayerName(pb.ID, nick, playerConst.MaxPlayerName)
	return pb
}

func (pb *PlayerBuilder) GetAngle() float32 {
	var angle float32
	connected := natives.GetPlayerFacingAngle(pb.ID, &angle)
	if !connected {
		return -1
	}
	return angle
}

func (pb *PlayerBuilder) GetPos() (float32, float32, float32) {
	var (
		x, y, z float32
	)
	connected := natives.GetPlayerPos(pb.ID, &x, &y, &z)
	if !connected {
		return -1, -1, -1
	}
	return x, y, z
}


func (pb *PlayerBuilder) InVehicle() bool {
  return natives.IsPlayerInAnyVehicle(pb.ID)
}


