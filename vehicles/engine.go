package vehicles

import (
	"math"
	"math/rand"
	// "time"

	"github.com/goseventh/rakstar/internal/natives"
)

/*
  SetFuelEconomy é um método que altera a economia de combustível do veículo.
  - Ele recebe um argumento float32 'fe' que representa a nova economia de combustível.
  - Se 'fe' for maior que 100, ele será definido como 100. Se for menor que 0, será 
    definido como 0.
  - O método então define o campo fuelEconomy do objeto engine do veículo para o valor de 'fe'.
  - O método retorna o próprio objeto engineBuilder, permitindo que chamadas de método sejam 
    encadeadas em uma única linha.
  - Nota: Uma economia de combustível de 100 significa que o veículo provavelmente não 
          consumirá quantidades significativas de combustível e pode até ter combustível 
          infinito. 
          Uma economia de combustível de 0 significa que o veículo consumirá combustível 
          drasticamente e provavelmente de forma imediata.
*/
func (e *engineBuilder) SetFuelEconomy(fe float32) *engineBuilder {
	if fe > 100 {
		fe = 100
	} else if fe < 0 {
		fe = 0
	}
	e.v.engine.fuelEconomy = fe
	return e
}

// SetFuel é um método que altera a quantidade de combustível no tanque do veículo.
//  - Ele recebe um argumento float32 'f' que representa a nova quantidade de combustível.
//  - Se 'f' for maior que 100, ele será definido como 100. Se for menor que 0, será definido como 0.
//  - O método então define o campo fuel do objeto engine do veículo para o valor de 'f'.
//  - O método retorna o próprio objeto engineBuilder, permitindo que chamadas de método sejam encadeadas 
//    em uma única linha.
//  - Nota: Um valor de combustível de 100 representa um tanque cheio, enquanto um valor de 0 representa 
//         um tanque vazio.
func (e *engineBuilder) SetFuel(f float32) *engineBuilder {
	if f > 100 {
		f = 100
	} else if f < 0 {
		f = 0
	}
	e.v.engine.fuel = f
	return e
}

/*
Fuel é um método que retorna a quantidade atual de combustível no tanque do veículo.
- O método retorna o valor do campo fuel do objeto engine do veículo.
- Nota: Um valor de retorno de 100 representa um tanque cheio, enquanto um valor 
        de retorno de 0 representa um tanque vazio.
*/
func (e *engineBuilder) Fuel() float32 {
	return e.v.engine.fuel
}

// FuelEconomy é um método que retorna a economia atual de combustível do veículo.
//   - O método retorna o valor do campo fuelEconomy do objeto engine do veículo.
//   - Nota: Um valor de retorno de 100 representa a máxima economia de combustível, 
//               enquanto um valor de retorno de 0 representa nenhuma economia/minima 
//               economia de combustível.
func (e *engineBuilder) FuelEconomy() float32 {
	return e.v.engine.fuelEconomy
}

/*
Ignite é um método que tenta dar partida no motor do veículo. As chances de sucesso 
dependem da quantidade atual de combustível no tanque e da carga atual da bateria.

  - Ele recebe um argumento booleano 'status' por referência. Se o motor for iniciado 
    com sucesso, 'status' será atualizado para true.
  - O método chama a função nativa SetVehicleParamsEx com o ID do veículo e outros 
    parâmetros relevantes. Se a função nativa indicar que o motor foi iniciado com sucesso, 
    'status' será atualizado para true.
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
TurnOff desliga o motor do veículo
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
