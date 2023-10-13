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

/* Builder é um construtor e deve ser usado
 para instanciar um construtor:
  func(){
      command.Builder().
        Command("hello").
        Alias("Hi").
        Handler(func(context *CommandContext){
          ...
        }).
        Conditionals().
        Index(0).
        TypePlayer().
        MustConnected().
        EndConditionals().
        ...
    }*/
func Builder() *commandBuilder {
	c := new(commandBuilder)
	c.conditionals = make(map[int][]condition)
	return c
}

// Wait é um operador lógico que aguardará N time.Duration [time]. 
// Em um dos testes da equipe RakStar, detectamos um comportamento
// estranho. Evite utilizar este operador 
func (c *commandBuilder) Wait(wait ...time.Duration) *commandBuilder {
	if wait[0] < time.Second {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])

	return c
}

// Conditionals é um construtor de expressões lógicas
func (c *commandBuilder) Conditionals() *conditionalsBuilder {
	if c.conditionalsBuilder != nil {
		return c.conditionalsBuilder
	}

	condit := new(conditionalsBuilder)
	condit.c = c
	c.conditionalsBuilder = condit
	return condit
}
