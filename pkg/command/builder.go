package command

import (
	"time"

	"github.com/goseventh/rakstar/internal/command"
	"github.com/goseventh/rakstar/internal/utils/common"
	"github.com/goseventh/rakstar/pkg/chat"
)

type commandBuilder struct {
	cmd     string
	alias   []string
	handler command.CommandHandler
}

func Builder() *commandBuilder {
	cb := new(commandBuilder)
	chat := chat.Builder()
	chat.Message("Nenhum comando correspondente foi encontrado!").
		Color(common.ErrorColorStr).
		Tag("servidor")

	cb.SetFail(chat)
	return new(commandBuilder)
}

func (cb *commandBuilder) Wait(wait ...time.Duration) *commandBuilder {
	if wait[0] < time.Second {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])

	return cb
}
