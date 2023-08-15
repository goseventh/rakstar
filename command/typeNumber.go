package command

const (
	MustEqual = iota
	MustBeDivisibleBy
	MustBeGreaterThan
	MustBeLessThan
	MustBeBetween
	MustBeMultipleOf
	MustBeSquareRootOf
)

type TypeNumber struct {
	c *conditionalsBuilder
}

func (c *conditionalsBuilder) TypeNumber() *TypeNumber {
	c.typeIdx = typeNumber
	tNumber := new(TypeNumber)
	tNumber.c = c
	return tNumber
}

func (t *TypeNumber) MustEqual(x int) *TypeNumber {
	t.c.createConditional(MustEqual, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeDivisibleBy(x int) *TypeNumber {
	t.c.createConditional(MustBeDivisibleBy, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeGreaterThan(x int) *TypeNumber {
	t.c.createConditional(MustBeGreaterThan, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeLessThan(x int) *TypeNumber {
	t.c.createConditional(MustBeLessThan, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeMultipleOf(x int) *TypeNumber {
	t.c.createConditional(MustBeMultipleOf, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeSquareRootOf(x int) *TypeNumber {
	t.c.createConditional(MustBeSquareRootOf, t.c.typeIdx, x)
	return t
}

func (t *TypeNumber) MustBeBetween(min, max int) *TypeNumber {
	var values []int
	values = (append(values, min, max))
	t.c.createConditional(MustBeBetween, t.c.typeIdx, values)
	return t
}

func (t *TypeNumber) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
