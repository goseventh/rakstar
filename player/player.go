//player fornece funções de gerenciamento de jogadores conectados. 
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

// Life define a vida do player
func (pb *PlayerBuilder) Life(life float32) *PlayerBuilder {
	natives.SetPlayerHealth(pb.ID, life)
	return pb
}

// Armour define a armadura(colete) do jogador
func (pb *PlayerBuilder) Armour(Armour float32) *PlayerBuilder {
	natives.SetPlayerArmour(pb.ID, Armour)
	return pb
}

// Spawn faz o jogador reaparecer no jogo.
// Ele chama a função nativa SpawnPlayer com o ID do jogador como argumento.
// O ID do jogador é acessado através do campo ID do objeto PlayerBuilder.
// Após chamar a função nativa, o método retorna o próprio objeto PlayerBuilder.
func (pb *PlayerBuilder) Spawn() *PlayerBuilder {
	natives.SpawnPlayer(pb.ID)
	return pb
}

// Nick obtém o apelido (nick) do jogador.
// Ele chama a função nativa GetPlayerName com o ID do jogador 
// e uma variável de string como argumentos.
// O ID do jogador é acessado através do campo ID do objeto PlayerBuilder.
// A variável de string é fornecida como um argumento para o método e 
// será preenchida com o apelido do jogador.
// O tamanho máximo do apelido é definido pela constante MaxPlayerName 
// em playerConst.
// Após chamar a função nativa, o método retorna o próprio objeto PlayerBuilder.
func (pb *PlayerBuilder) Nick(nick *string) *PlayerBuilder {
	natives.GetPlayerName(pb.ID, nick, playerConst.MaxPlayerName)
	return pb
}

// GetCoordinate obtém a coordenada do jogador e a direção de 
// que o jogador está olhando
func (pb *PlayerBuilder) Coordinate() (float32, float32, float32, float32, error) {
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

// InVehicle é verdadeiro se o jogador está dentro de um
// veículo, caso contrário será falso
func (pb *PlayerBuilder) InVehicle() bool {
	return natives.IsPlayerInAnyVehicle(pb.ID)
}

// Vehicle retornará o ID do veículo do jogador.
// Retornará -1 para um ID inválido ou se o jogador 
// não estiver em nenhum veículo
func (pb *PlayerBuilder) Vehicle() int {
	return natives.GetPlayerVehicleID(pb.ID)
}

// DeleteCurrentVehicle destruirá o veículo do jogador
func (pb *PlayerBuilder) DeleteCurrentVehicle() bool {
	vehID := natives.GetPlayerVehicleID(pb.ID)
	return natives.DestroyVehicle(vehID)
}


// SelectCharacter seleciona o personagem (skin) do jogador. Ela recebe
// um ID [ID SKINS] de skin do SA-MP
//
// [ID SKINS]: https://sampwiki.blast.hk/wiki/Skins:All
func (pb *PlayerBuilder) SelectCharacter(skin int) error {
	sucess := natives.SetPlayerSkin(pb.ID, skin)
	if !sucess {
		return ErrFailSetCharacter
	}
	return nil
}
