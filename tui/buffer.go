package tui

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

func NewBuffer(w int, h int) OutputBuffer {
	rows := createRows(w, h)
	return OutputBuffer{height: h, width: w, rows: rows, autogrow: false}
}

func createRows(w int, h int) [][]BufferCell {
	rows := make([][]BufferCell, h)
	for y := range h {
		rows[y] = createRow(w)
	}

	return rows
}

func createRow(w int) []BufferCell {
	row := make([]BufferCell, w)
	for x := range w {
		row[x] = BufferCell{R: ' ', Style: StyleDefault}
	}

	return row
}

func (b *OutputBuffer) Width() int {
	return b.width
}

func (b *OutputBuffer) Height() int {
	return len(b.rows)
}
