package paragraph

import (
	"gantry/geometry"
	"gantry/tui"
)

type Paragraph struct {
	text   string
	scroll geometry.Position
}

func (p *Paragraph) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	if len(p.text) == 0 {
		return
	}

	tmpBuffer := tui.NewAutogrowingBuffer(area.Width, 1)

	col := 0
	row := 0
	remainingCols := area.Width

	for _, r := range p.text {
		if r == '\n' {
			row++
			col = 0
			remainingCols = area.Width
			continue
		}
		if remainingCols == 0 {
			continue
		}
		tmpBuffer.SetContent(col, row, r, tui.StyleDefault)
		col++
		remainingCols--
	}

	buf.FillFrom(&tmpBuffer, area)

	// m := fmt.Sprintf("Size: %dx%d", tmpBuffer.Width(), tmpBuffer.Height())
	// col := area.X
	// row := area.Y
	// for _, r := range m {
	// 	buf.SetContent(col, row, r, tui.StyleDefault)
	// 	col++
	// }

	// for _, c := range p.text {
	// 	// Move to the next line
	// 	if c == '\n' {
	// 		row++
	// 		col = area.X
	// 		continue
	// 	}
	//
	// 	// For now just hard wrap
	// 	if col > w {
	// 		row++
	// 		col = area.X
	// 	}
	//
	// 	// Ignore output if no more space left
	// 	if row > h {
	// 		break
	// 	}
	//
	// 	screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
	// 	col++
	// }
}

func New(text string) Paragraph {
	return Paragraph{text: text, scroll: geometry.Position{X: 0, Y: 0}}
}
