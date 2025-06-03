package main

import "github.com/gdamore/tcell/v2"

type Msg any

type ResizeMsg struct{}
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

type MessageBus struct {
	ch chan Msg
}

func (b *MessageBus) send(msg Msg) {
	select {
	case b.ch <- msg:
		// ok
	default:
		// dropped
	}
}

func NewMessageBus() MessageBus {
	msgChannel := make(chan Msg, 16)

	return MessageBus{ch: msgChannel}
}
