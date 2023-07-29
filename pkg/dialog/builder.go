package dialog

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
}

func Builder() *DialogBuilder {
	return new(DialogBuilder)
}
