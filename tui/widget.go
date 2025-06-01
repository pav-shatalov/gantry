package tui

import (
	"gantry/geometry"
)

type Widget interface {
	Render(buf *ScreenBuffer, area geometry.Rect)
}

type BlockWidget struct{}

type InlineWidget struct{}
