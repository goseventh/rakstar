package chat

import (
	"fmt"
	"strings"

	"github.com/goseventh/rakstar/internal/player"
)

type SendPlayerMessageRequest struct {
	Player          *player.Player
	Message         string
	Color           string
	Local           bool
	Range           float32
	Tag             string
	DisableEncoding bool
}

type ChatBuilder struct {
	requestMsg SendPlayerMessageRequest
}

var isChatEnable = true

func Builder() *ChatBuilder {
	return new(ChatBuilder)
}

func (chat *ChatBuilder) PlayerID(playerid int) *ChatBuilder {
	player := player.Player{ID: playerid}
	chat.requestMsg.Player = &player
	return chat
}

func (chat *ChatBuilder) Message(msg string) *ChatBuilder {
	chat.requestMsg.Message = msg
	return chat
}

func (chat *ChatBuilder) Tag(tag string) *ChatBuilder {
	chat.requestMsg.Tag = tag
	return chat
}

func (chat *ChatBuilder) Range(r float32) *ChatBuilder {
	chat.requestMsg.Range = r
	return chat
}

func (chat *ChatBuilder) Send() *ChatBuilder {
	if chat.requestMsg.Message == "" {
		return chat
	}

	if *chat.requestMsg.Player == nil {
		return chat
	}

	if chat.requestMsg.Color == "" {
		chat.requestMsg.Color = "{ffffff}"
	}

	if chat.requestMsg.Tag != "" {
		chat.requestMsg.Tag = fmt.Sprintf("[%s] ", strings.ToUpper(chat.requestMsg.Tag))
	}

	chat.requestMsg.Message =
		fmt.Sprintf(
			"%s%s%s: {ffffff}%s",
			chat.requestMsg.Color,
			chat.requestMsg.Tag,
			chat.requestMsg.Player.GetName(),
			chat.requestMsg.Message,
		)

	if !chat.requestMsg.DisableEncoding {
		chat.requestMsg.Message = Encode(chat.requestMsg.Message)
	}

	switch chat.requestMsg.Range {
	case Global:
		SendClientMessageToAll(-1, chat.requestMsg.Message)
	case Local:
		chat.requestMsg.Range = 15
		err := sendRange(chat)
		if err != nil {
			fmt.Println(err)
		}
	case Grito:
		chat.requestMsg.Range = 45
		err := sendRange(chat)
		if err != nil {
			fmt.Println(err)
		}
	default:
		err := sendRange(chat)
		if err != nil {
			fmt.Println(err)
		}
	}

	return chat

}

func sendRange(chat *ChatBuilder) error {
	x, y, z, err := chat.requestMsg.Player.GetPos()

	for playerID := 0; playerID < MaxPlayers; playerID++ {
		if !IsPlayerConnected(playerID) {
			continue
		}

		if !IsPlayerInRangeOfPoint(playerID, chat.requestMsg.Range, x, y, z) {
			continue chat
		}

		SendClientMessage(playerID, -1, chat.requestMsg.Message)
	}
}

func Active() {
	print("active")
}

func Disable() {}

func Flush() {
	SendClientMessageToAll(-1, " ")
}
