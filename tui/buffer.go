package tui

import "github.com/gdamore/tcell/v2"

type ScreenBuffer struct {
	height int
	width  int

	rows    [][]BufferCell
	isDirty bool
}

type BufferCell struct {
	r     rune
	style tcell.Style
}

func (b *ScreenBuffer) SetContent(x int, y int, r rune, style tcell.Style) {
	b.rows[y][x] = BufferCell{r: r, style: style}
	b.isDirty = true
}

func (b *ScreenBuffer) GetContent() [][]BufferCell {
	b.isDirty = false
	return b.rows
}

func NewBuffer(w int, h int) ScreenBuffer {
	rows := make([][]BufferCell, 0)
	for y := range h {
		for x := range w {
			rows[y][x] = BufferCell{r: ' ', style: tcell.StyleDefault}
		}
	}

	return ScreenBuffer{height: h, width: w, rows: rows, isDirty: true}
}
