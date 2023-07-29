package command

import (
	"strconv"
)



func (self *ArgHandler) Next(defaultValues ...string) *string {
	var defaultValue *string = nil

	if len(defaultValues) > 0 {
		defaultValue = &defaultValues[0]
	}

	if self.currentArg == len(self.args) {
		return defaultValue
	}

	if self.currentArg == len(self.args) {
		return defaultValue
	}

	arg := self.args[self.currentArg]
	self.currentArg++

	return &arg
}

func (self *ArgHandler) NextInt(defaultValues ...int) *int {
	var defaultValue *int = nil

	if len(defaultValues) > 0 {
		defaultValue = &defaultValues[0]
	}

	nextValue := self.Next()

	if nextValue == nil {
		return defaultValue
	}

	converted, err := strconv.ParseInt(*nextValue, 0, 64)

	if err != nil {
		return defaultValue
	}

	asInt := int(converted)

	return &asInt
}

func (self *ArgHandler) NextFloat(defaultValues ...float64) *float64 {
	var defaultValue *float64 = nil

	if len(defaultValues) > 0 {
		defaultValue = &defaultValues[0]
	}

	nextValue := self.Next()

	if nextValue == nil {
		return defaultValue
	}

	converted, err := strconv.ParseFloat(*nextValue, 0)

	if err != nil {
		return defaultValue
	}

	asFloat64 := float64(converted)

	return &asFloat64
}

func (self *ArgHandler) Restore() {
	self.currentArg = 0
}

func (self *ArgHandler) GetInput() *string {
	return &self.input
}
