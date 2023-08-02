package vehicle

type eletrics struct {
	batteryCharger float32
	v              *vehicleBuilder
}

func (e *eletrics) BatteryCharger(charger float32) *eletrics {
	e.batteryCharger = charger
	return e
}

func (e *eletrics) ToggleLights() *eletrics {
	return e
}
