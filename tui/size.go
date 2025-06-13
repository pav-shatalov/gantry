package tui

type Size struct {
	Width  int
	Height int
}

func NewSize(w int, h int) Size {
	return Size{Width: w, Height: h}
}
