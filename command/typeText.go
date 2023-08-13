package command

const (
	MustBeUppercase = iota
	MustBeLowercase
  MustCompileRegex
)

type TypeText struct {
	c *conditionalsBuilder
}

func (c *conditionalsBuilder) TypeText() *TypeText {
	c.typeIdx = typeText
	tText := new(TypeText)
	tText.c = c
	return tText
}

func (t *TypeText) MustBeUppercaser() *TypeText {
	t.c.createConditional(MustBeUppercase, t.c.index, nil)
	return t
}

func (t *TypeText) MustBeLowercase() *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.index, nil)
	return t
}


func (t *TypeText) MustCompileRegex(regex string) *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.index, regex)
	return t
}

func (t *TypeText) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
