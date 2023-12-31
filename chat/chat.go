package chat

import (
	"fmt"
	"strings"
	"time"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/utils/constants/playerConst"
	"github.com/goseventh/rakstar/internal/utils/sampstr"
)

type SendPlayerMessageRequest struct {
	Player          *natives.Player
	Message         string
	Color           string
	Local           bool
	Range           float32
	Tag             string
	EnableEncoding bool
}

type ChatBuilder struct {
	requestMsg SendPlayerMessageRequest
}

var isChatEnable = true

func Builder() *ChatBuilder {
  chat := new(ChatBuilder)
  chat.EnableEncodding()
	return chat
}
func (chat *ChatBuilder) Wait(wait ...time.Duration) *ChatBuilder {
	if wait[0].Seconds() < 1 {
		wait[0] = time.Second
	}
	time.Sleep(wait[0])
	return chat
}

func (chat *ChatBuilder) Select(playerid int) *ChatBuilder {
	player := natives.Player{ID: playerid}
	chat.requestMsg.Player = &player
	return chat
}

func (chat *ChatBuilder) Message(msg string) *ChatBuilder {
	chat.requestMsg.Message = msg
	if chat.requestMsg.EnableEncoding {
		chat.requestMsg.Message = sampstr.Encode(chat.requestMsg.Message)
	}
	return chat
}

// EnableEncodding ativará a codificação necessária
// para converter seu texto em um texto compatível com acentuações
// samp. Tentar codificar uma mensagem recebida diretamente do samp
// causará comportamentos estranhos, use somente se você utilizar uma
// string Go. O padrão para EnableEncodding é true
//
// OBS: Esta função deve ser invocada antes de invocar
// função message
func (chat *ChatBuilder) EnableEncodding() *ChatBuilder {
	chat.requestMsg.EnableEncoding = true
	return chat
}

// DisableEncodding desativará codificação necessária
// para converter seu texto em um texto compatível com acentuações
//
// OBS: Esta função deve ser invocada antes de invocar
// a função message
// samp. Consulte:
// - https://pkg.go.dev/github.com/goseventh/rakstar/player#chat.EnableEncoding
func (chat *ChatBuilder) DisableEncodding() *ChatBuilder {
	chat.requestMsg.EnableEncoding = false
	return chat
}

func (chat *ChatBuilder) Color(color string) *ChatBuilder {
	chat.requestMsg.Color = color
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

	if chat.requestMsg.Player == nil {
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

	if chat.requestMsg.Player.ID == Global {
		natives.SendClientMessageToAll(-1, chat.requestMsg.Message)
	}

	switch chat.requestMsg.Range {
	case Local:
		chat.requestMsg.Range = PotencyLocal
		err := sendRange(chat)
		if err != nil {
			fmt.Println(err)
		}
	case Grito:
		chat.requestMsg.Range = PotencyGrito
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

	for playerID := 0; playerID < playerConst.MaxPlayers; playerID++ {
		if !natives.IsPlayerConnected(playerID) {
			continue
		}

		if !natives.IsPlayerInRangeOfPoint(playerID, chat.requestMsg.Range, x, y, z) {
			continue
		}

		natives.SendClientMessage(playerID, -1, chat.requestMsg.Message)
	}
	return err
}

func Active() {
	print("active")
}

func Disable() {}

func Flush() {
	natives.SendClientMessageToAll(-1, " ")
}
