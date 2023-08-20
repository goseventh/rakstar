package server

import (
	"github.com/panjf2000/ants"
)

var pool *ants.Pool

const defaultRuntimes = 7777

type ServerBuild struct {
	msgRestart string
	msgLoop    string
	tag        string
	playerID   int
	message    string
}


func Boot() error {
  return nil
}

func Builder() *ServerBuild {
	b := new(ServerBuild)
	b.playerID = -1
	return b
}

