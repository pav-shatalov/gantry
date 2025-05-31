package widget

import (
	"gantry/geometry"
	"regexp"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Paragraph struct {
	text   string
	scroll geometry.Position
}

func (p *Paragraph) Render(screen tcell.Screen, area geometry.Rect) {
	if len(p.text) == 0 {
		return
	}

	lines := strings.Split(p.text, "\n")
	ansiRegexp := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

	for y := range area.Height {
		if y > len(lines)-1 {
			break
		}
		cleanLine := ansiRegexp.ReplaceAllString(lines[y], "")
		runes := []rune(cleanLine)
		for x := range area.Width {
			runeIdx := x
			if runeIdx >= len(runes) {
				continue
			}
			r := runes[runeIdx]
			screen.SetContent(area.X+x, area.Y+y, r, []rune{}, tcell.StyleDefault)
		}
	}

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

func NewParagraph(text string) Paragraph {
	return Paragraph{text: text, scroll: geometry.Position{X: 0, Y: 0}}
}
