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

// TypeText define o tipo lógico do parâmetro para
// o tipo Text
func (c *conditionalsBuilder) TypeText() *TypeText {
	c.typeIdx = typeText
	tText := new(TypeText)
	tText.c = c
	return tText
}

func (t *TypeText) MustBeUppercaser() *TypeText {
// MustBeUppercase é válido se o texto estiver em caixa
// alta. Caso contrário o comando falhará
func (t *TypeText) MustBeUppercase() *TypeText {
	t.c.createConditional(MustBeUppercase, t.c.typeIdx, nil)
	return t
}

// MustBeLowercase é valido se o texto estiver em caixa
// baixa. Caso contrário o comando falhará
func (t *TypeText) MustBeLowercase() *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.typeIdx, nil)
	return t
}

func (t *TypeText) MustHavePrefix() *TypeText {
	t.c.createConditional(MustHavePrefix, t.c.typeIdx, nil)
	return t
}

func (t *TypeText) MustHaveSufix() *TypeText {
	t.c.createConditional(MustHaveSufix, t.c.typeIdx, nil)
	return t
}

func (t *TypeText) MustCompileRegex(regex string) *TypeText {
	t.c.createConditional(MustBeLowercase, t.c.typeIdx, regex)
	return t
}

// MustHavePrefix compara o preffix com o início do
// parâmetro, se ambos são iguais então é válido.
// Caso contrário, o comando falhará
func (t *TypeText) MustHavePrefix(preffix string) *TypeText {
	t.c.createConditional(MustHavePrefix, t.c.typeIdx, preffix)
	return t
}

// MustHaveSufix compara o sufix com o fim do parâmetro,
// se ambos são iguais então é válido. Caso contrário
// o comando falhará
func (t *TypeText) MustHaveSufix(sufix string) *TypeText {
	t.c.createConditional(MustHaveSufix, t.c.typeIdx, sufix)
	return t
}

// MustCompileRegex compara o regex com o parâmetro, se
// a expressão for verdadeira então é válido. Caso
// contrário o comando falhará
func (t *TypeText) MustCompileRegex(regex string) *TypeText {
	t.c.createConditional(MustCompileRegex, t.c.typeIdx, regex)
	return t
}

// Encerra a expressão lógica
func (t *TypeText) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
