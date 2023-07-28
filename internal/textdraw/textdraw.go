package textDraw

import (
	"fmt"

	"github.com/goseventh/rakstar/internal/natives"
	"github.com/goseventh/rakstar/internal/player"
	"github.com/goseventh/rakstar/internal/utils/constants/textDrawConst"
)

const (
	FontSanAndreas = iota
	FontClear
	FontCapitalClear
	FontGTA
	FontSprite
)

type PlayerTextDraw struct {
	player   *player.Player
	textDraw int
	align    int
}

func (p *player.Player) NewTextDraw(x, y float32, text string) (PlayerTextDraw, error) {
	td := PlayerTextDraw{player: p, textDraw: natives.CreatePlayerTextDraw(p.ID, x, y, text)}
	if td.textDraw == textDrawConst.InvalidTextDraw {
		return td, fmt.Errorf("invalid playertextdraw")
	}
	return td, nil
}

func (p *PlayerTextDraw) Destroy() {
	natives.PlayerTextDrawDestroy(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) SetString(text string) {
	natives.PlayerTextDrawSetString(p.player.ID, p.textDraw, text)
}

func (p *PlayerTextDraw) Show() {
	natives.PlayerTextDrawShow(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Hide() {
	natives.PlayerTextDrawHide(p.player.ID, p.textDraw)
}

func (p *PlayerTextDraw) Font(font int) {
	natives.PlayerTextDrawFont(p.player.ID, p.textDraw, font)
}

func (p *PlayerTextDraw) UseBox(use bool) {
	natives.PlayerTextDrawUseBox(p.player.ID, p.textDraw, use)
}

func (p *PlayerTextDraw) SetAlignment(align int) {
	p.align = align
	natives.PlayerTextDrawAlignment(p.player.ID, p.textDraw, p.align)
}

func (p *PlayerTextDraw) SetTextSize(x, y float32) {
	if p.align == 2 {
		x, y = y, x
	}
	natives.PlayerTextDrawTextSize(p.player.ID, p.textDraw, x, y)
}

func (p *PlayerTextDraw) SetColor(color int) {
	natives.PlayerTextDrawColor(p.player.ID, p.textDraw, color)
}

var SetColour = (*PlayerTextDraw).SetColor

func (p *PlayerTextDraw) SetBoxColor(color int) {
	natives.PlayerTextDrawBoxColor(p.player.ID, p.textDraw, color)
}

var SetBoxColour = (*PlayerTextDraw).SetBoxColor

func (p *PlayerTextDraw) SetBackgroundColor(color int) {
	natives.PlayerTextDrawBackgroundColor(p.player.ID, p.textDraw, color)
}

var SetBackgroundColour = (*PlayerTextDraw).SetBackgroundColor
