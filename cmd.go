package main

type Cmd func(*MessageBus)

func NewCmd(msg Msg) Cmd {
	return func(bus *MessageBus) {
		bus.send(msg)
	}
}
