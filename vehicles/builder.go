package vehicle

type vehicleBuilder struct {
	id                           int
	model                        int
	health                       float32
	colorPrimary, colorSecondary int
	posX, posY, posZ, rotate     float32
}

func Builder() *vehicleBuilder {
	return new(vehicleBuilder)
}

func (v *vehicleBuilder) Engine() *engine {
	e := new(engine)
	e.v = v
	return e
}

func (v *vehicleBuilder) Eletrics() *eletrics {
	e := new(eletrics)
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
