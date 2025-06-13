package paragraph

import (
	"gantry/geometry"
	"gantry/tui"

	"github.com/mattn/go-runewidth"
)

type Paragraph struct {
	lines  []string
	scroll int
}

func (p *Paragraph) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	remainingRows := area.Height
	idx := 0

	for y := p.scroll; y < len(p.lines); y++ {
		if remainingRows == 0 {
			break
		}

		remainingCols := area.Width
		for x, r := range p.lines[y] {
			if remainingCols == 0 {
				break
			}

			buf.SetContent(x+area.Col, area.Row+idx, r, tui.StyleDefault)
			w := max(runewidth.RuneWidth(r), 1)
			remainingCols -= w
			x += w
		}
		remainingRows--
		idx++
	}
}

func New(lines []string) Paragraph {
	return Paragraph{lines: lines, scroll: 0}
}

func (p Paragraph) Scroll(s int) Paragraph {
	p.scroll = s
	return p
}
