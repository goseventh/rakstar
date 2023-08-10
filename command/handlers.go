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

type rawCommand struct {
	name string
	args []string
}

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
			processCommand(player, cmdtext)
		})

	return true
}

func parseCommandArgs(cmdtext string) rawCommand {
	splitCmdText := strings.Split(sampstr.Decode(cmdtext), " ")

	name := strings.Replace(splitCmdText[0], "/", "", -1)
	args := splitCmdText[1:]

	return rawCommand{
		name,
		args,
	}
}

func parseArgHandler(args []string) ArgHandler {
	var argHandler = ArgHandler{}

	if len(args) >= 1 {
		argHandler.args = args
		argHandler.input = strings.Join(args, " ")
		argHandler.currentArg = 0
	}

	return argHandler
}

func validateArgs(command *Command, args []string) bool {
	if len(command.conditionals_) > 0 && len(args) <= 0 {
		return false
	}

	for idx, arg := range args {
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
						return false
					}

				case MustNickIs:
					var nick string
					natives.GetPlayerName(id, &nick, playerConst.MaxPlayerName)
					if cond.value != nick {
						log.Printf("[rakstar-cmd idx(%v)] falha na comparação de nicks entre %v:%v", idx, nick, cond.value)
						return false
					}
				}
			}

		}
	}

	return true
}

func processCommand(player natives.Player, cmdtext string) bool {
	rawCommand := parseCommandArgs(cmdtext)
	command, distance := SearchCommand(rawCommand.name)

	if command == nil {
		NotFoundChat.
			Select(player.ID).
			Tag("rakstar").
			Send()

		return false
	}

	if distance >= 1 && distance <= 2 {
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

	argHandler := parseArgHandler(rawCommand.args)

	context := CommandContext{
		Player:     &player,
		ArgHandler: &argHandler,
	}

	isValidArgs := validateArgs(command, rawCommand.args)

	if !isValidArgs {
		return false
	}

	log.Printf("[rakstar] running command [%s] for player %s\n", command.Name, player.GetName())

	command.Handler(&context)

	return true
}
