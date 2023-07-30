package server

import (
	"github.com/goseventh/rakstar/internal/samp"
	"github.com/goseventh/rakstar/pkg/chat"
)

func (rb *ServerBuild) Spawn(cb *chat.ChatBuilder) *ServerBuild {

	if cb != nil {
		cb.Send()
	}
	samp.SpawnPlayer(rb.playerID)
	return rb
}
