package command

const (
	MustPlayerConnected = iota
)

type conditionalsBuilder struct {
	index         int
	typeIdx       int
	condils       []int
	conditionals_ map[int][]int
}

func (cb *conditionalsBuilder) Index(index int) *conditionalsBuilder {
	cb.index = index
	return cb
}

func (cb *conditionalsBuilder) Type(typeIdx int) *conditionalsBuilder {
	cb.typeIdx = typeIdx
	return cb
}

func (cb *conditionalsBuilder) MustPlayerConnected() *conditionalsBuilder {
	cb.condils = append(cb.condils, MustPlayerConnected)
	return cb
}

func (cb *conditionalsBuilder) Set() *conditionalsBuilder {
	cb.condils = nil
	cb.conditionals_[cb.index] = cb.condils
	return cb
}
