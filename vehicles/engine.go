package vehicle

type engine struct {
	fuel        float32
	fuelEconomy float32
	v           *vehicleBuilder
}

func (e *engine) Egnite() *engine {
	return e
}

func (e *engine) Fuel(f float32) *engine {
	e.fuel = f
	return e

}

func (e *engine) Ignite() *engine {
	return e
}

func (e *engine) TurnOff() *engine {
	return e
}
