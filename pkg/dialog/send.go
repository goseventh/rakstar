package dialog

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/dialogConst"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
	"github.com/goseventh/rakstar/internal/utils/sampstr"
)

func (db *DialogBuilder) Send() *DialogBuilder {
	if db.DialogRequest.ID < 0 || db.DialogRequest.ID > playerConst.MaxPlayers {
		return db
	}

	if db.DialogRequest.Style < dialogConst.DialogStyleMsgbox ||
		db.DialogRequest.Style > dialogConst.DialogStyleTablistHeaders {
		return db
	}

	if len(db.DialogRequest.Buttons) == 0 {
		return db
	}

	var button1 = ""
	var button2 = ""

	switch len(db.DialogRequest.Buttons) {
	case 2:
		button2 = db.DialogRequest.Buttons[1]
		fallthrough
	case 1:
		button1 = db.DialogRequest.Buttons[0]
	}

	natives.ShowPlayerDialog(
		db.DialogRequest.ID,
		dialogHandlerID,
		db.DialogRequest.Style,
		sampstr.Encode(db.DialogRequest.Caption),
		sampstr.Encode(db.DialogRequest.Info),
		sampstr.Encode(button1),
		sampstr.Encode(button2),
	)

	channel := poolPush(db.DialogRequest.ID)

	dialogResponse := <-channel

	if dialogResponse == nil {
		return db
	}

	return db
}
