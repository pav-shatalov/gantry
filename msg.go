package main

type Msg any

type ResizeMsg struct{}
type ExitMsg struct{}
type KeyPressMsg struct {
	KeyString string
}
type SelectNextContainerMsg struct{}
type SelectPrevContainerMsg struct{}
type LoadContainerListMsg struct{}
type EnterContainerMsg struct{}
