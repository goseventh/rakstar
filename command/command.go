// command fornece funções úteis para criação e manipulações de comandos do SA-MP.
// Criar comandos com este pacote é intuitivo, e permite operações lógicas
// utilizando os operadoes lógicos do pacote, para estabelecer regras e condições
// antes de executar o comando. 
package command

import (
	"reflect"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/goseventh/rakstar/internal/natives"
)

type CommandHandler func(context *CommandContext)

type Command struct {
	Name          string
  conditions map[int][]condition
	Handler       CommandHandler
	Interceptors  []CommandInterceptorHandler
	Description   string
	RequireLogin  bool
	Aliases       []string
	Roles         []string
}

type CommandContext struct {
	Player     *natives.Player
	ArgHandler *ArgHandler
}

var commands = make(map[string]*Command)

func (self *CommandInterceptorContext) Next() {
	self.next = true
}

func RegisterCommand(command *Command) (bool, error) {
	commandToRegist := commands[strings.ToLower(command.Name)]

	if commandToRegist != nil {
		command.Interceptors = commandToRegist.Interceptors
		commandToRegist = command
	}

	if commandToRegist == nil {
		commandToRegist = command
	}

	commands[strings.ToLower(command.Name)] = commandToRegist

	for _, name := range commandToRegist.Aliases {
		commands[strings.ToLower(name)] = commandToRegist
	}

	return true, nil
}

func SearchCommand(inputName string) (*Command, int) {
	inputName = strings.ToLower(inputName)

	commandsKeys := reflect.ValueOf(commands).MapKeys()

	var lastCommand *Command = nil
	var lastDistance = -1

	for _, commandName := range commandsKeys {
		command := commands[commandName.String()]

		if command.Handler == nil {
			continue
		}

		distance := levenshtein.ComputeDistance(commandName.String(), inputName)

		if lastCommand == nil || lastDistance == -1 || distance < lastDistance {

			lastDistance = distance
			lastCommand = command

			if distance == 0 {
				break
			}
		}
	}

	if lastDistance == -1 || lastDistance > 2 {
		return nil, lastDistance
	}

	return lastCommand, lastDistance
}
