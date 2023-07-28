package rakstar

import (
	"fmt"
	"strings"
)

type SendPlayerMessageRequest struct {
	Player          *Player
	Message         string
	Color           string
	Local           bool
	Range           float32
	Tag             string
	DisableEncoding bool
}

type SendMessageRequest struct {
	Player          *Player
	Message         string
	Color           string
	Tag             string
	DisableEncoding bool
}

var isChatEnable = true

func CreateMessage(sendPlayerMessageRequest *SendPlayerMessageRequest) {
	var message string
	var tag string = ""

	if sendPlayerMessageRequest.Message == "" {
		return
	}

	if sendPlayerMessageRequest.Player == nil {
		return
	}

	if sendPlayerMessageRequest.Color == "" {
		sendPlayerMessageRequest.Color = "{ffffff}"
	}

	if sendPlayerMessageRequest.Local && sendPlayerMessageRequest.Range <= 0.0 {
		sendPlayerMessageRequest.Range = 10.0
	}

	if sendPlayerMessageRequest.Tag != "" {
		tag = fmt.Sprintf("[%s] ", strings.ToUpper(sendPlayerMessageRequest.Tag))
	}

	message =
		fmt.Sprintf(
			"%s%s%s: {ffffff}%s",
			sendPlayerMessageRequest.Color,
			tag,
			sendPlayerMessageRequest.Player.GetName(),
			sendPlayerMessageRequest.Message,
		)

	if !sendPlayerMessageRequest.DisableEncoding {
		message = Encode(message)
	}

	if sendPlayerMessageRequest.Local {
		x, y, z, err := sendPlayerMessageRequest.Player.GetPos()

		if err != nil {
			fmt.Println(err)
			return
		}

		for playerID := 0; playerID < MaxPlayers; playerID++ {
			if !IsPlayerConnected(playerID) {
				continue
			}

			if !IsPlayerInRangeOfPoint(playerID, sendPlayerMessageRequest.Range, x, y, z) {
				continue
			}

			SendClientMessage(playerID, -1, message)
		}

		return
	}

	SendClientMessageToAll(-1, message)
}

func SendMessage(sendMessageRequest *SendMessageRequest) {
	var message string
	var tag string = ""

	if sendMessageRequest.Message == "" {
		return
	}

	if sendMessageRequest.Color == "" {
		sendMessageRequest.Color = "{ffffff}"
	}

	if sendMessageRequest.Tag != "" {
		tag = fmt.Sprintf("[%s] ", strings.ToUpper(sendMessageRequest.Tag))
	}

	message =
		fmt.Sprintf(
			"%s%s%s",
			sendMessageRequest.Color,
			tag,
			sendMessageRequest.Message,
		)

	if !sendMessageRequest.DisableEncoding {
		message = Encode(message)
	}

	if sendMessageRequest.Player != nil {
		SendClientMessage(sendMessageRequest.Player.ID, -1, message)
		return
	}

	SendClientMessageToAll(-1, message)
}

func Active() {}

func Disable() {}

func Flush() {
	SendClientMessageToAll(-1, " ")
}
