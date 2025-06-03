package main

import (
	"fmt"
	"gantry/docker"
	"time"

	"github.com/gdamore/tcell/v2"
)

type ApplicationState struct {
	debug                 string
	isRunning             bool
	startTime             time.Time
	lastKey               string
	client                *docker.Client
	containers            []docker.Container
	selectedContainerIdx  int
	selectedContainerLogs []string
	next                  string
	dockerClientVersion   string
	dockerServerVersion   string
	isDirty               bool
	counter               int
}

func NewState() (ApplicationState, error) {
	state := ApplicationState{isDirty: true}
	client, err := docker.NewClient()
	if err != nil {
		return state, err
	}

	state.isRunning = true
	state.startTime = time.Now()
	state.client = &client
	state.dockerClientVersion = client.Version()
	serverVersion, err := client.ServerVersion()
	if err == nil {
		state.dockerServerVersion = serverVersion
	}

	return state, nil
}

func (s *ApplicationState) Update(msg Msg, msgBus *MessageBus) {
	switch m := msg.(type) {
	case KeyPressMsg:
		if m.Key == tcell.KeyEsc || m.Key == tcell.KeyCtrlC || m.Rune == 'q' {
			msgBus.send(ExitMsg{})
		} else if m.Key == tcell.KeyUp || m.Rune == 'k' {
			msgBus.send(SelectPrevContainerMsg{})
		} else if m.Key == tcell.KeyDown || m.Rune == 'j' {
			msgBus.send(SelectNextContainerMsg{})
		} else if m.Key == tcell.KeyCR {
			msgBus.send(EnterContainerMsg{})
		}
	case ExitMsg:
		s.isRunning = false
	case LoadContainerListMsg:
		containers, err := s.client.LoadContainerList()
		if err != nil {
			s.debug = fmt.Sprintf("%s", err)
			break
		}

		s.containers = containers
		msgBus.send(LoadContainerLogsMsg{})
	case LoadContainerLogsMsg:
		s.isDirty = true
		logs, err := s.client.ContainerLogs(s.containers[s.selectedContainerIdx].Id)
		if err != nil {
			s.debug = fmt.Sprint(err)
		}
		s.selectedContainerLogs = logs
		s.debug = fmt.Sprintf("Loaded container logs. %d", s.counter)
		s.counter++
	case SelectNextContainerMsg:
		s.isDirty = true
		if len(s.containers)-1 > s.selectedContainerIdx {
			s.selectedContainerIdx++
			msgBus.send(LoadContainerLogsMsg{})
		}
	case SelectPrevContainerMsg:
		s.isDirty = true
		if s.selectedContainerIdx > 0 {
			s.selectedContainerIdx--
			msgBus.send(LoadContainerLogsMsg{})
		}
	case EnterContainerMsg:
		s.isRunning = false
		s.next = s.containers[s.selectedContainerIdx].Id
	}
}

func (s *ApplicationState) ContainerNames() []string {
	var names []string
	for _, ctr := range s.containers {
		names = append(names, ctr.Name)
	}

	return names
}

func (s *ApplicationState) ContainerTableData() [][]string {
	var rows [][]string
	for _, ctr := range s.containers {
		row := []string{ctr.Name, ctr.Image, ctr.Id}
		rows = append(rows, row)
	}

	return rows
}
