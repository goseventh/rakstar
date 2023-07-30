package command

const (
	typePlayer = iota
	typeNumber
	typeText
)

type conditionalsBuilder struct {
	index   int
	typeIdx int
	condils []tCondils
	//	conditionals_ map[int][]int
	cb *commandBuilder
}

type tCondils struct {
	typeIdx int
	cond    int
	value   interface{}
}

func (cb *conditionalsBuilder) Index(index int) *conditionalsBuilder {
	cb.index = index
	return cb
}

// registra as condicionais para cada indice(index)
func (cb *conditionalsBuilder) Set() *conditionalsBuilder {

	if cb.index < 0 {
		return cb
	}

	if cb.condils == nil || len(cb.condils) == 0 {
		return cb
	}

	//cb.conditionals_s //alocado

	cb.cb.conditionals_[cb.index] = cb.condils
	cb.condils = nil
	return cb
}

func (tp *TPlayer) EndConditionals() *commandBuilder {
	tp.End()
	return tp.cb.cb
}
