package paragraph

import "github.com/gdamore/tcell/v2"

type Paragraph struct {
	text string
}

func (p *Paragraph) Render(screen tcell.Screen) {
	w,h := screen.Size()
	col := 0
	row := 0
	for _, c := range p.text {
		// Move to the next line
		if (c == '\n') {
			row++
			col = 0
			continue
		}

		// For now just hard wrap
		if (col > w) {
			row++;
			col = 0;
		}

		// Ignore output if no more space left
		if (row > h) {
			break;
		}

		screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
		col++
	}
}

func New(text string) Paragraph {
	return Paragraph{text: text}
}
