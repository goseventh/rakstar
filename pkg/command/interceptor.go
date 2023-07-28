package command

import (
	"strings"
	"github.com/goseventh/rakstar"
)

type CommandInterceptorContext struct {
	Player     *rakstar.Player
	ArgHandler *ArgHandler
	next       bool
}

type CommandInterceptorHandler func(context *CommandInterceptorContext)

func RegisterInterceptor(commandName string, interceptor CommandInterceptorHandler) {
	command := commands[strings.ToLower(commandName)]

	if command == nil {
		command = &Command{
			Name:         commandName,
			Interceptors: []CommandInterceptorHandler{interceptor},
		}

		commands[strings.ToLower(commandName)] = command

		return
	}

	command.Interceptors = append(command.Interceptors, interceptor)
}
