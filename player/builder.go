package player

type PlayerBuilder struct {
	ID          int
	Name        string
	State
}

type State struct {
  ListWeapons []Weapon //TODO: implementar auto-save/load-save
	Life       float32
	Armour     float32
	Coordinate []float32
	Character  int
}

func Builder() *PlayerBuilder {
	pb := new(PlayerBuilder)
	pb.ID = -1
  pb.loadStates()
	return pb
}
