package main

import (
	"fmt"
	"gantry/docker"
	"time"
)

type ApplicationState struct {
	debug      string
	isRunning  bool
	startTime  time.Time
	lastKey    string
	client     docker.Client
	containers []docker.Container
}

func NewState() (ApplicationState, error) {
	state := ApplicationState{}
	client, err := docker.NewClient()
	if err != nil {
		return state, err
	}

	state.isRunning = true
	state.startTime = time.Now()
	state.client = client

	return state, nil
}

func (s ApplicationState) Update(msg Msg) ApplicationState {
	switch m := msg.(type) {
	case ExitMsg:
		s.isRunning = false
	case KeyPressMsg:
		s.lastKey = m.KeyString
	case LoadContainerListMsg:
		s.debug = fmt.Sprintf("Loading containers")
		containers, err := s.client.LoadContainerList()
		if err != nil {
			s.debug = fmt.Sprintf("%s", err)
			break
		}

		s.containers = containers
		s.debug = fmt.Sprint("Loaded containers")
	}

	return s
}
