package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/geometry"
	"gantry/widget/paragraph"
	"gantry/widget/selectablelist"
)

type surface struct {
	width  int
	height int
}

type store struct {
	surface surface
	screen *tcell.Screen
	containers map[int]string
	selected_container_idx int
}

type cmd int

type Command interface{}

type ResizeCommand struct {}
type ExitCommand struct {}
type SelectNextContainer struct {}
type SelectPrevContainer struct {}

type Message struct {
	command Command
}

func initialState(screen *tcell.Screen, surface surface) store {
	return store{
		surface: surface,
		screen: screen,
		containers: map[int]string{0: "Container #1", 1: "Container #2"},
	}
}

func (s *store) update(msg Message) {
	screen := *s.screen;
	cmd := msg.command;
	switch cmd.(type) {
	case ResizeCommand: 
		newWidth, newHeight := screen.Size();
		s.surface.width = newWidth
		s.surface.height = newHeight
	case ExitCommand:
		screen.Fini()
		os.Exit(0)
	case SelectNextContainer:
		_,ok := s.containers[s.selected_container_idx + 1]
		if ok {
			s.selected_container_idx = s.selected_container_idx + 1
		}
	case SelectPrevContainer:
		_,ok := s.containers[s.selected_container_idx - 1]
		if ok {
			s.selected_container_idx = s.selected_container_idx - 1
		}
	}
}

func (s *store) view() {
	screen := *s.screen;
	message := fmt.Sprintf("Screen size: %dx%d\nAnother text", s.surface.width, s.surface.height)
	para := paragraph.New(message)
	para.Render(screen)

	list := selectablelist.New(s.containers, s.selected_container_idx);
	list.Render(screen, geometry.Position{X: 0, Y: 2})
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
				msg := Message{command: ResizeCommand{}}
				state.update(msg)
			case *tcell.EventKey:
				if (ev.Key() == tcell.KeyEsc) {
					msg := Message{command: ExitCommand{}}
					state.update(msg)
				}

				if (ev.Key() == tcell.KeyUp) {
					msg := Message{command: SelectPrevContainer{}}

					state.update(msg)
				}
				if (ev.Key() == tcell.KeyDown) {
					msg := Message{command: SelectNextContainer{}}

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
