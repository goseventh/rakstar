package player

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

// Select é um operador lógico que seleciona um jogador com base em um apelido ou ID 
// fornecido.
// Ele recebe um argumento que pode ser uma string (representando o apelido do jogador)
// ou um inteiro (representando o ID do jogador).
// Se o argumento for uma string, o método percorrerá todos os jogadores e comparará
// o apelido de cada jogador com a string fornecida.
// Se encontrar um jogador cujo apelido corresponda à string, ele definirá o campo ID
// do objeto PlayerBuilder para o ID desse jogador.
// Se o argumento for um inteiro, o método simplesmente definirá o campo ID do objeto
// PlayerBuilder para esse inteiro.
func (pb *PlayerBuilder) Select(arg interface{}) *PlayerBuilder {
	switch v := arg.(type) {
	case string:
		var name string
		for i := 0; i < playerConst.MaxPlayers; i++ {
			natives.GetPlayerName(i, &name, playerConst.MaxPlayerName)
			if name == v {
				pb.ID = i
			}
			return pb

		}
	case int:
		pb.ID = v
		return pb

	}

	return pb
}
