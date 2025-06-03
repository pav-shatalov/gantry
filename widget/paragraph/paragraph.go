package paragraph

import (
	"gantry/geometry"
	"gantry/tui"

	"github.com/mattn/go-runewidth"
)

type Paragraph struct {
	text   string
	scroll int
}

func reflow(s string, maxLength int) []string {
	var (
		lines []string
		line  []rune
		width int
	)

	for _, r := range s {
		if r == '\n' {
			lines = append(lines, string(line))
			line = nil
			width = 0
			continue
		}

		w := runewidth.RuneWidth(r)
		if width+w > maxLength {
			lines = append(lines, string(line))
			line = []rune{r}
			width = w
		} else {
			line = append(line, r)
			width += w
		}
	}

	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines
}

func (p *Paragraph) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	if len(p.text) == 0 {
		return
	}

	t := reflow(p.text, area.Width-1)

	remainingRows := area.Height
	idx := 0

	for y := p.scroll; y < len(t); y++ {
		if remainingRows == 0 {
			break
		}
		for x, r := range t[y] {
			buf.SetContent(x+area.X, area.Y+idx, r, tui.StyleDefault)
		}
		remainingRows--
		idx++
	}
}

func New(text string) Paragraph {
	return Paragraph{text: text, scroll: 0}
}

func (p Paragraph) Scroll(s int) Paragraph {
	p.scroll = s
	return p
}
