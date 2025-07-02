package span

import (
	"gantry/tui"
)

type Span struct {
	text    string
	padding int
	style   tui.Style
}

func New(t string) Span {
	return Span{text: t, padding: 0}
}

func (s Span) Style(style tui.Style) Span {
	s.style = style
	return s
}

func (s Span) Padding(padding int) Span {
	s.padding = padding
	return s
}

func (s *Span) Render(buf *tui.OutputBuffer, pos tui.Position) {
	row := pos.Row
	col := pos.Col
	text := ""

	for range s.padding {
		text += " "
	}

	text += s.text

	for _, c := range text {
		buf.SetContent(col, row, c, s.style)
		col++
	}
}

func (s *Span) SetStyle(style tui.Style) {
	s.style = style
}

func (s *Span) SetPadding(padding int) {
	s.padding = padding
}

func (s *Span) Length() int {
	return len(s.text)
}
