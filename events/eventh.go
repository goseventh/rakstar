package events

import "github.com/goseventh/rakstar/internal/callbacks"

type Event struct {
	handler interface{}
	event   string
}

func NewEvent() *Event {
	return new(Event)
}

func (e *Event) SetEvent(event string) {
	e.event = event
}

func (e *Event) SetHandler(event string) {
	e.event = event
}

func (e *Event) Subscribe() error {
	err := callbacks.On(e.event, e.handler)
	return err
}
