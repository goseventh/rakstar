package server

import (
	"github.com/goseventh/rakstar/goroutines"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

func (ServerBuild) InAllPlayers(task func(id int)) {
	goroutines.Submit(func() {
		for i := 0; i < playerConst.MaxPlayers; i++ {
			task(i)
		}
	})
}
