package vehicles

import "github.com/goseventh/rakstar/internal/natives"
import "math/rand"
import "time"

/*
  SetBatteryCharger é um método que altera a carga da bateria do veículo.
  - Ele recebe um argumento float32 'charger' que representa a nova carga da bateria.
  - Se 'charger' for maior que 100, ele será definido como 100. Se for menor que 0, 
    será definido como 0.
  - O método então define o campo batteryCharger do objeto eletrics do veículo para o 
    valor de 'charger'.
*/
func (e *eletricsBuilder) SetBatteryCharger(charger float32) *eletricsBuilder {
	if charger > 100 {
		charger = 100
	} else if charger < 0 {
		charger = 0
	}

	e.v.eletrics.batteryCharger = charger
	return e

}
// BatteryCharger é um método que retorna a carga atual da bateria do veículo.
func (e *eletricsBuilder) BatteryCharger() float32 {
	return e.v.eletrics.batteryCharger
}

/*
 ToggleLights é um método que altera o estado dos faróis do veículo.
   - Se os faróis estiverem ligados, o método os desligará. 
   - Se estiverem desligados, o método os ligará.
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

// IntroduceElectricalDrain é um método que simula uma drenagem elétrica no veículo.
// - O método verifica se os faróis estão ligados e se a carga da bateria é maior que 25. 
//   Se qualquer uma dessas condições não for atendida, o método retorna imediatamente.
// - Caso contrário, o método entra em um loop onde alterna os faróis várias vezes. 
//   Isso simula uma drenagem elétrica no veículo.
// - Este método deve ser usado para simular uma ignição.
func (e *eletricsBuilder) IntroduceElectricalDrain() {
	vehicleID := e.v.id
	engine := 0
	lights := 0
	alarm := 0
	doors := 0
	bonnet := 0
	boot := 0
	objective := 0

	natives.GetVehicleParamsEx(
		vehicleID,
		&engine,
		&lights,
		&alarm,
		&doors,
		&bonnet,
		&boot,
		&objective,
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
