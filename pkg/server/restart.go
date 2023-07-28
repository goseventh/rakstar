package server

import (
	"fmt"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/pkg/chat"
)



func (rb *ServerBuild) Message(msg string) *ServerBuild {
	rb.msgRestart = msg
	return rb
}

func (rb *ServerBuild) MessageLoop(msg string) *ServerBuild {
	rb.msgLoop = msg
	return rb
}

func (rb *ServerBuild) RestartNow() *ServerBuild {
	if rb.msgRestart == "" {
		rb.msgRestart = "ordem manual para reiniciar servidor em momentos..."
	}

	if rb.tag == "" {
		rb.tag = "rakstar"
	}

	chat.Builder().Message(rb.msgLoop).Tag(rb.tag)
	time.Sleep(7 * time.Second)
	for i := 5; i > 0; i-- {
		time.Sleep(time.Second)
		chat.Builder().
			Message(
				fmt.
					Sprintf("%v - %v", rb.msgLoop, i),
			)
	}

	natives.SendRconCommand("gmx")
	return rb
}
