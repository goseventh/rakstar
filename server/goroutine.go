package server

import (
	"fmt"
	"time"

	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

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

func (gb *Goroutine) Loop(startLoop, endLoop, goroutines int) {
	x := startLoop
	s := endLoop / goroutines
	y := s

	for g := 0; g < goroutines; g++ {
    time.Sleep(time.Second*5)
    for i := x; i <= y; i++ {
    fmt.Printf("worker %v, iniciou com x:%v y:%v\n", g, x, y)
    }
    x = y + 1
    y += s
		pool.Submit(func() {
				// fmt.Printf("worker %v - loop: (%v/%v): %v\n", g, x, y, i)
		})
	}
}
