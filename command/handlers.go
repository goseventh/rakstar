package command

import (
	"fmt"
	"log"
	"strconv"

	//"main/pkg/server"
	//	"main/pkg/utils/sampstr"
	"strings"

	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/common"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
	"github.com/goseventh/rakstar/internal/utils/sampstr"
	"github.com/goseventh/rakstar/server"
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
func SetConfig(notFoundChat *chat.ChatBuilder, similarFoundMsg string) {
	NotFoundChat = notFoundChat
	SimiularFoundMSG = &similarFoundMsg
}

/*
Função que deve ser chamada na callback "OnPlayerCommand"
*/
func HandlePlayerCommandText(player natives.Player, cmdtext string) bool {
	server.Builder().
		Goroutine().
		Submit(func() {
			splitCmdText := strings.Split(sampstr.Decode(cmdtext), " ")

			commandName := strings.Replace(splitCmdText[0], "/", "", -1)
			commandArgs := splitCmdText[1:]

			command, distance := SearchCommand(commandName)

			if command == nil {
				NotFoundChat.
					Select(player.ID).
					Tag("rakstar").
					Send()
				return
			}

			if distance == 2 {

				chat.Builder().
					Select(player.ID).
					Color(common.WarnColorStr).
					Tag("rakstar").
					Message(fmt.
						Sprintf(
							"%v: %v",
							*SimiularFoundMSG,
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

			commandContext := CommandContext{
				Player:     &player,
				ArgHandler: argHandler,
			}

			var pass bool = true

		condLoop:
			for idx, arg := range commandArgs {
				for _, cond := range command.conditionals_[idx] {
					switch cond.typeIdx {
					case typePlayer:
						var id int = -1
						var err error
						id, err = strconv.Atoi(arg)
						if err != nil {
							var nick string
							for i := 0; i < playerConst.MaxPlayers; i++ {
								natives.GetPlayerName(i, &nick, playerConst.MaxPlayerName)
								if nick == arg {
									id = i
									break
								}
							}

						}

						switch cond.cond {
						case MustPlayerConnected:
							if !natives.IsPlayerConnected(id) {
								log.Printf("[rakstar-cmd idx(%v)] o jogador %v não está conectado", idx, id)
								pass = false
								break condLoop
							}

							log.Printf("[rakstar-cmd idx(%v)] o jogador %v está conectado", idx, id)
						case MustNickIs:
							var nick string
							natives.GetPlayerName(id, &nick, playerConst.MaxPlayerName)
							if cond.value != nick {
								log.Printf("[rakstar-cmd idx(%v)] falha na comparação de nicks entre %v:%v", idx, nick, cond.value)
								pass = false
								break condLoop
							}

							log.Printf("[rakstar-cmd idx(%v)] boa comparação de nicks entre %v:%v", idx, nick, cond.value)

						}
					}

				}
			}

			if !pass {
				return
			}

			log.Printf("[rakstar] rodando o comando [%s] para o jogador %s\n", command.Name, player.GetName())
			command.Handler(&commandContext)
		})

	return true
}
