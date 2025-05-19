package span

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type Span struct {
	text    string
	padding int
}

func New(t string) Span {
	return Span{text: t, padding: 0}
}

func (s *Span) Render(screen tcell.Screen, pos geometry.Position) {
	row := pos.Y
	col := pos.X
	text := ""

	for range s.padding {
		text += " "
	}

	text += s.text

	for _, c := range text {
		screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
		col++
	}
}

func (s *Span) Length() int {
	return len(s.text)
}

func (s *Span) Padding(padding int) {
	s.padding = padding
}
