package dialog

import (
	"main/pkg/utils/sampstr"
)

func HandleDialogResponse(id int, dialogID int, response int, listitem int, inputtext string) bool {
	if dialogID != dialogHandlerID {
		return false
	}

	channel := poolNext(id)

	if channel == nil {
		return false
	}

	channel <- &DialogResponse{response, listitem, sampstr.Decode(inputtext)}

	return true
}
