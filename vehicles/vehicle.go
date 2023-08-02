package vehicle

func (v *vehicleBuilder) Pos(x, y, z, rotate float32) *vehicleBuilder {
	v.posX, v.posY, v.posZ, rotate = x, y, z, rotate
	return v
}

func (v *vehicleBuilder) Health(h float32) *vehicleBuilder {
	v.health = h
	return v
}

func (v *vehicleBuilder) Color(prim, secon int) *vehicleBuilder {
	v.colorPrimary, v.colorSecondary = prim, secon
	return v
}

/*
Escolhe o modelo do veículo para futuras criações
*/
func (v *vehicleBuilder) Model(m int) *vehicleBuilder {
	v.model = m
	return v
}
