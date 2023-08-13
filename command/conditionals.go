package command

const (
	typePlayer = iota
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

func (c *conditionalsBuilder) createConditional(typeCondiction, typeIdx int, value interface{}) {
	cond := condition{
		typeIdx: c.index,
		value:   value,
	}
	c.conditions = append(c.conditions, cond)
}

// registra as condicionais para cada indice(index)
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

func (t *TypePlayer) EndConditionals() *commandBuilder {
	t.End()
	return t.c.c
}
