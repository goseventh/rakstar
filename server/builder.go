package server

import (
	"log"
	"time"

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

type Goroutine struct {
	runtimes int
}

func Boot() error {
	var err error
	pool, err = ants.NewPool(defaultRuntimes, ants.WithOptions(ants.Options{
		PreAlloc:       false,
		Nonblocking:   false , //bloqueando pra testar
		ExpiryDuration: time.Duration(7) * time.Second,
		PanicHandler: func(i interface{}) {
			log.Println("[rakstar] a panic occurred in the server manager.")
		},
	}))
	return err
}

func Builder() *ServerBuild {
	b := new(ServerBuild)
	b.playerID = -1
	return b
}

func (rb *ServerBuild) Goroutine() *Goroutine {
	gb := new(Goroutine)
	gb.runtimes = defaultRuntimes
	return gb
}
