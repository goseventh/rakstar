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
	c.index = index
	return c
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

func (t *TypePlayer) createConditional(typeCondiction, typeIdx int, value interface{}) {
	cond := condition{
		cond:    typeCondiction,
		typeIdx: t.c.index,
		value:   value,
	}

	t.c.conditions = append(t.c.conditions, cond)
}
