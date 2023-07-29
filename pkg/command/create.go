package command

import "github.com/goseventh/rakstar/internal/command"

func (cb *commandBuilder) Command(cmd string) *commandBuilder {
	cb.cmd = cmd
	return cb
}

func (cb *commandBuilder) Alias(cmd string) *commandBuilder {
	cb.alias = append(cb.alias, cmd)
	return cb
}

func (cb *commandBuilder) Handler(handler command.CommandHandler) *commandBuilder {
	cb.handler = handler
	return cb
}

func (cb *commandBuilder) Create() {
	command.RegisterCommand(&command.Command{
		Name: cb.cmd,
		Aliases: cb.alias,
		Handler: cb.handler,
	})
	
}