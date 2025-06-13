package tui

type Widget interface {
	Render(buf *OutputBuffer, area Rect)
}

type BlockWidget struct{}

type InlineWidget struct{}
