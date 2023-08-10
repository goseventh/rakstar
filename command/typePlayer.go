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
	cond := condition{
		cond:    MustPlayerConnected,
		typeIdx: t.c.index,
	}

	t.c.conditions = append(t.c.conditions, cond)
	return t
}

func (t *TypePlayer) MustNickIs(nick string) *TypePlayer {
	cond := condition{
		cond:    MustNickIs,
		typeIdx: t.c.index,
		value:   nick,
	}

	t.c.conditions = append(t.c.conditions, cond)
	return t
}

func (t *TypePlayer) End() *conditionalsBuilder {
	t.c.Set()
	return t.c
}
