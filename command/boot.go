package command

import (
	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/internal/callbacks"
	"github.com/goseventh/rakstar/internal/utils/common"
)

/*
Função que é chamada pelo inicializador do RakStar
*/
func Boot() bool {
	cb := chat.Builder()
	cb.Color(common.WarnColorStr).
		Message("Nenhum comando foi encontrado")
	SetConfig(cb, "Comando errado, similar")

	err := callbacks.On("playerCommandText", HandlePlayerCommandText)
	return err == nil

	/* Registra o handler na callback

	server.Builder().
		SetHandler(HandlePlayerCommandText).
		Subscribe(OnPlayerCommand)

	*/

}
