package divider

import (
	"gantry/tui"
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

func NewVertical() Divider {
	return Divider{direction: vertical, r: '│'}
}

func NewHorizontal() Divider {
	return Divider{direction: horizontal, r: '─'}
}

func (s *Divider) Render(buf *tui.OutputBuffer, area tui.Rect) {
	dim := area.Width
	row := area.Row
	col := area.Col

	if s.direction == vertical {
		dim = area.Height
	}

	for range dim {
		buf.SetContent(col, row, s.r, tui.StyleDefault.Fg(tui.ColorBlack))
		if s.direction == vertical {
			row++
		} else {
			col++
		}
	}
}
