package command

import (
	"fmt"
	//"main/pkg/server"
	//	"main/pkg/utils/sampstr"
	"strings"

	"github.com/goseventh/rakstar/internal/player"
	"github.com/goseventh/rakstar/internal/utils/sampstr"
	"github.com/goseventh/rakstar/pkg/chat"
	"github.com/goseventh/rakstar/pkg/server"
)

var NotFoundChat *chat.ChatBuilder
var SimiularFoundMSG *string


/*
	seta as mensagens que são executadas qunado um comando similar é encontrado, ou quando nenhum é encontrado.
	
	# Exemplo:

	cb := chat.builder()
	
	cb.Message("nenhum comando encontrado")

	similarFound := "comando similar"

	SetConfig(cb, similarFound)

	# # Resultado:

		- *Jogador digita "/command", mas não existe:
	
		> chat: nenhum comando foi encontrado

		- *Jogador digita "/aujad", e similar foi encontrado: "ajuda"
	
		> chat: comando similar: ajuda

*/
func SetConfig(NotFoundChat *chat.ChatBuilder, SimilarFoundMsg string){
	NotFoundChat = NotFoundChat
	SimilarFoundMsg = SimilarFoundMsg
}


/*
Função que deve ser chamada na callback "OnPlayerCommand"
*/
func HandlePlayerCommandText(player player.Player, cmdtext string) bool {
	server.Builder().
		Goroutine().
		Submit(func() {
			splitCmdText := strings.Split(sampstr.Decode(cmdtext), " ")

			commandName := strings.Replace(splitCmdText[0], "/", "", -1)
			commandArgs := splitCmdText[1:]

			command, distance := SearchCommand(commandName)

			if command == nil {
				NotFoundChat.
					PlayerID(player.ID).
					Tag("rakstar").
					Send()
			}

			if distance == 2 {

				chat.Builder().
					PlayerID(player.ID).
					Tag("rakstar").
					Message(fmt.
						Sprintf(
							"%v: %v",
							SimiularFoundMSG,
							command.Name,
						)).
					Send()

			}

			var argHandler *ArgHandler = &ArgHandler{}

			if len(commandArgs) >= 1 {
				argHandler.args = commandArgs
				argHandler.input = strings.Join(commandArgs, " ")
				argHandler.currentArg = 0
			}

			interceptorContext := &CommandInterceptorContext{
				Player:     &player,
				ArgHandler: argHandler,
				next:       true,
			}

			for _, intercept := range command.Interceptors {
				interceptorContext.next = false
				intercept(interceptorContext)

				if !interceptorContext.next {
					return
				}
			}

			fmt.Printf("[rakstar] rodando o comando [%s] para o jogador %s\n", command.Name, player.GetName())

			commandContext := CommandContext{
				Player:     &player,
				ArgHandler: argHandler,
			}

			command.Handler(&commandContext)
		})

	return true
}
