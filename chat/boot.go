package chat
//teste com lunarvim
import (
	"fmt"

	"github.com/goseventh/rakstar/internal/callbacks"
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/common"
)

/*
Função que é chamada pelo inicializador do RakStar
*/
func Boot() {
	callbacks.On("playerText", handlerChat)
}

func handlerChat(p natives.Player, text string) bool {
	Builder().
		Message(fmt.Sprintf("%s: %s", p.GetName(), text)).
    Select(p.ID).
		Range(Local).
		Color(common.ChatLocalColorStr).
		Tag("local").
		Send()

	return false
}
