package server

import (
	"fmt"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/pkg/chat"
)

type RestartBuild struct {
	msgRestart string
	msgLoop    string
	tag        string
}

func Builder() *RestartBuild {
	return new(RestartBuild)
}

func (rb *RestartBuild) Message(msg string) *RestartBuild {
	rb.msgRestart = msg
	return rb
}

func (rb *RestartBuild) MessageLoop(msg string) *RestartBuild {
	rb.msgLoop = msg
	return rb
}

func (rb *RestartBuild) RestartNow() *RestartBuild {
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
