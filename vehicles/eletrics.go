package vehicle

import "github.com/goseventh/rakstar/internal/natives"
import "math/rand"
import "time"

/*
Alterar a carga da bateria: valores 0~100
*/
func (e *eletricsBuilder) BatteryCharger(charger float32) *eletricsBuilder {
	if charger > 100 {
		charger = 100
	} else if charger < 0 {
		charger = 0
	}

	e.v.eletrics.batteryCharger = charger
	return e
}

func (e *eletricsBuilder) GetBatteryCharger() float32 {
	return e.v.eletrics.batteryCharger
}

/*
Altera o estado do faról do veículo

- Se estiver ligado, o faról desligará

- Se estiver desligado, o faról ligará
*/
func (e *eletricsBuilder) ToggleLights() *eletricsBuilder {
	var (
		engine    int
		lights    int
		alarm     int
		doors     int
		bonnet    int
		boot      int
		objective int
	)
	natives.GetVehicleParamsEx(
		e.v.id,
		&engine,
		&lights,
		&alarm,
		&doors,
		&bonnet,
		&boot,
		&objective,
	)
	if lights == 1 {
		lights = 0
	} else if lights == 0 {
		lights = 1
	}
	natives.SetVehicleParamsEx(
		e.v.id,
		engine,
		lights,
		alarm,
		doors,
		bonnet,
		boot,
		objective,
	)
	return e
}

// Quando invocada verificará se o veículo tem carga
// suficiente na bateria: <25% acontecerá uma
// pane em todo o sistema elétrico.
// Se tudo ocorrer bem, a função dormirá entre 7-3
// segundos para simular o dreno da bateria.
//
// - Utilize essa função para simular uma ignição
func (e *eletricsBuilder) IntroduceElectricalDrain() {
	var lights int
	natives.GetVehicleParamsEx(
		e.v.id,
		nil,
		&lights,
		nil,
		nil,
		nil,
		nil,
		nil,
	)
	if lights != 1 {
		return
	}
	if e.v.eletrics.batteryCharger > 25 {
		return
	}
	state := rand.Intn(1 - (-1) + 1)
	rounds := rand.Intn(150 - (-30) + 150)
	if state != 0 {
		time.Sleep(time.Duration(rand.Intn(7000 - 3000 + 7000)))
		return
	}
	for i := 0; i < rounds; i++ {
		time.Sleep(time.Duration(rand.Intn(1000 - 100 + 1000)))
		e.v.Eletrics().ToggleLights()
	}
}
