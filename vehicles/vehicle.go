package vehicles

import (
	"math"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/vehiclesConst"
	"github.com/goseventh/rakstar/player"
	"golang.org/x/exp/slices"
)

// Invocar esta função retornará o ID de criação do veículo,
// ou seja, o ID em ordem crescente de criação de veíuclo do
// servidor.
func (v *vehicleBuilder) GetID() int {
	return v.id
}

// Invocar esta função retornará o modelo do veículo
func (v *vehicleBuilder) GetModel() int {
	return v.model
}

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
func (v *vehicleBuilder) Plate(plate string) {
	v.plate = plate

	if v.id == -1 {
		return
	}

	natives.SetVehicleNumberPlate(v.id, plate)
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
	occupiedSeats := []int{}

	for i := 0; i <= natives.GetMaxPlayers(); i++ {
		vehicleID := natives.GetPlayerVehicleID(i)

		if vehicleID != v.id {
			continue
		}

		seat := natives.GetPlayerVehicleSeat(i)

		if seat == -1 {
			continue
		}

		occupiedSeats = append(occupiedSeats, seat)
	}

	availableSeat := -1

	for seat := 0; seat < 30; seat++ {
		if slices.Contains(occupiedSeats, seat) {
			continue
		}

		availableSeat = seat
		break
	}

	if availableSeat == -1 {
		return v
	}

	natives.PutPlayerInVehicle(p.ID, v.id, availableSeat)

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
	requestX, requestY, requestZ, _, err := player.GetCoordinate()
	if err != nil {
		return
	}
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

// Invocar esta função destruirá o veículo selecionado
// Se ocorrer falhas durante a execução, ela retornará false
// Se ocorrer com êxito, ela retornará true
func (v *vehicleBuilder) Destroy() bool {
	return natives.DestroyVehicle(v.id)
}

func (v *vehicleBuilder) GetVehicleModelName(modelID int) string {
	switch modelID {
	case vehiclesConst.VehicleLandstalker:
		return "Land Stalker"
	case vehiclesConst.VehicleBravura:
		return "Bravura"
	case vehiclesConst.VehicleBuffalo:
		return "Buffalo"
	case vehiclesConst.VehicleLinerunner:
		return "Liner Runner"
	case vehiclesConst.VehiclePerrenial:
		return "Perrenial"
	case vehiclesConst.VehicleSentinel:
		return "Sentinel"
	case vehiclesConst.VehicleDumper:
		return "Dumper"
	case vehiclesConst.VehicleFiretruck:
		return "FireTruck"
	case vehiclesConst.VehicleTrashmaster:
		return "Trash Master"
	case vehiclesConst.VehicleStretch:
		return "Stretch"
	case vehiclesConst.VehicleManana:
		return "Manana"
	case vehiclesConst.VehicleInfernus:
		return "Infernus"
	case vehiclesConst.VehicleVoodoo:
		return "Voodoo"
	case vehiclesConst.VehiclePony:
		return "Pony"
	case vehiclesConst.VehicleMule:
		return "Mule"
	case vehiclesConst.VehicleCheetah:
		return "Cheetah"
	case vehiclesConst.VehicleAmbulance:
		return "Ambulance"
	case vehiclesConst.VehicleLeviathan:
		return "Leviathan"
	case vehiclesConst.VehicleMoonbeam:
		return "Moonbeam"
	case vehiclesConst.VehicleEsperanto:
		return "Esperanto"
	case vehiclesConst.VehicleTaxi:
		return "Taxi"
	case vehiclesConst.VehicleWashington:
		return "Washigton"
	case vehiclesConst.VehicleBobcat:
		return "Bob Cat"
	case vehiclesConst.VehicleMrwhoopee:
		return "Mr. whoopee"
	case vehiclesConst.VehicleBfinjection:
		return "Bf Injection"
	case vehiclesConst.VehicleHunter:
		return "Hunter"
	case vehiclesConst.VehiclePremier:
		return "Premier"
	case vehiclesConst.VehicleEnforcer:
		return "Enforcer"
	case vehiclesConst.VehicleSecuricar:
		return "Securicar"
	case vehiclesConst.VehicleBanshee:
		return "Banshee"
	case vehiclesConst.VehiclePredator:
		return "Predator"
	case vehiclesConst.VehicleBus:
		return "Bus"
	case vehiclesConst.VehicleRhino:
		return "Rhino"
	case vehiclesConst.VehicleBarracks:
		return "VehicleBarracks"
	case vehiclesConst.VehicleHotknife:
		return "Hotknife"
	case vehiclesConst.VehicleArticletrailer1:
		return "Article trailer1"
	case vehiclesConst.VehiclePrevion:
		return "Previon"
	case vehiclesConst.VehicleCoach:
		return "Coach"
	case vehiclesConst.VehicleCabbie:
		return "Cabbie"
	case vehiclesConst.VehicleStallion:
		return "Stallion"
	case vehiclesConst.VehicleRumpo:
		return "Rumpo"
	case vehiclesConst.VehicleRcbandit:
		return "RC Bandit"
	case vehiclesConst.VehicleRomero:
		return "Romero"
	case vehiclesConst.VehiclePacker:
		return "Packer"
	case vehiclesConst.VehicleMonster:
		return "Monster"
	case vehiclesConst.VehicleAdmiral:
		return "Admiral"
	case vehiclesConst.VehicleSqualo:
		return "Squalo"
	case vehiclesConst.VehicleKart:
		return "Kart"
	default:
		return "undefined"
	}
}
