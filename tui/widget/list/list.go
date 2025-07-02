package list

import (
	"gantry/tui"
	"gantry/tui/widget/span"
	runewidth "github.com/mattn/go-runewidth"
)

type List struct {
	tui.Block
	options  []string
	selected int
}

func New(options []string, selectedIdx int) List {
	return List{options: options, selected: selectedIdx, Block: tui.NewBlock()}
}

func (s *List) Render(buf *tui.OutputBuffer, area tui.Rect) {
	s.Block.Render(buf, area)
	a := s.Block.InnerArea(area)

	col := a.Col
	row := a.Row

	for itemIdx, item := range s.options {
		isSelected := itemIdx == s.selected
		sp := span.New(item).Padding(1)
		if isSelected {
			marker := '›'
			markerSpan := span.New(string(marker)).
				Style(tui.StyleDefault.Fg(tui.ColorYellow))
			markerSpan.Render(buf, tui.NewPosition(col, row))
			col += runewidth.RuneWidth(marker)
			sp.SetStyle(tui.StyleDefault.Fg(tui.ColorYellow))
		}
		if !isSelected {
			sp.SetPadding(2)
		}
		sp.Render(buf, tui.NewPosition(col, row))

		col = a.Col
		row++
	}
}
