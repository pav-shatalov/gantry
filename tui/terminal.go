package tui

import (
	"github.com/gdamore/tcell/v2"
	_ "github.com/gdamore/tcell/v2/encoding"
)

type Terminal struct {
	Screen       tcell.Screen
	EventChannel chan tcell.Event
	quitChannel  chan struct{}
	colorMap     map[Color]tcell.Color
	Area         Rect
}

func colorMap() map[Color]tcell.Color {
	colorMap := map[Color]tcell.Color{
		ColorReset:  tcell.ColorReset,
		ColorBlack:  tcell.ColorBlack,
		ColorBlue:   tcell.ColorBlue,
		ColorRed:    tcell.ColorRed,
		ColorYellow: tcell.ColorYellow,
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

	w, h := screen.Size()
	app.Area = NewRect(0, 0, w, h)

	app.EventChannel = make(chan tcell.Event, 16)
	app.quitChannel = make(chan struct{})

	go screen.ChannelEvents(app.EventChannel, app.quitChannel)
	app.Screen = screen

	return app, nil
}

func (a *Terminal) Draw(widget Widget) {
	w := a.Area.Width
	h := a.Area.Height
	buf := NewBuffer(w, h)
	widget.Render(&buf, a.Area)
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
