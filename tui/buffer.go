package tui

import (
	"gantry/geometry"
)

type OutputBuffer struct {
	height   int
	width    int
	autogrow bool

	rows [][]BufferCell
}

type BufferCell struct {
	R     rune
	Style Style
}

func (b *OutputBuffer) SetContent(x int, y int, r rune, style Style) {
	if b.autogrow && y >= b.Height() {
		rowsToCreate := 64
		newRows := createRows(b.width, rowsToCreate)
		for range newRows {
			b.rows = append(b.rows, createRow(b.width))
		}
	}
	b.rows[y][x] = BufferCell{R: r, Style: style}
}

func (b *OutputBuffer) SetCell(x int, y int, cell *BufferCell) {
	b.rows[y][x] = *cell
}

func (b *OutputBuffer) GetContent() [][]BufferCell {
	return b.rows
}

func (b *OutputBuffer) GetCell(x int, y int) *BufferCell {
	return &b.rows[y][x]
}

func (b *OutputBuffer) FillFrom(tmpBuf *OutputBuffer, area geometry.Rect) {
	h := tmpBuf.Height()
	w := tmpBuf.Width()
	remainingRows := area.Height
	offset := 0
	if tmpBuf.Height() > 10 {
		offset = 20
	}
	rowsProcessed := 0
	for y := range h {
		remainingRows--

		remainingCols := area.Width
		for x := range w {
			remainingCols--

			cell := tmpBuf.GetCell(x, y+offset)
			b.SetCell(x+area.X, y+area.Y, cell)
			if remainingCols == 0 {
				break
			}
		}
		rowsProcessed++

		if remainingRows == 0 {
			break
		}
	}
}

func NewBuffer(w int, h int) OutputBuffer {
	rows := createRows(w, h)
	return OutputBuffer{height: h, width: w, rows: rows, autogrow: false}
}

func createRows(w int, h int) [][]BufferCell {
	rows := make([][]BufferCell, h, 256)
	for y := range h {
		rows[y] = createRow(w)
	}

	return rows
}

func createRow(w int) []BufferCell {
	row := make([]BufferCell, w, 256)
	for x := range w {
		row[x] = BufferCell{R: ' ', Style: StyleDefault}
	}

	return row
}

func NewAutogrowingBuffer(w int, initialHeight int) OutputBuffer {
	buf := NewBuffer(w, initialHeight)
	buf.autogrow = true
	return buf
}

func (b *OutputBuffer) Width() int {
	return b.width
}

func (b *OutputBuffer) Height() int {
	return len(b.rows)
}
