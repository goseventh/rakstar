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

/*
	Efetiva a ordem de reinício, recebendo um objeto builder 
	
	# # Exemplo:

	cb := chat.Build().
	PlayerID(chat.global).
	Tag("servidor").
	Color(common.WarnColorStr).
	Message("ordem de reinício")
	    
	server.
	Build().
	MessageLoop("O servidor reiniciará").
	RestartNow(cb)

	# * Resultado do chat:
		- [SERVIDOR] ordem de reínicio
		- [SERVIDOR] o servidor reiniciará - 5
		- [SERVIDOR] o servidor reiniciará - 4
		- [SERVIDOR] o servidor reiniciará - 3
		- [SERVIDOR] o servidor reiniciará - 2
		- [SERVIDOR] o servidor reiniciará - 1
		- [SERVIDOR] o servidor reiniciará - 0
	... servidor reiniciou
	
*/
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
