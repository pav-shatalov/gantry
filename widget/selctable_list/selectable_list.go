package selectablelist

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type SelectableList struct {
	options []string
	selected int
}

func New(options []string) SelectableList {
	return SelectableList{options: options, selected: 0}
}

func (s *SelectableList) Render(screen tcell.Screen) {
	col := 0
	row := 2

	for itemIdx, item := range s.options {
		if itemIdx == s.selected {
			marker := fmt.Sprintf("> ")
			for _, c := range marker {
				screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
				col++
			}
		}
		for _, c := range item {
			screen.SetContent(col, row, c, []rune{}, tcell.StyleDefault)
			col++
		}
		col = 0
		row++
	}
}
