package dialog

import "github.com/goseventh/rakstar/internal/utils/constants/dialogConst"

func (db *DialogBuilder) TypeInput() *DialogBuilder {
	db.DialogRequest.Style = dialogConst.DialogStyleInput
	return db
}

func (db *DialogBuilder) TypeList() *DialogBuilder {
	db.DialogRequest.Style = dialogConst.DialogStyleList
	return db
}

func (db *DialogBuilder) TypePassword() *DialogBuilder {
	db.DialogRequest.Style = dialogConst.DialogStylePassword
	return db
}

func (db *DialogBuilder) TypeTabList() *DialogBuilder {
	db.DialogRequest.Style = dialogConst.DialogStyleTablist
	return db
}

func (db *DialogBuilder) TypeTabListHeaders() *DialogBuilder {
	db.DialogRequest.Style = dialogConst.DialogStyleTablistHeaders
	return db
}
