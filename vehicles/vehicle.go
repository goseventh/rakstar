package vehicles

import (
	"math"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/vehiclesConst"
	"github.com/goseventh/rakstar/player"
)

// Invocar esta função setará a coordenada do veículo
func (v *vehicleBuilder) Coordinate(x, y, z, rotate float32) *vehicleBuilder {
	v.posX, v.posY, v.posZ, v.rotate = x, y, z, rotate
	return v
}

// Invocar esta funcão setará a saúde do veículo
func (v *vehicleBuilder) Health(h float32) *vehicleBuilder {
	v.health = h
	return v
}

// Invocar esta função setará a cor do veículo, incluíndo cores
// primárias e secundárias
func (v *vehicleBuilder) Color(prim, secon int) *vehicleBuilder {
	v.colorPrimary, v.colorSecondary = prim, secon
	return v
}

// Escolhe o modelo do veículo para futuras criações
func (v *vehicleBuilder) Model(m int) *vehicleBuilder {
	v.model = m
	return v
}

// Invocar esta função setará o texto que aparecerá
// na placa do veículo
func (v *vehicleBuilder) Place(place string) {
	natives.SetVehicleNumberPlate(v.id, place)
}

// Cria o veículo e o spawna
// seta o ID de criação do veículo após o spawn
func (v *vehicleBuilder) Create() *vehicleBuilder {
	v.id = natives.CreateVehicle(
		v.model,
		v.posX,
		v.posY,
		v.posZ,
		v.rotate,
		v.colorPrimary,
		v.colorSecondary,
		-1,
		false,
	)
	return v
}

// Adiciona um jogador a um assento disponível de um veículo
//
// a função tentará encontrar quais assentos estão disponíveis
// para o veiculo, caso houver setará automaticamente o jogador
func (v *vehicleBuilder) AttachPlayer(p *player.PlayerBuilder) *vehicleBuilder {
	var seats []int
	for i := 0; i <= natives.GetMaxPlayers(); i++ {
		vehID := natives.GetPlayerVehicleID(i)
		if vehID != v.id {
			continue
		}
		seat := natives.GetPlayerVehicleSeat(i)
		if seat <= -1 {
			continue
		}
		seats = append(seats, seat+1)
	}

	if len(seats) > 4 {
		v.AttachPlayer(p)
	}

	for s := 0; s <= 4; s++ {
		if seats[s] != 0 {
			continue
		}
		natives.PutPlayerInVehicle(p.ID, v.id, s-1)
		break
	}
	return v
}

// Invocar esta função destruirá todos os veiculos
// próximos do player que estejam na distancia X
//
// Exemplo:
//
//	p := player.Builder().Select("alph4b3eth")
//	DeleteInRange(p, 30)
func (v *vehicleBuilder) DeleteInRange(player player.PlayerBuilder, distance float64) {
	requestX, requestY, requestZ := player.GetPos()
	if distance < 0 {
		distance = 5.0
	}
	for vehicle := 0; vehicle < vehiclesConst.MaxVehicles; vehicle++ {
		var vehicleX, vehicleY, vehicleZ float32
		sucess := natives.GetVehiclePos(vehicle, &vehicleX, &vehicleY, &vehicleZ)
		if !sucess {
			continue
		}
		distanceX := (requestX - vehicleX)
		distanceY := (requestY - vehicleY)
		distanceZ := (requestZ - vehicleZ)

		vehicleDistance := math.Sqrt(
			float64(distanceX)*float64(distanceX) +
				float64(distanceY)*float64(distanceY) +
				float64(distanceZ)*float64(distanceZ),
		)
		if vehicleDistance > distance {
			continue
		}

		natives.DestroyVehicle(vehicle)
		break
	}
}

// Invocar esta função destruirá todos os veículos do servidor
// Se ocorrer falhas durante a execução, ela retornará false
// Se ocorrer com êxito, ela retornará true
func (v *vehicleBuilder) DestroyAll() bool {
	for vehicle := 0; vehicle < vehiclesConst.MaxVehicles; vehicle++ {
		sucess := natives.DestroyVehicle(vehicle)
		if !sucess {
			return false
		}
	}
	return true
}
