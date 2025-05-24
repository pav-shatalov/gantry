package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/geometry"
	"gantry/layout"
	"gantry/widget/block"
	"gantry/widget/paragraph"
	"gantry/widget/selectablelist"
)

type ContainerInfo struct {
	name string
	id   string
}

func (s *Store) view() {
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

	statsParagraph := paragraph.New(fmt.Sprintf("STATS: %+v", s.containerStats))
	statsParagraph.Render(screen, block1.InnerArea(innerAreas[0]))

	block2 := block.New()
	block2.Title("Block 2").BorderStyle(borderStyle).Render(screen, innerAreas[1])

	debugParagraph := paragraph.New(s.debug)
	debugParagraph.Render(screen, block2.InnerArea(innerAreas[1]))
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
	state := NewStore(&screen, surface)
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
				if ev.Key() == tcell.KeyEsc || ev.Rune() == 'q' {
					msg := Message{command: ExitCommand{}}
					state.update(msg)
				}

				if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
					msg := Message{command: SelectPrevContainer{}}

					state.update(msg)
				}
				if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
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
