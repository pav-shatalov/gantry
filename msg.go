package main

import (
	"gantry/tui"

	"github.com/gdamore/tcell/v2"
)

type Msg any

type ResizeMsg struct {
	area tui.Rect
}
type ExitMsg struct{}
type KeyPressMsg struct {
	KeyString string
	Key       tcell.Key
	Rune      rune
}
type SelectNextContainerMsg struct{}
type SelectPrevContainerMsg struct{}
type LoadContainerListMsg struct{}
type LoadContainerLogsMsg struct{}
type EnterContainerMsg struct{}
type ScrollUpMsg struct{}
type ScrollDownMsg struct{}
type TickMsg struct{}

type MessageBus struct {
	ch chan Msg
}

func (b *MessageBus) send(msg Msg) {
	select {
	case b.ch <- msg:
		// ok
	default:
		// dropped?
	}
}

func NewMessageBus() MessageBus {
	msgChannel := make(chan Msg, 4)

	return MessageBus{ch: msgChannel}
}
