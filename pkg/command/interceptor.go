package command

import (
	"strings"

	"github.com/goseventh/rakstar/internal/natives"
)

type CommandInterceptorContext struct {
	Player     *natives.Player
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
