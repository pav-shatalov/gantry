package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/widget/paragraph"
	selectablelist "gantry/widget/selctable_list"
)

type surface struct {
	width  int
	height int
}

type store struct {
	surface surface
	screen *tcell.Screen
	containers []string
}

type cmd int

const (
	Resize cmd = iota
	Exit
)

type message struct {
	command cmd
}

func initialState(screen *tcell.Screen, surface surface) store {
	return store{surface: surface, screen: screen, containers: []string{"Container #1", "Container #2"}}
}

func (s *store) update(msg message) {
	screen := *s.screen;
	switch msg.command {
	case Resize: 
		newWidth, newHeight := screen.Size();
		s.surface.width = newWidth
		s.surface.height = newHeight
	case Exit:
		screen.Fini()
		os.Exit(0)
	}
}

func (s *store) view() {
	screen := *s.screen;
	message := fmt.Sprintf("Screen size: %dx%d\nAnother text", s.surface.width, s.surface.height)
	para := paragraph.New(message)
	para.Render(screen)

	list := selectablelist.New(s.containers);
	list.Render(screen)
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	width, height := screen.Size()
	surface := surface{width: int(width), height: int(height)}
	state := initialState(&screen, surface)

	evChannel := make(chan tcell.Event, 10)

	go func() {
		for {
			ev := screen.PollEvent()
			evChannel <- ev
		}
	}()

	// main loop
	for {
		select {
		case event := <- evChannel:
			switch ev := event.(type) {
			case *tcell.EventResize:
				msg := message{command: Resize}
				state.update(msg)
			case *tcell.EventKey:
				if (ev.Key() == tcell.KeyEsc) {
					msg := message{command: Exit}
					state.update(msg)
				}
			}
		default:
		}
		screen.Clear()
		state.view()
		screen.Show()
		time.Sleep(50 * time.Millisecond)
	}
}
