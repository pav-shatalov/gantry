package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/tui"
)

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
	var renderCalls int
	appWidget := AppWidget{state: &state}

	for {
		defer func() {
			if r := recover(); r != nil {
				state.isRunning = false
				terminal.RestoreTerm()
				log.Fatal(r)
			}
		}()
		if state.isDirty {
			terminal.Draw(appWidget)
			state.isDirty = false
			renderCalls++
		}
		frames++

		handleEvent(&messageBus, &terminal)
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
	fmt.Printf("Render calls: %d\n", renderCalls)

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

func handleEvent(msgBus *MessageBus, terminal *tui.Terminal) {
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
		state.debug = fmt.Sprintf("New Msg %#v", msg)
		state.Update(msg, msgBus)
	default:
	}
}
