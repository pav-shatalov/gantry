package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
		layout.NewMin(1),
	}
	verticalAreas := layout.NewVertical(surface).Constraints(constraints).Areas()

	topArea := widget.NewParagraph("Top")
	bottomArea := widget.NewParagraph(
		fmt.Sprintf(
			"Client v%s, Server v%s, Last KeyPress: %s; Debug: %s",
			s.dockerClientVersion,
			s.dockerServerVersion,
			s.lastKey,
			s.debug,
		),
	)
	midAreaSplit := layout.NewHorizontal(verticalAreas[1]).Constraints([]layout.Constraint{
		layout.NewPercentage(30),
		layout.NewLength(1),
		layout.NewPercentage(70),
	}).Areas()

	containerListBlock := widget.NewContainer(midAreaSplit[2], geometry.Position{X: 0, Y: 0}).WithPadding(0, 0, 0, 1)
	containerList := widget.NewList(s.ContainerNames(), s.selectedContainerIdx)
	containerInfo := widget.NewParagraph(strings.Join(s.selectedContainerLogs, "\n"))
	divider := widget.NewVerticalDivider()

	topArea.Render(screen, verticalAreas[0])
	bottomArea.Render(screen, verticalAreas[2])
	containerList.Render(screen, midAreaSplit[0])
	divider.Render(screen, midAreaSplit[1])
	containerInfo.Render(screen, containerListBlock.InnerArea())
}

func main() {
	state, err := NewState()
	messageBus := NewMessageBus()
	if err != nil {
		log.Fatal(err)
	}
	messageBus.send(LoadContainerListMsg{})
	terminal, err := tui.InitTerminal()
	if err != nil {
		log.Fatal(err)
	}

	var frames int

	for {
		terminal.Draw(func(screen tcell.Screen) {
			if state.isDirty {
				screen.Clear()
				view(state, screen)
				screen.Show()
				state.isDirty = false
			}
		})
		frames++

		handleEvent(&messageBus, terminal)
		handleMsg(&messageBus, &state)

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

func handleEvent(msgBus *MessageBus, terminal tui.Terminal) {
	select {
	case event := <-terminal.EventChannel:
		switch ev := event.(type) {
		case *tcell.EventKey:
			msgBus.send(KeyPressMsg{KeyString: ev.Name(), Key: ev.Key(), Rune: ev.Rune()})
		case *tcell.EventResize:
			msgBus.send(ResizeMsg{})
		}
	default:
	}
}

func handleMsg(msgBus *MessageBus, state *ApplicationState) {
	select {
	case msg := <-msgBus.ch:
		state.Update(msg, msgBus)
		state.isDirty = true
	default:
	}
}
