package server

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/pkg/chat"
)

func (rb *ServerBuild) Spawn(cb *chat.ChatBuilder) *ServerBuild {

	if cb != nil {
		cb.Send()
	}
	natives.SpawnPlayer(rb.playerID)
	return rb
}
