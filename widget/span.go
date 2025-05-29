package widget

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type Span struct {
	text    string
	padding int
	style   tcell.Style
}

func NewSpan(t string) Span {
	return Span{text: t, padding: 0}
}

func (s Span) Style(style tcell.Style) Span {
	s.style = style
	return s
}

func (s Span) Padding(padding int) Span {
	s.padding = padding
	return s
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
		screen.SetContent(col, row, c, []rune{}, s.style)
		col++
	}
}

func (s *Span) Length() int {
	return len(s.text)
}
