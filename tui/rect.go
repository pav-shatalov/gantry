package tui

type Rect struct {
	Position
	Size
}

func NewRect(x int, y int, w int, h int) Rect {
	return Rect{Position: NewPosition(x, y), Size: NewSize(w, h)}
}

// func (r *Rect) Move(x int, y int) {
// 	r.Col = x
// 	r.Row = y
// }
//
// func (r *Rect) Resize(w int, h int) {
// 	r.Width = w
// 	r.Height = h
// }
