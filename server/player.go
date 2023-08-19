package server

import (
	"errors"

	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/player"
)

var (
	ErrFailureGetIp = errors.New("Failure to obtain player's IP address.")
)

func (rb *ServerBuild) Spawn(cb *chat.ChatBuilder) *ServerBuild {

	if cb != nil {
		cb.Send()
	}
	natives.SpawnPlayer(rb.playerID)
	return rb
}

// Invocar esta função retornará o IP do jogador
func (rb *ServerBuild) GetIP(player player.PlayerBuilder) (string, error) {
	var ip string
	sucess := natives.GetPlayerIp(player.ID, &ip, 16)
	if !sucess {
		return "", ErrFailureGetIp
	}
  return ip, nil
}
