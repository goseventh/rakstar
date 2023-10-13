package vehicles

import (
	"math"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/vehiclesConst"
	"github.com/goseventh/rakstar/player"
	"golang.org/x/exp/slices"
)

/*
ID é um método que retorna o ID de criação do veículo.
  - O método retorna o valor do campo id do objeto vehicleBuilder.
  - Nota: O ID de criação do veículo é o ID em ordem crescente
    de criação de veículo do servidor.
*/
func (v *vehicleBuilder) ID() int {
	return v.id
}

/*
Model é um método que retorna o modelo do veículo.
  - O método retorna o valor do campo model do objeto vehicleBuilder.
*/
func (v *vehicleBuilder) Model() int {
	return v.model
}

/*
SetCoordinate é um método que define a coordenada do veículo.
  - Ele recebe quatro argumentos float32: 'x', 'y', 'z' e 'rotate',
    que representam a posição e a rotação do veículo.
  - O método define os campos posX, posY, posZ e rotate do objeto
    vehicleBuilder para os valores fornecidos.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que
    chamadas de método sejam encadeadas em uma única linha.
*/
func (v *vehicleBuilder) SetCoordinate(x, y, z, rotate float32) *vehicleBuilder {
	v.posX, v.posY, v.posZ, v.rotate = x, y, z, rotate
	return v
}

/*
SetHealth é um método que define a saúde do veículo.
  - Ele recebe um argumento float32 'h' que representa a nova saúde do veículo.
  - O método define o campo health do objeto vehicleBuilder para o valor de 'h'.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que chamadas
    de método sejam encadeadas em uma única linha.
*/
func (v *vehicleBuilder) SetHealth(h float32) *vehicleBuilder {
	v.health = h
	return v
}

/*
SetColor é um método que define a cor do veículo, incluindo as cores primária e secundária.
  - Ele recebe dois argumentos inteiros: 'prim' e 'secon', que representam as cores primária
    e secundária, respectivamente.
  - O método define os campos colorPrimary e colorSecondary do objeto vehicleBuilder para
    os valores fornecidos.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que chamadas de método
    sejam encadeadas em uma única linha.
*/
func (v *vehicleBuilder) SetColor(prim, secon int) *vehicleBuilder {
	v.colorPrimary, v.colorSecondary = prim, secon
	return v
}

/*
SetModel é um método que escolhe o modelo do veículo para futuras criações.
  - Ele recebe um argumento inteiro 'm' que representa o novo modelo do veículo.
  - O método define o campo model do objeto vehicleBuilder para o valor de 'm'.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que chamadas de
    método sejam encadeadas em uma única linha.

Consulte os ID de cada model do SA-MP: [ID DOS MODELOS]

[ID DOS MODELOS]: https://sampwiki.blast.hk/wiki/Vehicles:All
*/
func (v *vehicleBuilder) SetModel(m int) *vehicleBuilder {
	v.model = m
	return v
}

/*
SetPlate é um método que define o texto que aparecerá na placa do veículo.
  - Ele recebe um argumento string 'plate' que representa o novo texto da placa.
  - O método define o campo plate do objeto vehicleBuilder para o valor de 'plate'.
  - Se o ID do veículo for diferente de -1, o método também chamará a função nativa
    SetVehicleNumberPlate com o ID do veículo e o novo texto da placa.
*/
func (v *vehicleBuilder) SetPlate(plate string) {
	v.plate = plate

	if v.id == -1 {
		return
	}

	natives.SetVehicleNumberPlate(v.id, plate)
}

/*
Create é um método que cria e faz spawn de um veículo no jogo.
  - O método chama a função nativa CreateVehicle com os parâmetros apropriados obtidos
    dos campos do objeto vehicleBuilder.
  - A função nativa retorna um ID de veículo, que é então armazenado no campo id do
    objeto vehicleBuilder.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que chamadas de método
    sejam encadeadas em uma única linha.
*/
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

/*
AttachPlayer é um método que adiciona um jogador a um assento disponível de um veículo.
  - Ele recebe um argumento do tipo [*player.PlayerBuilder] que representa o jogador a ser
    adicionado ao veículo.
  - O método verifica quais assentos estão ocupados e seleciona o primeiro assento disponível.
  - Em seguida, ele chama a função nativa PutPlayerInVehicle com o ID do jogador, o ID do
    veículo e o assento disponível.
  - O método retorna o próprio objeto vehicleBuilder, permitindo que chamadas de método sejam
    encadeadas em uma única linha.
*/
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

/*
DeleteInRange é um método que destrói todos os veículos próximos a um jogador que estejam
dentro de uma certa distância.
  - Ele recebe dois argumentos: um objeto player.PlayerBuilder que representa o jogador
    e um float64 que representa a distância.
  - O método verifica a distância entre cada veículo no servidor e o jogador. Se um
    veículo estiver dentro da distância especificada, ele será destruído.
  - Nota: Se a distância fornecida for menor que 0, ela será definida como 5.0.
*/
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

/*
DestroyAll é um método que destrói todos os veículos do servidor.
  - O método percorre todos os veículos no servidor e tenta destruir cada um chamando
    a função nativa DestroyVehicle.
  - Se qualquer chamada para DestroyVehicle retornar false, indicando que o veículo
    não pôde ser destruído, o método retornará false.
  - Se todos os veículos forem destruídos com sucesso, o método retornará true.
*/
func (v *vehicleBuilder) DestroyAll() bool {
	for vehicle := 0; vehicle < vehiclesConst.MaxVehicles; vehicle++ {
		sucess := natives.DestroyVehicle(vehicle)
		if !sucess {
			return false
		}
	}
	return true
}

/*
  Destroy é um método que destrói o veículo selecionado.
    - O método chama a função nativa DestroyVehicle com o ID do veículo.
    - Se a função nativa retornar false, indicando que o veículo não pôde
      ser destruído, o método retornará false.
    - Se o veículo for destruído com sucesso, o método retornará true.
*/
func (v *vehicleBuilder) Destroy() bool {
	return natives.DestroyVehicle(v.id)
}

/*
  VehicleModelName é um método que está em construção, que recebe um
  ID de modelo [ID DOS MODELOS] e retorna o nome respectivo do veículo
    - Nota: A lista de veículos está em construção e alguns tipos de veículos
    não possuem suporte. Também não há garantias de estabilidade.
    Chamar esse método poderá resultar em retornos incorretos ou
    vazios.

  [ID DOS MODELOS]: https://sampwiki.blast.hk/wiki/Vehicles:All
*/
func (v *vehicleBuilder) VehicleModelName(modelID int) string {
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
