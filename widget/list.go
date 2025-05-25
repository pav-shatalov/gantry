package widget

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
	runewidth "github.com/mattn/go-runewidth"
)

type List struct {
	options  []string
	selected any
}

func NewList(options []string, selectedIdx int) List {
	return List{options: options, selected: selectedIdx}
}

func (s *List) Render(screen tcell.Screen, area geometry.Rect) {
	col := area.X
	row := area.Y

	for itemIdx, item := range s.options {
		isSelected := itemIdx == s.selected
		if isSelected {
			marker := '›'
			markerSpan := NewSpan(string(marker))
			markerSpan.Render(screen, geometry.Position{X: col, Y: row})
			col += runewidth.RuneWidth(marker)
		}
		sp := NewSpan(item)
		if !isSelected {
			sp.Padding(1)
		}
		sp.Render(screen, geometry.Position{X: col, Y: row})

		col = area.X
		row++
	}
}
