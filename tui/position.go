package tui

type Position struct {
	Col int
	Row int
}

func NewPosition(col int, row int) Position {
	return Position{Col: col, Row: row}
}
