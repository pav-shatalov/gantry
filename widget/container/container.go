package container

import (
	"gantry/tui"
)

type Padding struct {
	top    int
	right  int
	bottom int
	left   int
}

// TODO: Remove it?
type Container struct {
	area    tui.Rect
	scroll  tui.Position
	padding Padding
}

func New(area tui.Rect, scroll tui.Position) Container {
	return Container{area: area, scroll: scroll}
}

func (c Container) WithPadding(top int, right int, bottom int, left int) Container {
	c.padding = Padding{top: top, right: right, bottom: bottom, left: left}
	return c
}

func (c *Container) InnerArea() tui.Rect {
	return tui.NewRect(
		c.area.Col+c.padding.left,
		c.area.Row+c.padding.top,
		c.area.Width-(c.padding.left+c.padding.right),
		c.area.Height-(c.padding.top+c.padding.bottom),
	)
}

func (c *Container) Render(buf tui.OutputBuffer, area tui.Rect) {

}
