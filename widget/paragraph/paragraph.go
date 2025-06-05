package paragraph

import (
	"gantry/geometry"
	"gantry/tui"
	"regexp"

	"github.com/mattn/go-runewidth"
)

type Paragraph struct {
	text   string
	scroll int
}

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func Strip(str string) string {
	return re.ReplaceAllString(str, "")
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

		w := max(runewidth.RuneWidth(r), 1)
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
	withoutAnsi := Strip(p.text)
	if len(withoutAnsi) == 0 {
		return
	}

	t := reflow(withoutAnsi, area.Width)

	remainingRows := area.Height
	idx := 0

	for y := p.scroll; y < len(t); y++ {
		if remainingRows == 0 {
			break
		}

		remainingCols := area.Width
		for x, r := range t[y] {
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

func New(text string) Paragraph {
	return Paragraph{text: text, scroll: 0}
}

func (p Paragraph) Scroll(s int) Paragraph {
	p.scroll = s
	return p
}
