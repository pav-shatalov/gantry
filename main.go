package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gdamore/tcell/v2"

	"gantry/geometry"
	"gantry/layout"
	"gantry/widget/block"
	"gantry/widget/selectablelist"
	// "gantry/widget/selectablelist"
)

type store struct {
	surface              geometry.Rect
	screen               *tcell.Screen
	containers           []ContainerInfo
	selectedContainerIdx int
}

type cmd int

type Command interface{}

type ResizeCommand struct{}
type ExitCommand struct{}
type SelectNextContainer struct{}
type SelectPrevContainer struct{}
type LoadContainerList struct{}

type Message struct {
	command Command
}

type ContainerInfo struct {
	name string
	id   string
}

func initialState(screen *tcell.Screen, surface geometry.Rect) store {
	return store{
		surface:    surface,
		screen:     screen,
		containers: []ContainerInfo{},
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
		if len(s.containers) > s.selectedContainerIdx+1 {
			s.selectedContainerIdx++
		}
	case SelectPrevContainer:
		if s.selectedContainerIdx >= 1 {
			s.selectedContainerIdx--
		}
	case LoadContainerList:
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		defer cli.Close()

		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			panic(err)
		}

		for _, ctr := range containers {
			s.containers = append(s.containers, ContainerInfo{id: ctr.ID, name: strings.Join(ctr.Names, "|")})
		}
	}
}

func (s *store) view() {
	globalLayout := layout.NewHorizontal(s.surface)
	globalLayout.Add(layout.NewPercentage(20))
	globalLayout.Add(layout.NewPercentage(80))
	globalAreas := globalLayout.Areas()

	innerLayout := layout.NewVertical(globalAreas[1])
	innerLayout.Add(layout.NewPercentage(50))
	innerLayout.Add(layout.NewPercentage(50))
	innerAreas := innerLayout.Areas()

	screen := *s.screen
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack)

	aside := block.New()
	aside.BorderStyle(borderStyle).Render(screen, globalAreas[0])

	var containerNames []string

	for _, containerInfo := range s.containers {
		containerNames = append(containerNames, containerInfo.name)
	}

	list := selectablelist.New(containerNames, s.selectedContainerIdx)
	list.Render(screen, aside.InnerArea(globalAreas[0]))

	block1 := block.New()
	block1.Title("Block 1").BorderStyle(borderStyle).Render(screen, innerAreas[0])

	block2 := block.New()
	block2.Title("Block 2").BorderStyle(borderStyle).Render(screen, innerAreas[1])
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
	state.update(Message{command: LoadContainerList{}})

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
		state.view()
		screen.Show()
		time.Sleep(50 * time.Millisecond)
	}
}
