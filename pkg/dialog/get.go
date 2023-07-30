package dialog

func (db *DialogBuilder) GetButton(response *string) *DialogBuilder {
	button := ""
	if db.DialogResponse.Response == 1 {
		button = "left"
	} else if db.DialogResponse.Response == 2 {
		button = "right"
	}
	*response = button
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
