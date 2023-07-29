package command

import (
	"github.com/goseventh/rakstar/internal/command"
	"github.com/goseventh/rakstar/pkg/chat"
)

func (cb *commandBuilder) SetFail(chat *chat.ChatBuilder) *commandBuilder {
	if chat == nil{
		return cb
	}
	command.Fail = chat
	return cb
}
