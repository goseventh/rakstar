package command

const (
	MustBeUppercase = iota
	MustBeLowercase
	MustHavePrefix
	MustHaveSufix
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

func (t *TypeText) MustBeUppercase() *TypeText {
	t.c.createConditional(MustBeUppercase, t.c.typeIdx, nil)
	return t
}

func (t *TypeText) MustBeLowercase() *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.typeIdx, nil)
	return t
}

func (t *TypeText) MustHavePrefix(preffix string) *TypeText {
	t.c.createConditional(MustHavePrefix, t.c.typeIdx, preffix)
	return t
}

func (t *TypeText) MustHaveSufix(sufix string) *TypeText {
	t.c.createConditional(MustHaveSufix, t.c.typeIdx, sufix)
	return t
}

func (t *TypeText) MustCompileRegex(regex string) *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.typeIdx, regex)
	return t
}

func (t *TypeText) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
