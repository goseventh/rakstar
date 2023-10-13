package command

const (
	undefined = iota
	typePlayer
	typeNumber
	typeText
)

type conditionalsBuilder struct {
	index      int
	typeIdx    int
	conditions []condition
	c          *commandBuilder
}

type condition struct {
	typeIdx int
	cond    int
	value   interface{}
}

func (c *conditionalsBuilder) Index(index int) *conditionalsBuilder {
	if index < 0 {
		index = 0
	}
	c.index = index
	return c
}

func (c *conditionalsBuilder) createConditional(typeCond, typeIdx int, value interface{}) {
	cond := condition{
		cond:    typeCond,
		typeIdx: typeIdx,
		value:   value,
	}

	c.conditions = append(c.conditions, cond)
}

// Set registra as condicionais para cada indice(index)
func (c *conditionalsBuilder) Set() *conditionalsBuilder {
	if c.index < 0 {
		return c
	}

	if c.conditions == nil || len(c.conditions) == 0 {
		return c
	}

	c.c.conditionals[c.index] = c.conditions
	c.conditions = nil
	return c
}

//EndConditionals encerra o bloco lógico atual.
func (t *TypePlayer) EndConditionals() *commandBuilder {
	t.End()
	return t.c.c
}

func (t *TypeNumber) EndConditionals() *commandBuilder {
	t.End()
	return t.c.c
}

func (t *TypeText) EndConditionals() *commandBuilder {
	t.End()
	return t.c.c
}
