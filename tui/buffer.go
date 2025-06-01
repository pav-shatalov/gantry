package tui

type ScreenBuffer struct {
	height int
	width  int

	rows [][]BufferCell
}

type BufferCell struct {
	r     rune
	style Style
}

func (b *ScreenBuffer) SetContent(x int, y int, r rune, style Style) {
	b.rows[y][x] = BufferCell{r: r, style: style}
}

func (b *ScreenBuffer) GetContent() [][]BufferCell {
	return b.rows
}

func (b *ScreenBuffer) GetCell(x int, y int) *BufferCell {
	return &b.rows[y][x]
}

func NewBuffer(w int, h int) ScreenBuffer {
	rows := make([][]BufferCell, h)
	for y := range h {
		rows[y] = make([]BufferCell, w)
		for x := range w {
			rows[y][x] = BufferCell{r: ' ', style: StyleDefault}
		}
	}

	return ScreenBuffer{height: h, width: w, rows: rows}
}

func (b *ScreenBuffer) Width() int {
	return b.width
}

func (b *ScreenBuffer) Height() int {
	return b.height
}
