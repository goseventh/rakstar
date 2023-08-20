package command

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/goseventh/rakstar/chat"
	"github.com/goseventh/rakstar/goroutines"
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/common"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
	"github.com/goseventh/rakstar/internal/utils/sampstr"
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
  goroutines.Submit(func() {
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

func getIDfromName(nick string) int {
	var id int = -1
	var err error
	id, err = strconv.Atoi(nick)
	if err != nil {
		var nickSearch string
		for i := 0; i < playerConst.MaxPlayers; i++ {
			natives.GetPlayerName(i, &nickSearch, playerConst.MaxPlayerName)
			if compareNicks(nickSearch, nick) {
				break
			}
		}
	}
	return id
}

func compareNicks(nick, nick2 string) bool {
	return nick == nick2
}

func getNickFromID(id int) string {
	nick := ""
	natives.GetPlayerName(id, &nick, playerConst.MaxPlayerName)
	return nick
}

func isConnected(id int) bool {
	return natives.IsPlayerConnected(id)
}

func verifyTypePlayer(cond condition, idx int, arg string) bool {
	id := getIDfromName(arg)
	switch cond.cond {
	case MustPlayerConnected:
		if !isConnected(id) {
			log.Printf("[rakstar-cmd idx(%v)] o jogador %v não está conectado", idx, id)
			return false
		}

	case MustNickIs:
		nick := getNickFromID(id)
		if !compareNicks(nick, cond.value.(string)) {
			log.Printf("[rakstar-cmd idx(%v)] falha na comparação de nicks entre %v:%v",
				idx, nick, cond.value)
			return false
		}
	}
	return true
}

func valueStrBeBetween(xStr string, lessGreater []int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	if len(lessGreater) == 0 {
		return false
	}
	if len(lessGreater) > 2 {
		return false
	}
	return x < lessGreater[0] && x > lessGreater[1]
}

func valueStrBeGreeter(xStr string, y int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	return x > y
}

func valueStrBeLess(xStr string, y int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	return x < y
}

func valueStrEqual(xStr string, y int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	return x == y
}

func valueStrDivisibleBy(xStr string, y int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	return x%y == 0
}

func valueStrSquareRootOf(xStr string, y int) bool {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return false
	}
	return x*x == y
}

func verifyTypeNumber(cond condition, idx int, arg string) bool {
	switch cond.cond {
	case MustBeBetween:
		if !valueStrBeBetween(arg, cond.value.([]int)) {
			return false
		}
	case MustBeGreaterThan:
		if !valueStrBeGreeter(arg, cond.value.(int)) {
			return false
		}
	case MustBeLessThan:
		if !valueStrBeLess(arg, cond.value.(int)) {
			return false
		}
	case MustEqual:
		if !valueStrEqual(arg, cond.value.(int)) {
			return false
		}
	case MustBeDivisibleBy:
		if !valueStrDivisibleBy(arg, cond.value.(int)) {
			return false
		}
	case MustBeMultipleOf:
		if !valueStrDivisibleBy(arg, cond.value.(int)) {
			return false
		}
	case MustBeSquareRootOf:
		if !valueStrSquareRootOf(arg, cond.value.(int)) {
			return false
		}
	}
	return true
}

func textIsUpper(text string) bool {
	return text == strings.ToUpper(text)
}

func textIsLower(text string) bool {
	return text == strings.ToLower(text)
}

func textIsPrefix(text, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}

func textIsSuffix(text, prefix string) bool {
	return strings.HasSuffix(text, prefix)
}

func textIsRegMatch(text, regex string) bool {
	ok, err := regexp.Match(regex, []byte(text))
	if err != nil {
		return false
	}
	return ok
}

func verifyTypeText(cond condition, idx int, arg string) bool {
	switch cond.cond {
	case MustBeUppercase:
		if !textIsUpper(arg) {
			return false
		}
	case MustBeLowercase:
		if !textIsLower(arg) {
			return false
		}
	case MustHavePrefix:
		if !textIsPrefix(arg, cond.value.(string)) {
			return false
		}
	case MustHaveSufix:
		if !textIsSuffix(arg, cond.value.(string)) {
			return false
		}
	case MustCompileRegex:
		if !textIsRegMatch(arg, cond.value.(string)) {
			return false
		}
	}
	return true
}

func validateConditions(command *Command, idx int, arg string) bool {
	var status bool = true
	for _, cond := range command.conditions[idx] {
		switch cond.typeIdx {
		case typePlayer:
			ok := verifyTypePlayer(cond, idx, arg)
			log.Printf("[validateConditions] typePlayer is valid? %v", ok)
			if !ok {
        status = false
				// return false
			}
		case typeNumber:
			ok := verifyTypeNumber(cond, idx, arg)
			log.Printf("[validateConditions] typeNumber is valid? %v", ok)
			if !ok {
        status = false
				// return false
			}
		case typeText:
			ok := verifyTypeText(cond, idx, arg)
			log.Printf("[validateConditions] typeText is valid? %v", ok)
			if !ok {
				status = false
				// return false
			}
		}
	}
	return status
}

func validateArgs(command *Command, args []string) bool {
	if len(command.conditions) > 0 && len(args) <= 0 {
		return false
	}

	for idx, arg := range args {
		ok := validateConditions(command, idx, arg)
		log.Printf("[validateArgs] validateConditions is valid? - %v", ok)
		if !ok {
			return false
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

	log.Printf("[processCommand] args is valid? - %v", isValidArgs)
	if !isValidArgs {
		return false
	}

	// log.Printf("[rakstar] running command [%s] for player %s\n", command.Name, player.GetName())

	command.Handler(&context)

	return true
}
