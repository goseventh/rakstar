package command

import (
	"time"
)

type commandBuilder struct {
	cmd                 string
	alias               []string
	handler             CommandHandler
	conditionals        map[int][]condition
	conditionalsBuilder *conditionalsBuilder
}

type ArgHandler struct {
	input      string
	args       []string
	currentArg int
}

func Builder() *commandBuilder {
	c := new(commandBuilder)
	c.conditionals = make(map[int][]condition)
	return c
}

func (c *commandBuilder) Wait(wait ...time.Duration) *commandBuilder {
	if wait[0] < time.Second {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])

	return c
}

func (c *commandBuilder) Conditionals() *conditionalsBuilder {
	if c.conditionalsBuilder != nil {
		return c.conditionalsBuilder
	}

	condit := new(conditionalsBuilder)
	condit.c = c
	c.conditionalsBuilder = condit
	return condit
}
