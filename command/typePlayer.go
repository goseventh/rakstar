package command

const (
	MustPlayerConnected = iota
	MustNickIs
)

type TypePlayer struct {
	c *conditionalsBuilder
}

// TypePlayer define o tipo lógico do parâmetro do comando
// para o tipo Player
func (c *conditionalsBuilder) TypePlayer() *TypePlayer {
	c.typeIdx = typePlayer
	tPlayer := new(TypePlayer)
	tPlayer.c = c
	return tPlayer
}

// MustConnected é válido se o Player estiver conectado
// ao servidor. Caso contrário, o comando falhará
func (t *TypePlayer) MustConnected() *TypePlayer {
	t.c.createConditional(MustPlayerConnected, t.c.typeIdx, nil)
	return t
}

// MustNickIs compara se o nick é igual ao nickname 
// do jogador conectado
func (t *TypePlayer) MustNickIs(nick string) *TypePlayer {
	t.c.createConditional(MustNickIs, t.c.typeIdx, nick)
	return t
}

// Encerra a expressão lógica
func (t *TypePlayer) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
