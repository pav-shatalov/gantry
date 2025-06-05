package widget

import (
	"gantry/geometry"
	"gantry/layout"

	"github.com/gdamore/tcell/v2"
)

type Cell = string

type Row []Cell

type Table struct {
	rows []Row
}

func NewTable(inputRows [][]string) Table {
	var rows []Row
	for _, r := range inputRows {
		var row []Cell
		for _, c := range r {
			cell := c
			row = append(row, cell)
		}
		rows = append(rows, row)
	}
	return Table{rows: rows}
}

func (t *Table) Render(screen tcell.Screen, area geometry.Rect) {
	col := area.Col
	row := area.Row

	var constraints []layout.Constraint
	for range len(t.rows[0]) {
		constraints = append(constraints, layout.NewLength(area.Width/len(t.rows[0])))
	}
	layout := layout.NewHorizontal(area).Constraints(constraints).Areas()

	for _, r := range t.rows {
		cellCol := col
		for i, c := range r {
			if i > 0 {
				cellCol += layout[i-1].Width
			}
			cellArea := geometry.Rect{
				Col:    cellCol,
				Row:    row,
				Width:  layout[i].Width,
				Height: layout[i].Height,
			}
			renderCell(c, screen, cellArea)
		}
		row++
	}
}

func renderCell(cell Cell, screen tcell.Screen, area geometry.Rect) {
	col := area.Col
	row := area.Row

	runes := []rune(cell)

	for i := range area.Width {
		var r rune
		if len(runes) > i {
			r = runes[i]
		}
		screen.SetContent(col, row, r, []rune{}, tcell.StyleDefault)
		col++
	}
}
