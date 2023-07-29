package server

import (
	"fmt"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/pkg/chat"
)

/*
	Define a mensagem que será enviada durante a contagem regresiva para reiniciar

	Exemplo: "O servidor reiniciará"

	# # Resultado: 
		- O servidor reiniciará - 5 
		- O servidor reiniciará - 4 
		- O servidor reiniciará - 3 
		- O servidor reiniciará - 2 
		- O servidor reiniciará - 1 
		- O servidor reiniciará - 0 

*/
func (rb *ServerBuild) MessageLoop(msg string) *ServerBuild {
	rb.msgLoop = msg
	return rb
}

func (rb *ServerBuild) RestartNow(cb *chat.ChatBuilder) *ServerBuild {
	if cb != nil {
		cb.Send()
	}
	if rb.msgLoop == "" {
		rb.msgLoop = "o servidor reiniciará"
	}

	time.Sleep(7 * time.Second)
	for i := 5; i > 0; i-- {
		time.Sleep(time.Second)
		cb.
			PlayerID(chat.Global).
			Message(fmt.
				Sprintf("%v - %v", rb.msgLoop, i)).
			Send()
	}

	natives.SendRconCommand("gmx")
	return rb
}
