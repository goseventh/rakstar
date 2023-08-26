package player

import (
	"time"

	"github.com/goseventh/rakstar/goroutines"
	"github.com/goseventh/rakstar/internal/natives"
)

func (pb *PlayerBuilder) initCollectionState() {
	ticker := time.NewTicker(time.Second)
	goroutines.Builder().Submit(func() {
		for {
			<-ticker.C
			pb.syncStates()
		}
	})
}

func (pb *PlayerBuilder) syncStates() {
	syncStateLife(pb)
  syncStateArmour(pb)
  syncStateCharacter(pb)
	syncStateCoordinate(pb)
}

func syncStateCoordinate(pb *PlayerBuilder) {
	var posX, posY, posZ float32
	ok := natives.GetPlayerPos(pb.ID, &posX, &posY, &posZ)
	if !ok {
		return
	}
	pb.Coordinate = []float32{posX, posY, posZ}
}

func syncStateLife(pb *PlayerBuilder) {
	var life float32
	ok := natives.GetPlayerHealth(pb.ID, &life)
	if !ok {
		return
	}
	pb.State.Life = life
}

func syncStateArmour(pb *PlayerBuilder) {
	var armour float32
	ok := natives.GetPlayerArmour(pb.ID, &armour)
	if !ok {
		return
	}
	pb.State.Armour = armour
}

func syncStateCharacter(pb *PlayerBuilder) {
	character := natives.GetPlayerSkin(pb.ID)
  if character < 0 {
    return
  }
  pb.Character = character
}
