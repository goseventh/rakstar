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

func (rb *ServerBuild) Wait(wait... time.Duration) *ServerBuild {
	if wait[0].Seconds() < 1 {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])
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

	natives.Kick(rb.playerID)

	return rb
}
