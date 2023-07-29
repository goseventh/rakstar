package dialog

func (db *DialogBuilder) GetResponse(response *int) *DialogBuilder {
	*response = db.DialogResponse.Response
	return db
}

func (db *DialogBuilder) GetListItem(listItem *int) *DialogBuilder {
	*listItem = db.DialogResponse.Listitem
	return db
}

func (db *DialogBuilder) GetText(text *string) *DialogBuilder {
	*text = db.DialogResponse.Inputtext
	return db
}

func (db *DialogBuilder) GetErr(err *error) *DialogBuilder {
	*err = db.Err
	return db
}
