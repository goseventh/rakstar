package goroutines

import (
	"fmt"
	"github.com/panjf2000/ants"
	"time"

	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)


var pool *ants.Pool
const defaultRuntimes = 7777

// Submit envia uma tarefa para os workers
func (gb *goroutine) Submit(task func()) {
	pool.Submit(task)
}

// InAllPlayers invocará [Submit] passando como parâmetro task
// em um loop para todos os jogadores, de forma que task seja
// executada simultaneamente para todos os jogadores ao mesmo
// tempo
func (gb *goroutine) InAllPlayers(task func(id int)) {
	pool.Submit(func() {
		for i := 0; i < playerConst.MaxPlayers; i++ {
			task(i)
		}
	})

}

// Loop é uma função que permite o processamento paralelo de um bloco de código.
// Ela divide o trabalho de incremento entre várias goroutines, acelerando a execução do loop.
// A função recebe três argumentos: startLoop, endLoop e goroutines.
// startLoop é o valor inicial do loop.
// endLoop é o valor final do loop.
// goroutines é o número de goroutines que serão usadas para processar o loop.
// Cada goroutine processa uma parte do loop, começando em 'startLoop' e terminando em 'endLoop'.
// A parte do loop que cada goroutine processa é determinada dividindo o intervalo do loop pelo número de goroutines.
// Por exemplo, se startLoop é 0, endLoop é 100 e goroutines é 10, então cada goroutine processará 10 iterações do loop.
// A função usa um pool de goroutines para executar as tarefas.
// Cada tarefa é submetida ao pool usando a função Submit do pool.
//
// # ATENÇÃO: Esta função contém um bug
func (gb *goroutine) Loop(startLoop, endLoop, goroutines int) {
	x := startLoop
	s := endLoop / goroutines
	y := s

	for g := 0; g < goroutines; g++ {
		time.Sleep(time.Second * 5)
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
