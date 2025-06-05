package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	// "net/http"
	// _ "net/http/pprof"

	"github.com/gdamore/tcell/v2"

	"gantry/tui"
)

func main() {
	// go func() {
	// 	http.ListenAndServe(":6060", nil)
	// }()
	messageBus := NewMessageBus()
	messageBus.send(LoadContainerListMsg{})
	terminal, err := tui.InitTerminal()
	if err != nil {
		log.Fatal(err)
	}

	model, err := NewModel()
	if err != nil {
		log.Fatal(err)
	}
	var frames int
	var renderCalls int
	appWidget := AppWidget{model: &model}

	for {
		defer func() {
			if r := recover(); r != nil {
				model.isRunning = false
				terminal.RestoreTerm()
				log.Fatal(r)
			}
		}()

		if model.shouldRedraw {
			terminal.Draw(appWidget)
			model.shouldRedraw = false
			renderCalls++
		}
		frames++

		handleEvent(&messageBus, &terminal)
		handleMsg(&messageBus, &model)

		if !model.isRunning {
			break
		}

		time.Sleep(16 * time.Millisecond)
	}

	terminal.RestoreTerm()
	duration := time.Since(model.startTime)
	fps := int(float64(frames) / duration.Seconds())
	fmt.Printf("FPS: %d\n", fps)
	fmt.Printf("Render calls: %d\n", renderCalls)

	if model.next != "" {
		cmd := exec.Command("docker", "exec", "-ti", model.next, "bash")
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
			if ev.Key() == tcell.KeyEsc || ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				msgBus.send(ExitMsg{})
			} else if ev.Key() == tcell.KeyUp || ev.Rune() == 'k' {
				msgBus.send(SelectPrevContainerMsg{})
			} else if ev.Key() == tcell.KeyDown || ev.Rune() == 'j' {
				msgBus.send(SelectNextContainerMsg{})
			} else if ev.Key() == tcell.KeyCR {
				msgBus.send(EnterContainerMsg{})
			} else if ev.Key() == tcell.KeyCtrlD {
				msgBus.send(ScrollDownMsg{})
			} else if ev.Key() == tcell.KeyCtrlU {
				msgBus.send(ScrollUpMsg{})
			}
		case *tcell.EventResize:
			msgBus.send(ResizeMsg{})
		}
	default:
	}
}

func handleMsg(msgBus *MessageBus, state *ApplicationModel) {
	select {
	case msg := <-msgBus.ch:
		state.debug = fmt.Sprintf("New Msg %#v", msg)
		cmd := state.Update(msg)
		if cmd != nil {
			go cmd(msgBus)
		}
	default:
	}
}
