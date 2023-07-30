package command

import (
	"time"
)

type commandBuilder struct {
	cmd           string
	alias         []string
	handler       CommandHandler
	conditionals_ map[int][]tCondils
}

type ArgHandler struct {
	input      string
	args       []string
	currentArg int
}

func Builder() *commandBuilder {
	cb := new(commandBuilder)
	cb.conditionals_ = make(map[int][]tCondils)
	return cb
}

func (cb *commandBuilder) Wait(wait ...time.Duration) *commandBuilder {
	if wait[0] < time.Second {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])

	return cb
}

func (cb *commandBuilder) Conditionals() *conditionalsBuilder {
	condit := new(conditionalsBuilder)
	//condit.conditionals_ = make(map[int][]int)
	condit.cb = cb

	return condit
}
