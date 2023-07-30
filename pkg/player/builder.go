package player
type PlayerBuilder struct {
	ID int
	Name string

}

func Builder() *PlayerBuilder{
	pb := new(PlayerBuilder)
	pb.ID = -1
	return pb
}

