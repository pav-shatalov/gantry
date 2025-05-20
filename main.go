package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/geometry"
	"gantry/layout"
	"gantry/widget/block"
	// "gantry/widget/paragraph"
	// "gantry/widget/selectablelist"
)

type store struct {
	surface                geometry.Rect
	screen                 *tcell.Screen
	containers             map[int]string
	selected_container_idx int
}

type cmd int

type Command interface{}

type ResizeCommand struct{}
type ExitCommand struct{}
type SelectNextContainer struct{}
type SelectPrevContainer struct{}

type Message struct {
	command Command
}

func initialState(screen *tcell.Screen, surface geometry.Rect) store {
	return store{
		surface:    surface,
		screen:     screen,
		containers: map[int]string{0: "Container #1", 1: "Container #2"},
	}
}

func (s *store) update(msg Message) {
	screen := *s.screen
	cmd := msg.command
	switch cmd.(type) {
	case ResizeCommand:
		newWidth, newHeight := screen.Size()
		s.surface.Width = newWidth
		s.surface.Height = newHeight
	case ExitCommand:
		screen.Fini()
		os.Exit(0)
	case SelectNextContainer:
		_, ok := s.containers[s.selected_container_idx+1]
		if ok {
			s.selected_container_idx = s.selected_container_idx + 1
		}
	case SelectPrevContainer:
		_, ok := s.containers[s.selected_container_idx-1]
		if ok {
			s.selected_container_idx = s.selected_container_idx - 1
		}
	}
}

func (s *store) view() {
	l := layout.New(s.surface, layout.Vertical)
	l.Add(layout.NewPercentage(20))
	l.Add(layout.NewPercentage(80))
	areas := l.Areas()
	// leftArea, rightArea := areas[0], areas[1]
	//
	screen := *s.screen
	// list := selectablelist.New(s.containers, s.selected_container_idx);
	// list.Render(screen, geometry.Position{X: 0, Y: 2})
	//
	block1 := block.New()
	block1.Render(screen, areas[1])

	block2 := block.New()
	block2.Render(screen, areas[0])

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
	surface := geometry.Rect{X: 0, Y: 0, Width: int(width), Height: int(height)}
	fmt.Printf("Surf: %+v", surface)
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
		case event := <-evChannel:
			switch ev := event.(type) {
			case *tcell.EventResize:
				msg := Message{command: ResizeCommand{}}
				state.update(msg)
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEsc {
					msg := Message{command: ExitCommand{}}
					state.update(msg)
				}

				if ev.Key() == tcell.KeyUp {
					msg := Message{command: SelectPrevContainer{}}

					state.update(msg)
				}
				if ev.Key() == tcell.KeyDown {
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
