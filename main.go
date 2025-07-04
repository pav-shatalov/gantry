package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
	"time"

	"github.com/gdamore/tcell/v2"

	"gantry/tui"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	messageBus := NewMessageBus()
	messageBus.send(LoadContainerListMsg{})
	terminal, err := tui.InitTerminal()
	if err != nil {
		log.Fatal(err)
	}

	model, err := NewModel(terminal.Area)
	if err != nil {
		log.Fatal(err)
	}
	var frames int
	var renderCalls int
	appWidget := AppWidget{model: &model}
	ticker := time.NewTicker(1 * time.Second)

	for {
		defer func() {
			if r := recover(); r != nil {
				model.isRunning = false
				terminal.RestoreTerm()
				debug.PrintStack()
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
		handleTick(ticker, &messageBus)

		if !model.isRunning {
			break
		}

		time.Sleep(16 * time.Millisecond)
	}

	terminal.RestoreTerm()

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
			} else if ev.Key() == tcell.KeyCtrlD || ev.Key() == tcell.KeyPgDn {
				msgBus.send(ScrollDownMsg{})
			} else if ev.Key() == tcell.KeyCtrlU || ev.Key() == tcell.KeyPgUp {
				msgBus.send(ScrollUpMsg{})
			}
		case *tcell.EventResize:
			w, h := terminal.Screen.Size()
			terminal.Area = tui.NewRect(0, 0, w, h)
			msgBus.send(ResizeMsg{
				area: terminal.Area,
			})
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

func handleTick(ticker *time.Ticker, msgBus *MessageBus) {
	select {
	case <-ticker.C:
		msgBus.send(TickMsg{})
	default:
	}
}
