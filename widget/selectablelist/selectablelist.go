package selectablelist

import (
	"fmt"
	"sort"

	"gantry/geometry"
	"gantry/widget/span"

	"github.com/gdamore/tcell/v2"
)

type SelectableList struct {
	options map[int]string
	selected int
}

func New(options map[int]string, selected_idx int) SelectableList {
	return SelectableList{options: options, selected: selected_idx}
}

func (s *SelectableList) Render(screen tcell.Screen, pos geometry.Position) {
	col := pos.X
	row := pos.Y

	// Let's get a slice of all keys
    keySlice := make([]int, 0)
    for key := range s.options {
        keySlice = append(keySlice , key)
    }

	// Now sort the slice
    sort.Ints(keySlice)

	for _, itemIdx := range keySlice {
		item := s.options[itemIdx]
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
		sp.Render(screen, geometry.Position{X: col, Y: row});
		
		col = 0
		row++
	}
}
