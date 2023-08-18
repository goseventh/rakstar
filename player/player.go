package player

import (
	"errors"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

var (
	FailTeleport = errors.New("Player teleportation failure")
  FailSetCharacter = errors.New("Failure to set player character.")
  FailSelectWeapon= errors.New("Failure to select player's weapon.")
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

// Obtem o angulo do jogador
func (pb *PlayerBuilder) GetAngle() float32 {
	var angle float32
	connected := natives.GetPlayerFacingAngle(pb.ID, &angle)
	if !connected {
		return -1
	}
	return angle
}

// Obtem a posição do jogador
func (pb *PlayerBuilder) GetPos() (float32, float32, float32) {
	var (
		x, y, z float32
	)
	connected := natives.GetPlayerPos(pb.ID, &x, &y, &z)
	if !connected {
		return -1, -1, -1
	}
	return x, y, z
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
		return FailTeleport
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

func (pb*PlayerBuilder) SelectCharacter(skin int) error{
  sucess := natives.SetPlayerSkin(pb.ID, skin)
  if !sucess{
    return FailSetCharacter
  }
  return nil
}

// Invocar esta função selecionará a arma que o jogador está
// segurando
func (pb*PlayerBuilder) SelectWeapon(weapon int) error{
 sucess := natives.SetPlayerArmedWeapon(pb.ID, weapon)
  if !sucess{
    return FailSelectWeapon
  }
  return nil
}
