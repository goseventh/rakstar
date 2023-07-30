package dialog

import "time"

const (
	InvalidStyleDialogError     = "dialog invalid style"
	InvalidPlayerIDDialogError  = "dialog invalid player id"
	ResponseTimeoutDialogError  = "dialog response timeout"
	ButtonIsRequiredDialogError = "at least one dialog button is required"
)

type DialogRequest struct {
	ID      int
	Style   int
	Caption string
	Info    string
	Buttons []string
}

type DialogResponse struct {
	Response  int
	Listitem  int
	Inputtext string
}

const dialogHandlerID = 9999

type DialogBuilder struct {
	DialogRequest  *DialogRequest
	DialogResponse *DialogResponse
	Err error
}

func  Builder() *DialogBuilder {
	return new(DialogBuilder)
}

func (db *DialogBuilder) Wait(wait... time.Duration) *DialogBuilder {
	if wait[0].Seconds() < 1 {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])
	return new(DialogBuilder)
}

