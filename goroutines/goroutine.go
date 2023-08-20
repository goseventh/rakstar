package goroutines

import "github.com/panjf2000/ants"
import "time"
import "log"

var pool *ants.Pool

const defaultRuntimes = 1000

type Handler func()
type manangerGoroutine struct {
}

func Boot() error {
	var err error
	pool, err = ants.NewPool(defaultRuntimes, ants.WithOptions(ants.Options{
		PreAlloc:       false,
		Nonblocking:    false, //bloqueando pra testar
		ExpiryDuration: time.Duration(7) * time.Second,
		PanicHandler: func(i interface{}) {
			log.Println("[rakstar] a panic occurred in the server manager:", i)
		},
	}))
	return err
}

func Submit(handler Handler) {
	pool.Submit(handler)
}

func GetPool() *ants.Pool{
  return pool
}

func SubmitWithLoop(startLoop, endLoop int, workers int, handler Handler) {
	x := startLoop
	s := endLoop / workers
	y := s

	for g := 0; g < workers; g++ {
		time.Sleep(time.Second * 5)
		for i := x; i <= y; i++ {
			log.Printf("worker %v, iniciou com x:%v y:%v\n", g, x, y)
		}
		x = y + 1
		y += s
		// Submit(func() {
		// 	// fmt.Printf("worker %v - loop: (%v/%v): %v\n", g, x, y, i)
		// })
		Submit(handler)
	}
}
