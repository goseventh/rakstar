package player

import (
	"github.com/goseventh/rakstar/goroutines"
)

type PlayerBuilder struct {
	ID          int
	Name        string
	ListWeapons []Weapon
  wallet
}

func Builder() *PlayerBuilder {
	pb := new(PlayerBuilder)
	pb.ID = -1
	goroutines.Submit(
		func() {
			pb.updateMoney(pb.balance)
		})
	return pb
}
