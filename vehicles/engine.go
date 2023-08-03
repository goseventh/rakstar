package vehicle

import (
	"math/rand"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
)

type engine struct {
	fuel        float32
	fuelEconomy float32
	v           *vehicleBuilder
}

func (e *engine) Fuel(f float32) *engine {
	e.fuel = f
	return e

}

func (e *engine) Ignite() *engine {
	lights := 0
	alarm := 0
	doors := 0
	bonnet := 0
	boot := 0
	objective := 0

	min := -1
	switch {
	case (e.v.Eletrics().
		batteryCharger > 80):

		min = rand.Intn((-1) - (-2) + (-1))

	case (e.v.Eletrics().
		batteryCharger < 50 &&
		e.v.Eletrics().
			batteryCharger > 30):

		min = rand.Intn((-1) - (-4) + (-1))

	case (e.v.Eletrics().
		batteryCharger < 25):

		min = rand.Intn((-1) - (-17) + (-1))
	}
	max := 1
	switch {
	case (e.fuel > 80):
		max = rand.Intn(1 - 0 + 1)
	case (e.fuel < 50 && e.fuel > 30):
		max = rand.Intn(1 - (-1) + 1)
	case (e.fuel < 25):
		max = rand.Intn(1 - (-5) + 1)
	}

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
	time.Sleep(time.Duration(rand.Intn(30 - (3) + 30)))
	natives.SetVehicleParamsEx(
		e.v.id,
		rand.Intn(max-(min)+max),
		lights,
		alarm,
		doors,
		bonnet,
		boot,
		objective,
	)
	return e
}

func (e *engine) TurnOff() *engine {
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
