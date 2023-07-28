package rakstar



func HandlePlayerText(player Player, text string) bool {
	if !isChatEnable {
		return false
	}

	if text == "" {
		return false
	}

	SendPlayerMessage(&SendPlayerMessageRequest{
		Player:          &player,
		Message:         text,
		Color:           ChatLocalColorStr,
		Range:           15.0,
		Local:           true,
		Tag:             "local",
		DisableEncoding: true,
	})

	return false
}
