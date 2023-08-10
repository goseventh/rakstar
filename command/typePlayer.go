package command

const (
	MustPlayerConnected = iota
	MustNickIs
)

type TypePlayer struct {
	c *conditionalsBuilder
}

func (c *conditionalsBuilder) TypePlayer() *TypePlayer {
	c.typeIdx = typePlayer
	tPlayer := new(TypePlayer)
	tPlayer.c = c

	return tPlayer
}

func (t *TypePlayer) MustConnected() *TypePlayer {
 t.c.createConditional(MustPlayerConnected, t.c.index, nil)

	return t

}

func (t *TypePlayer) MustNickIs(nick string) *TypePlayer {
  t.c.createConditional(MustNickIs, t.c.index, nick)
	return t
}

func (t *TypePlayer) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
