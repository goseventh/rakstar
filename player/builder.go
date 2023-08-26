package player

type PlayerBuilder struct {
	ID          int
	Name        string
	ListWeapons []Weapon
	State
}

type State struct {
	Life       float32
	Armour     float32
	Coordinate []float32
	Character  int
}

func Builder() *PlayerBuilder {
	pb := new(PlayerBuilder)
	pb.ID = -1
	return pb
}
