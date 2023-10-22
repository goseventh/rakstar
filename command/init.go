package command

import (
	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/internal/callbacks"
	"github.com/goseventh/rakstar/internal/utils/common"
)

func init() {
	chat := chat.Builder()
	chat.Color(common.WarnColorStr).
		Message("Nenhum comando foi encontrado")
	SetConfig(chat, "Comando errado, similar")

	err := callbacks.On("playerCommandText", HandlePlayerCommandText)
  if err != nil{
    panic(err)
  }
}
