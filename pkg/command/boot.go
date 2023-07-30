package command

import (
	rakstar "github.com/goseventh/rakstar/internal"
	"github.com/goseventh/rakstar/internal/utils/common"
	"github.com/goseventh/rakstar/pkg/chat"
)

/*
Função que é chamada pelo inicializador do RakStar
*/
func Boot() bool {
	cb := chat.Builder()
	cb.Color(common.WarnColorStr).
		Message("Nenhum comando foi encontrado")
	SetConfig(cb, "Comando errado, similar")
	
	err := rakstar.On("playerCommandText", HandlePlayerCommandText)
	return err == nil

	/* Registra o handler na callback

	server.Builder().
		SetHandler(HandlePlayerCommandText).
		Subscribe(OnPlayerCommand)

	*/

}
