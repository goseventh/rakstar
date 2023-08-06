/*
Este pacote oferece o sistema Advanced Vehicle Engine, permitindo
a criação e simulação avançada do comportamento de veículos,
inspirados no mundo real. Com este pacote, os desenvolvedores
têm a capacidade de criar e manipular veículos virtuais com
detalhes minuciosos, incluindo características como o sistema de
combustível, consumo, parte elétrica do veículo, bateria e sistemas
de ignição.
*/
package vehicle

type vehicleBuilder struct {
	id                           int
	model                        int
	health                       float32
	colorPrimary, colorSecondary int
	posX, posY, posZ, rotate     float32
	eletrics                     struct {
		batteryCharger float32
		v              *vehicleBuilder
	}
  engine struct{
    fuel float32
    fuelEconomy float32
  }
}

type eletricsBuilder struct {
	// batteryCharger float32
	v              *vehicleBuilder
}

type engineBuilder struct {
	// fuel        float32
	// fuelEconomy float32
	v           *vehicleBuilder
}
 
func Builder() *vehicleBuilder {
	v := new(vehicleBuilder)
	v.id = -1
	return v
}

/*Funcao par ser funcar*/
func (v *vehicleBuilder) Engine() *engineBuilder {
	e := new(engineBuilder)
	e.v = v
	return e
}

func (v *vehicleBuilder) Eletrics() *eletricsBuilder {
	e := new(eletricsBuilder)
	e.v = v
	return e
}

/*
Seleciona o veículo que será usado para manipulações

- Recebe o ID de criação do veículo
*/
func (v *vehicleBuilder) Select(id int) *vehicleBuilder {
	v.id = id
	return v
}
