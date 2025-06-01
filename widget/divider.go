package widget

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type direction = uint8

const (
	horizontal direction = iota
	vertical
)

type Divider struct {
	direction direction
	r         rune
}

func NewVerticalDivider() Divider {
	return Divider{direction: vertical, r: '│'}
}

func NewHorizontalDivider() Divider {
	return Divider{direction: horizontal, r: '─'}
}

func (s *Divider) Render(screen tcell.Screen, area geometry.Rect) {
	dim := area.Width
	row := area.Y
	col := area.X

	if s.direction == vertical {
		dim = area.Height
	}

	for range dim {
		screen.SetContent(col, row, s.r, []rune{}, tcell.StyleDefault.Foreground(tcell.ColorBlack))
		if s.direction == vertical {
			row++
		} else {
			col++
		}
	}
}
