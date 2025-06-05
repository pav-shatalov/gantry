package geometry

type Rect struct {
	Col int
	Row int

	Width  int
	Height int
}

func NewRect(x int, y int, w int, h int) Rect {
	return Rect{Col: x, Row: y, Width: w, Height: h}
}

func (r Rect) Pos(x int, y int) Rect {
	r.Col = x
	r.Row = y

	return r
}

func (r Rect) Size(w int, h int) Rect {
	r.Width = w
	r.Height = h

	return r
}
