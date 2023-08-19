package player

import (
	"errors"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

var (
	ErrFailTeleport      = errors.New("Player teleportation failure")
	ErrFailSetCharacter  = errors.New("Failure to set player character")
	ErrFailGetCoordinate = errors.New("Failure to get player coordinate")
)

// Seta a vida do player
func (pb *PlayerBuilder) Life(life float32) *PlayerBuilder {
	natives.SetPlayerHealth(pb.ID, life)
	return pb
}

// Seta a armadura(colete) do jogador
func (pb *PlayerBuilder) Armour(Armour float32) *PlayerBuilder {
	natives.SetPlayerArmour(pb.ID, Armour)
	return pb
}

// (re)Spawna o jogador
func (pb *PlayerBuilder) Spawn() *PlayerBuilder {
	natives.SpawnPlayer(pb.ID)
	return pb
}

// Obtem o nick do jogador, e seta na variavel recebida no parametro
func (pb *PlayerBuilder) Nick(nick *string) *PlayerBuilder {
	natives.GetPlayerName(pb.ID, nick, playerConst.MaxPlayerName)
	return pb
}

// Invocar esta função retornará a coordenada do jogador, bem como 
// a direção de orientação que corresponde, por exemplo, a rotação
// baseada na bússula. Ou seja, a direção que o jogador está olhando
func (pb *PlayerBuilder) GetCoordinate() (float32, float32, float32, float32, error) {
	var (
		rotation float32
		x, y, z  float32
	)

	sucess := natives.GetPlayerFacingAngle(pb.ID, &rotation)
	sucess2 := natives.GetPlayerPos(pb.ID, &x, &y, &z)
	if !sucess || !sucess2 {
		return -1, -1, -1, -1, ErrFailGetCoordinate
	}
	return x, y, z, rotation, nil
}


// Invocar esta função teletransportará o jogador para as
// coordenadas informadas no parametro, bem como sua direção
// de orientação - direção que corresponde, por exemplo, a
// rotação baseada na bússula. Ou seja, a direção que o jogador
// está olhando
func (pb *PlayerBuilder) Teleport(x, y, z, r float32) error {
	sucess := natives.SetPlayerPos(pb.ID, x, y, z)
	sucess2 := natives.SetPlayerFacingAngle(pb.ID, r)
	if !sucess || !sucess2 {
		return ErrFailTeleport
	}
	return nil
}

// Invocar esta função retornará se o jogador está dentro de um
// veículo
func (pb *PlayerBuilder) InVehicle() bool {
	return natives.IsPlayerInAnyVehicle(pb.ID)
}

// Invocar esta função retornará o ID do veículo que o jogador está
func (pb *PlayerBuilder) GetVehicle() int {
	return natives.GetPlayerVehicleID(pb.ID)
}

// Invocar esta função destruirá o veículo do jogador
func (pb *PlayerBuilder) DeleteCurrentVehicle() bool {
	vehID := natives.GetPlayerVehicleID(pb.ID)
	return natives.DestroyVehicle(vehID)
}

func (pb *PlayerBuilder) SelectCharacter(skin int) error {
	sucess := natives.SetPlayerSkin(pb.ID, skin)
	if !sucess {
		return ErrFailSetCharacter
	}
	return nil
}
