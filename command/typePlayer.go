package command

const (
	MustPlayerConnected = iota
	MustNickIs
)

type TPlayer struct {
	cb *conditionalsBuilder
}

func (cb *conditionalsBuilder) TypePlayer() *TPlayer {
	cb.typeIdx = typePlayer
	tPlayer := new(TPlayer)
	tPlayer.cb = cb

	return tPlayer
}

func (tp *TPlayer) MustConnected() *TPlayer {
	cond := tCondils{
		cond:    MustPlayerConnected,
		typeIdx: tp.cb.index,
	}

	tp.cb.condils = append(tp.cb.condils, cond)
	return tp
}

func (tp *TPlayer) MustNickIs(nick string) *TPlayer {
	cond := tCondils{
		cond:    MustNickIs,
		typeIdx: tp.cb.index,
		value:   nick,
	}

	tp.cb.condils = append(tp.cb.condils, cond)
	return tp
}

func (tp *TPlayer) End() *conditionalsBuilder {
	tp.cb.Set()
	return tp.cb
}
