package server

import (
	"main/pkg/utils/common"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/pkg/chat"
)

func (rb *ServerBuild) Player(playerid int) *ServerBuild {
	rb.playerID = playerid
	return rb
}

func (rb *ServerBuild) Wait(wait time.Duration) *ServerBuild {
	rb.wait = wait
	return rb
}

func (rb *ServerBuild) Expulse() *ServerBuild {
	if rb.playerID == -1 {
		return rb
	}

	chat.Builder().
		Color(common.WarnColorStr).
		Tag("servidor").
		Message(rb.message)

	if rb.wait.Seconds() < 1 {
		natives.Kick(rb.playerID)
		return rb
	}

	time.AfterFunc(rb.wait, func() {
		natives.Kick(rb.playerID)

	})

	return rb
}
