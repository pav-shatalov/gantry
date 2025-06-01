package container

import (
	"gantry/geometry"
	"gantry/tui"
)

type Padding struct {
	top    int
	right  int
	bottom int
	left   int
}

type Container struct {
	area    geometry.Rect
	scroll  geometry.Position
	padding Padding
}

func New(area geometry.Rect, scroll geometry.Position) Container {
	return Container{area: area, scroll: scroll}
}

func (c Container) WithPadding(top int, right int, bottom int, left int) Container {
	c.padding = Padding{top: top, right: right, bottom: bottom, left: left}
	return c
}

func (c *Container) InnerArea() geometry.Rect {
	return geometry.Rect{
		X:      c.area.X + c.padding.left,
		Y:      c.area.Y + c.padding.top,
		Width:  c.area.Width - (c.padding.left + c.padding.right),
		Height: c.area.Height - (c.padding.top + c.padding.bottom),
	}
}

func (c *Container) Render(buf tui.ScreenBuffer, area geometry.Rect) {

}
