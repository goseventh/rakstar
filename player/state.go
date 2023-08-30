package player

import (
	"time"

	"github.com/goseventh/rakstar/database/mongo"
	"github.com/goseventh/rakstar/goroutines"
	"github.com/goseventh/rakstar/internal/natives"
	"go.mongodb.org/mongo-driver/bson"
)

const playerStateColl string = "player_states"

func (pb *PlayerBuilder) initCollectionState() {
	ticker := time.NewTicker(time.Second)
	goroutines.Builder().Submit(func() {
		for {
			<-ticker.C
			pb.saveStates()
		}
	})
}

func (pb *PlayerBuilder) saveStates() {
	getStateLife(pb)
	getStateArmour(pb)
	getStateCharacter(pb)
	getStateCoordinate(pb)
	mongodb.Builder().
		UseDatabase("rk_state").
		UseCollection(playerStateColl).
		CreatePlayerStateWithWorkers(pb.State)
}

func (pb *PlayerBuilder) loadStates() {
	state := new(PlayerBuilder)
	mongodb.Builder().
		UseDatabase("rk_state").
		UseCollection(playerStateColl).
		GetPlayerState(state, bson.D{})
	pb.State = state.State
	setStateLife(pb)
	setStateArmour(pb)
	setStateCharacter(pb)
	setStateCoordinate(pb)
}

func getStateCoordinate(pb *PlayerBuilder) {
	var posX, posY, posZ float32
	ok := natives.GetPlayerPos(pb.ID, &posX, &posY, &posZ)
	if !ok {
		return
	}
	pb.Coordinate = []float32{posX, posY, posZ}
}

func getStateLife(pb *PlayerBuilder) {
	var life float32
	ok := natives.GetPlayerHealth(pb.ID, &life)
	if !ok {
		return
	}
	pb.State.Life = life
}

func getStateArmour(pb *PlayerBuilder) {
	var armour float32
	ok := natives.GetPlayerArmour(pb.ID, &armour)
	if !ok {
		return
	}
	pb.State.Armour = armour
}

func getStateCharacter(pb *PlayerBuilder) {
	character := natives.GetPlayerSkin(pb.ID)
	if character < 0 {
		return
	}
	pb.Character = character
}

func setStateCoordinate(pb *PlayerBuilder) bool {
	if len(pb.Coordinate) < 3 {
		return false
	}
	ok := natives.SetPlayerPos(pb.ID, pb.Coordinate[0], pb.Coordinate[1], pb.Coordinate[2])
	return ok
}

func setStateLife(pb *PlayerBuilder) bool {
	ok := natives.SetPlayerHealth(pb.ID, pb.State.Life)
	return ok
}

func setStateArmour(pb *PlayerBuilder) bool {
	ok := natives.SetPlayerArmour(pb.ID, pb.State.Armour)
	return ok
}

func setStateCharacter(pb *PlayerBuilder) bool {
	ok := natives.SetPlayerSkin(pb.ID, pb.Character)
	return ok
}
