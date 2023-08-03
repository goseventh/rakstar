package vehicle

import (
	"math/rand"
	"time"

	"github.com/goseventh/rakstar/server"
)

/*
Inicia o clock do veículo
  - O clock do veículo realiza todo o processamento e computação
    necessários para que os sistemas elétricos, motor, e outros
    funcionem corretamente.

Importante: invocar mais de uma vez essa função na mesma instância, causará
comportamentos estranhos
*/
func (v *vehicleBuilder) Start() *vehicleBuilder {
	ticker := time.NewTicker(time.Second)
	server.Builder().
		Goroutine().Submit(
		func() {
			for {
				<-ticker.C
        verifyBattery(v)
        verifyFuel(v)
			}

		},
	)
	return v
}

func verifyFuel(v *vehicleBuilder) {
	if v.Engine().fuel > 25 {
		return
	}
	state := rand.Intn(1 - (-1) + 1)
	if state != 0 {
		return
	}
	v.Engine().TurnOff()

}
func verifyBattery(v *vehicleBuilder) {
	if v.Eletrics().batteryCharger > 25 {
		return
	}
	state := rand.Intn(1 - (-1) + 1)
	rounds := rand.Intn(100 - (-30) + 100)
	if state != 0 {
		return
	}
	for i := 0; i < rounds; i++ {
		time.Sleep(time.Duration(rand.Intn(1000 - 100 + 1000)))
		v.Eletrics().ToggleLights()
	}
}
