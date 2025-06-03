package tui

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type Terminal struct {
	Screen       tcell.Screen
	EventChannel chan tcell.Event
	quitChannel  chan struct{}
	colorMap     map[Color]tcell.Color
}

func colorMap() map[Color]tcell.Color {
	colorMap := map[Color]tcell.Color{
		ColorReset: tcell.ColorReset,
		ColorBlack: tcell.ColorBlack,
	}

	return colorMap
}

func InitTerminal() (Terminal, error) {
	app := Terminal{colorMap: colorMap()}
	screen, err := tcell.NewScreen()
	if err != nil {
		return app, err
	}

	if e := screen.Init(); e != nil {
		return app, err
	}

	screen.EnableMouse()
	screen.Clear()

	app.EventChannel = make(chan tcell.Event, 8)
	app.quitChannel = make(chan struct{})

	go screen.ChannelEvents(app.EventChannel, app.quitChannel)
	app.Screen = screen

	return app, nil
}

func (a *Terminal) Draw(widget Widget) {
	w, h := a.Screen.Size()
	buf := NewBuffer(w, h)
	area := geometry.Rect{X: 0, Y: 0, Width: w, Height: h}
	widget.Render(&buf, area)
	a.flushBuf(a.Screen, &buf)
}

func (a *Terminal) RestoreTerm() {
	close(a.quitChannel)
	a.Screen.Fini()
}

func (a *Terminal) flushBuf(s tcell.Screen, buf *OutputBuffer) {
	for y := range buf.Height() {
		for x := range buf.Width() {
			cell := buf.GetCell(x, y)
			style := a.convertStyle(&cell.Style)
			s.SetContent(x, y, cell.R, []rune{}, style)
		}
	}

	s.Show()
}

func (a *Terminal) convertStyle(s *Style) tcell.Style {
	tcellStyle := tcell.StyleDefault

	if s.fg != ColorReset {
		tcellStyle = tcellStyle.Foreground(convertColor(a.colorMap, s.fg))
	}

	return tcellStyle
}

func convertColor(m map[Color]tcell.Color, c Color) tcell.Color {
	v, exists := m[c]
	if !exists {
		v = tcell.ColorReset
	}

	return v
}
