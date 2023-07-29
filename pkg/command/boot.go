package command

import (
	"github.com/goseventh/rakstar/internal/utils/common"
	"github.com/goseventh/rakstar/pkg/chat"
)

/*
Função que é chamada pelo inicializador do RakStar
*/
func Boot() {
	cb := chat.Builder()
	cb.Color(common.WarnColorStr).
		Message("Nenhum comando foi encontrado")
	SetConfig(cb, "Comando errado, similar")
	//Registra o handler na callback
}
