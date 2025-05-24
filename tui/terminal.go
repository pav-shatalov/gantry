package tui

import (
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

	app.EventChannel = make(chan tcell.Event)
	app.quitChannel = make(chan struct{})

	go screen.ChannelEvents(app.EventChannel, app.quitChannel)
	app.Screen = screen

	return app, nil
}

func (a *Terminal) Draw(t func(screen tcell.Screen)) {
	a.Screen.Clear() // todo: remove it?
	t(a.Screen)
	a.Screen.Show()
}

func (a *Terminal) RestoreTerm() {
	close(a.quitChannel)
	a.Screen.Fini()
}
