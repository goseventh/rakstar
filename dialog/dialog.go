package dialog

import (
	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/dialogConst"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
)

func (db *DialogBuilder) Select(arg interface{}) *DialogBuilder {
	switch v := arg.(type) {
	case string:
		var name string
		for i := 0; i < playerConst.MaxPlayers; i++ {
			natives.GetPlayerName(i, &name, playerConst.MaxPlayerName)
			if name == v {
				db.DialogRequest.ID = i
			}
			return db

		}
	case int:
		db.DialogRequest.ID = v
		return db

	}

	return db
}

func (db *DialogBuilder) Title(title string) *DialogBuilder {
	db.DialogRequest.Caption = title
	return db
}

func (db *DialogBuilder) Message(msg string) *DialogBuilder {
	db.DialogRequest.Info = msg
	return db
}

func (db *DialogBuilder) Buttons(buttons []string) *DialogBuilder {
	db.DialogRequest.Buttons = buttons
	return db
}

func (db *DialogBuilder) Close() *DialogBuilder {
	natives.ShowPlayerDialog(db.DialogRequest.ID, -1, dialogConst.DialogStyleMsgbox, "???", "???", "???", "???")
	return db
}
