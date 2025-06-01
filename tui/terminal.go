package tui

import (
	"gantry/geometry"

	"github.com/gdamore/tcell/v2"
)

type Terminal struct {
	Screen       tcell.Screen
	EventChannel chan tcell.Event
	quitChannel  chan struct{}
}

func InitTerminal() (Terminal, error) {
	app := Terminal{}
	screen, err := tcell.NewScreen()
	if err != nil {
		return app, err
	}

	if e := screen.Init(); e != nil {
		return app, err
	}

	screen.EnableMouse()
	screen.Clear()

	app.EventChannel = make(chan tcell.Event, 16)
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
	flushBuf(a.Screen, &buf)
}

func (a *Terminal) RestoreTerm() {
	close(a.quitChannel)
	a.Screen.Fini()
}

func flushBuf(s tcell.Screen, buf *ScreenBuffer) {
	for y := range buf.Height() {
		for x := range buf.Width() {
			s.SetContent(x, y, buf.GetCell(x, y).r, []rune{}, tcell.StyleDefault)
		}
	}

	s.Show()
}
