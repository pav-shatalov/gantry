package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/geometry"
	"gantry/layout"
	"gantry/tui"
	"gantry/widget"
)

func view(s ApplicationState, screen tcell.Screen) {
	width, height := screen.Size()
	surface := geometry.Rect{X: 0, Y: 0, Width: width, Height: height}

	constraints := []layout.Constraint{
		layout.NewMin(2),
		layout.NewPercentage(100),
		layout.NewMin(2),
	}
	verticalAreas := layout.NewVertical(surface).Constraints(constraints).Areas()

	topArea := widget.NewParagraph("Top")
	midArea := widget.NewList(s.ContainerNames(), s.selectedContainerIdx)
	bottomArea := widget.NewParagraph(fmt.Sprintf("Last KeyPress: %s; Debug: %s", s.lastKey, s.debug))

	topArea.Render(screen, verticalAreas[0])
	midArea.Render(screen, verticalAreas[1])
	bottomArea.Render(screen, verticalAreas[2])
}

func main() {
	state, err := NewState()
	if err != nil {
		log.Fatal(err)
	}
	state = state.Update(LoadContainerListMsg{})
	terminal, err := tui.InitTerminal()
	if err != nil {
		log.Fatal(err)
	}

	var frames int

	for {
		terminal.Draw(func(screen tcell.Screen) {
			view(state, screen)
		})
		frames++

		msg := handleEvent(terminal)
		state = state.Update(msg)

		if !state.isRunning {
			break
		}

		time.Sleep(16 * time.Millisecond)
	}

	terminal.RestoreTerm()
	duration := time.Since(state.startTime)
	fps := int(float64(frames) / duration.Seconds())
	fmt.Printf("FPS: %d\n", fps)

	if state.next != "" {
		cmd := exec.Command("docker", "exec", "-ti", state.next, "bash")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to run command: %v\n", err)
			os.Exit(1)
		}
	}
}

func handleEvent(terminal tui.Terminal) Msg {
	var msg Msg

	select {
	case event := <-terminal.EventChannel:
		switch ev := event.(type) {
		case *tcell.EventKey:
			// Quit application
			if ev.Key() == tcell.KeyEsc || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				msg = ExitMsg{}
			} else if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
				msg = SelectPrevContainerMsg{}
			} else if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
				msg = SelectNextContainerMsg{}
			} else if ev.Key() == tcell.KeyCR {
				msg = EnterContainerMsg{}
			} else {
				// Pass keypress to application
				msg = KeyPressMsg{KeyString: ev.Name()}
			}
		}
	default:
	}

	return msg
}
