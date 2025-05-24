package widget

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type Paragraph struct {
	text string
}

func (p *Paragraph) Render(screen tcell.Screen, area geometry.Rect) {
	w, h := screen.Size()
	col := area.X
	row := area.Y
	for _, c := range p.text {
		// Move to the next line
		if c == '\n' {
			row++
			col = area.X
			continue
		}

		// For now just hard wrap
		if col > w {
			row++
			col = area.X
		}

		// Ignore output if no more space left
		if row > h {
			break
		}

		screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
		col++
	}
}

func NewParagraph(text string) Paragraph {
	return Paragraph{text: text}
}
