package vehicles

import (
	"math"
	"math/rand"
	// "time"

	"github.com/goseventh/rakstar/internal/natives"
)

/*
Altera a economia de combustível do veículo: valores 0~100

- Se definido como 100, provavelmente o veículo não consumirá
quantias significantes e em alguns casos pode-se tornar combustível
infinito. Se definido 0, o veículo consumirá drasticamente e provavelmente,
de forma imediata.
*/
func (e *engineBuilder) FuelEconomy(fe float32) *engineBuilder {
	if fe > 100 {
		fe = 100
	} else if fe < 0 {
		fe = 0
	}
	e.v.engine.fuelEconomy = fe
	return e
}

/*
Altera o combustível que o veículo possuí no tanque: valores 0~100

- 100 -> Tanque cheio

- 0 -> Tanque vazio
*/
func (e *engineBuilder) Fuel(f float32) *engineBuilder {
	if f > 100 {
		f = 100
	} else if f < 0 {
		f = 0
	}
	e.v.engine.fuel = f
	return e
}

// Retorna a qunatidade de combustível no tanque
//
// 100 = cheio;
// 0 = vazio
func (e *engineBuilder) GetFuel() float32 {
	return e.v.engine.fuel
}

// Retorna a qunatidade de economia de combustível
//
// 100 = economia máxima;
// 0 = economia minima/nenhuma
func (e *engineBuilder) GetFuelEconomy() float32 {
	return e.v.engine.fuelEconomy
}

/*
Tentará dar partida ao motor do veículo, as chances de sucesso dependerá
da quantidade de combustível armazenado no tanque e da carga da bateria

- Atualiza o endereço da várivael status para true, se o motor efetuar a
partida
*/
func (e *engineBuilder) Ignite(status *bool) *engineBuilder {
	lights := 0
	alarm := 0
	doors := 0
	bonnet := 0
	boot := 0
	objective := 0

	var ignite int
	if e.canIgniteEngine() {
		ignite = 1
	}

	// e.v.Eletrics().IntroduceElectricalDrain()
	// time.Slieep(time.Duration(rand.Intn(30 - (3) + 30)))
	natives.SetVehicleParamsEx(
		e.v.id,
		ignite,
		lights,
		alarm,
		doors,
		bonnet,
		boot,
		objective,
	)
	return e
}

func (e *engineBuilder) canIgniteEngine() bool {
	charger := e.v.Eletrics().GetBatteryCharger()
	fuel := e.GetFuel()

	rand.New(rand.NewSource(0))

	if fuel == 0 || charger == 0 {
		return false
	}

	minConsumable := math.Min(float64(charger), float64(fuel))
	weight := rand.Intn(11) - 5

	failRange := rand.Intn(100) + 1
	successRange := int(minConsumable) + weight

	canIgnite := successRange > failRange

	return canIgnite
}

/*
Desliga o motor do veículo
*/
func (e *engineBuilder) TurnOff() *engineBuilder {
	lights := 0
	alarm := 0
	doors := 0
	bonnet := 0
	boot := 0
	objective := 0

	natives.GetVehicleParamsEx(
		e.v.id,
		nil,
		&lights,
		&alarm,
		&doors,
		&bonnet,
		&boot,
		&objective,
	)
	natives.SetVehicleParamsEx(
		e.v.id,
		0,
		lights,
		alarm,
		doors,
		bonnet,
		boot,
		objective,
	)
	return e
}
