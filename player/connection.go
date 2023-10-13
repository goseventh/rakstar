package player

import "github.com/goseventh/rakstar/internal/natives"

// Connected verifica se o jogador está conectado ao servidor,
// e a condição é armazenada em status *bool. Se jogador 
// está conectado então status é verdadeiro, caso
// contrário será é falso
func (pb *PlayerBuilder) Connected(status *bool) *PlayerBuilder {
	*status = natives.IsPlayerConnected(pb.ID)
	return pb
}
