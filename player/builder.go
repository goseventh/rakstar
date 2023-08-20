package player

type PlayerBuilder struct {
	ID   int
	Name string
	ListWeapons []Weapon
}

func Builder() *PlayerBuilder {
	pb := new(PlayerBuilder)
	pb.ID = -1
	return pb
}
