package player

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

func (pb *PlayerBuilder) Select(arg interface{}) *PlayerBuilder {
	switch v := arg.(type) {
	case string:
		var name string
		for i := 0; i < playerConst.MaxPlayers; i++ {
			natives.GetPlayerName(i, &name, playerConst.MaxPlayerName)
			if name == v {
				pb.ID = i
			}
			return pb

		}
	case int:
		pb.ID = v
		return pb

	}

	return pb
}
