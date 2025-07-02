package paragraph

import (
	"gantry/tui"

	"github.com/rivo/uniseg"
)

type Paragraph struct {
	tui.Block
	lines  []string
	scroll int
}

func (p *Paragraph) Render(buf *tui.OutputBuffer, area tui.Rect) {
	p.Block.Render(buf, area)
	a := p.InnerArea(area)
	remainingRows := a.Height
	idx := 0

	for y := p.scroll; y < len(p.lines); y++ {
		if remainingRows == 0 {
			break
		}

		remainingCols := a.Width
		line := uniseg.NewGraphemes(p.lines[y])
		x := 0
		for line.Next() {
			r := line.Runes()
			if remainingCols == 0 {
				break
			}

			buf.SetContent(x+a.Col, idx+a.Row, r[0], tui.StyleDefault)
			w := line.Width()
			remainingCols -= w
			x += w
		}
		remainingRows--
		idx++
	}
}

func New(lines []string) Paragraph {
	return Paragraph{lines: lines, scroll: 0, Block: tui.NewBlock()}
}

func (p *Paragraph) Scroll(s int) {
	p.scroll = s
}
