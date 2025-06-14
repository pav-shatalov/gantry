package span

import (
	"gantry/tui"
)

type Span struct {
	text    string
	padding int
}

func New(t string) Span {
	return Span{text: t, padding: 0}
}

// func (s Span) Style(style tcell.Style) Span {
// 	s.style = style
// 	return s
// }

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
		buf.SetContent(col, row, c, tui.StyleDefault)
		col++
	}
}

func (s *Span) Length() int {
	return len(s.text)
}
