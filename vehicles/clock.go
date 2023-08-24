package vehicles

import (
	"math/rand"
	"time"

	"github.com/goseventh/rakstar/goroutines"
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
	goroutines.Builder().
		Submit(
			func() {
				for {
					<-ticker.C
					v.Eletrics().IntroduceElectricalDrain()
					// verifyBattery(v)
					// verifyFuel(v)
					v.engine.fuel -= (0.001 - v.engine.fuelEconomy)
					v.eletrics.batteryCharger -= 0.001
				}

			},
		)
	return v
}

func verifyFuel(v *vehicleBuilder) {
	if v.engine.fuel > 25 {
		return
	}
	state := rand.Intn(1 - (-1) + 1)
	if state != 0 {
		return
	}
	v.Engine().TurnOff()

}

// func verifyBattery(v *vehicleBuilder) {
// 	server.
// 		Builder().
// 		Goroutine().Submit(
// 		func() {
// 			var lights int
// 			natives.GetVehicleParamsEx(
// 				v.id,
// 				nil,
// 				&lights,
// 				nil,
// 				nil,
// 				nil,
// 				nil,
// 				nil,
// 			)
// 			if lights != 1 {
// 				return
// 			}
// 			if v.eletrics.batteryCharger > 25 {
// 				return
// 			}
// 			state := rand.Intn(1 - (-1) + 1)
// 			rounds := rand.Intn(100 - (-30) + 100)
// 			if state != 0 {
// 				return
// 			}
// 			for i := 0; i < rounds; i++ {
// 				time.Sleep(time.Duration(rand.Intn(1000 - 100 + 1000)))
// 				v.Eletrics().ToggleLights()
// 			}
// 		})
// }
