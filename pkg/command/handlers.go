package command

import (
	"fmt"
	"main/pkg/server"
	"main/pkg/utils/sampstr"
	"strings"

	"github.com/goseventh/rakstar/internal/player"
	"github.com/goseventh/rakstar/pkg/chat"
)

func HandlePlayerCommandText(player player.Player, cmdtext string) bool {
	server.GoroutinePool.Submit(func() {
		splitCmdText := strings.Split(sampstr.Decode(cmdtext), " ")

		commandName := strings.Replace(splitCmdText[0], "/", "", -1)
		commandArgs := splitCmdText[1:]

		command, distance := SearchCommand(commandName)

		if command == nil {
			chat.Builder().
				Message("Nenhum comando correspondente foi encontrado!").
				Tag("servidor").
				PlayerID(player.ID).
				Send()

		}

		if distance == 2 {
			chat.Builder().
				Message(
					fmt.
						Sprintf(
							"Comando errado, resultado semelhante: %v",
							command.Name,
						),
				).
				Tag("advertência").
				PlayerID(player.ID)

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

		fmt.Printf("Executando o comando [%s] para o jogador %s\n", command.Name, player.GetName())

		commandContext := CommandContext{
			Player:     &player,
			ArgHandler: argHandler,
		}

		command.Handler(&commandContext)
	})

	return true
}
