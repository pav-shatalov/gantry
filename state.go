package main

import "time"

type ApplicationState struct {
	debug     string
	isRunning bool
	startTime time.Time
}

func NewState() ApplicationState {
	return ApplicationState{
		isRunning: true,
		startTime: time.Now(),
	}
}

func (s ApplicationState) Update(msg Msg) ApplicationState {
	switch msg.(type) {
	case ExitMsg:
		s.isRunning = false
	}

	return s
}
