package command

import (
	"time"
)

type commandBuilder struct {
	cmd     string
	alias   []string
	handler CommandHandler
}

type ArgHandler struct {
	input      string
	args       []string
	currentArg int
}

func Builder() *commandBuilder {

	return new(commandBuilder)
}

func (cb *commandBuilder) Wait(wait ...time.Duration) *commandBuilder {
	if wait[0] < time.Second {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])

	return cb
}

func (cb *commandBuilder) Conditionals() *conditionalsBuilder {
	return new(conditionalsBuilder)
}
