package paragraph

import (
	"gantry/tui"
	"math"

	"github.com/rivo/uniseg"
)

type Paragraph struct {
	tui.Block
	lines  []string
	scroll int
}

func (p *Paragraph) Render(buf *tui.OutputBuffer, area tui.Rect) {
	p.Block.Render(buf, area)
	contentLength := len(p.lines)
	if contentLength > p.InnerArea(area).Height {
		scrollBarArea := tui.NewRect(area.Col, area.Row+1, area.Width, area.Height-2)
		renderScrollbar(buf, scrollBarArea, p.scroll, contentLength)
	}
	a := p.InnerArea(area)
	remainingRows := a.Height
	idx := 0

	for y := p.scroll; y < len(p.lines); y++ {
		if remainingRows == 0 {
			break
		}

		remainingCols := a.Width
		grapheme := uniseg.NewGraphemes(p.lines[y])
		x := 0
		for grapheme.Next() {
			r := grapheme.Runes()
			if remainingCols == 0 {
				break
			}

			buf.SetContent(x+a.Col, idx+a.Row, r[0], tui.StyleDefault)
			w := grapheme.Width()
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

func renderScrollbar(buf *tui.OutputBuffer, area tui.Rect, scroll int, contentLength int) {
	thumbRune := '▌'
	col := area.Col + area.Width - 1

	visibleRatio := float64(area.Height) / float64(contentLength)
	thumbHeight := max(int(math.Floor(visibleRatio*float64(area.Height))), 1)
	if thumbHeight > area.Height {
		thumbHeight = area.Height
	}

	maxScrollOffset := max(contentLength-area.Height, 0)

	var scrollProgress float64
	if maxScrollOffset > 0 {
		scrollProgress = min(float64(scroll)/float64(maxScrollOffset), 1)
	} else {
		scrollProgress = 0
	}

	thumbY := min(
		area.Height-thumbHeight,
		max(
			0,
			int(math.Floor(scrollProgress*float64(area.Height-thumbHeight))),
		),
	)

	row := area.Row + thumbY
	for range thumbHeight {
		buf.SetContent(col, row, thumbRune, tui.StyleDefault)
		row += 1
	}
}
