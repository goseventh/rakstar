package goroutines

import "github.com/panjf2000/ants"
import "time"
import "log" 


type goroutine struct {
	runtimes int
}

func init() {
	var err error
	pool, err = ants.NewPool(defaultRuntimes, ants.WithOptions(ants.Options{
		PreAlloc:       false,
		Nonblocking:    false, //bloqueando pra testar
		ExpiryDuration: time.Duration(7) * time.Second,
		PanicHandler: func(i interface{}) {
			log.Println("[rakstar] a panic occurred in the server manager:", i)
		},
	}))
  panic (err)
}

func Builder() *goroutine {
	g := new(goroutine)
	return g
}
