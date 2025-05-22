package selectablelist

import (
	"fmt"

	"gantry/geometry"
	"gantry/widget/span"

	"github.com/gdamore/tcell/v2"
)

type SelectableList struct {
	options  []string
	selected any
}

func New(options []string, selectedIdx int) SelectableList {
	return SelectableList{options: options, selected: selectedIdx}
}

func (s *SelectableList) Render(screen tcell.Screen, area geometry.Rect) {
	col := area.X
	row := area.Y

	for itemIdx, item := range s.options {
		isSelected := itemIdx == s.selected
		if isSelected {
			marker := span.New(fmt.Sprintf("> "))
			marker.Render(screen, geometry.Position{X: col, Y: row})
			col += marker.Length()
		}
		sp := span.New(item)
		if !isSelected {
			sp.Padding(2)
		}
		sp.Render(screen, geometry.Position{X: col, Y: row})

		col = area.X
		row++
	}
}
