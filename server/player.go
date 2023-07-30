package server

import (
	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/internal/natives"
)

func (rb *ServerBuild) Spawn(cb *chat.ChatBuilder) *ServerBuild {

	if cb != nil {
		cb.Send()
	}
	natives.SpawnPlayer(rb.playerID)
	return rb
}
