package server

import "github.com/goseventh/rakstar/internal/utils/constants/playerConst"

func (gb *Goroutine) Submit(task func()) {
	pool.Submit(task)
}

func (gb *Goroutine) InAllPlayers(task func(id int)) {
	pool.Submit(func() {
		for i := 0; i < playerConst.MaxPlayers; i++ {
			task(i)
		}
	})

}
