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
	scrollOffset          int
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

func (s *ApplicationState) Update(msg Msg) Cmd {
	var cmd Cmd
	switch m := msg.(type) {

	// TODO: move it somewhere
	case KeyPressMsg:
		if m.Key == tcell.KeyEsc || m.Key == tcell.KeyCtrlC || m.Rune == 'q' {
			cmd = NewCmd(ExitMsg{})
		} else if m.Key == tcell.KeyUp || m.Rune == 'k' {
			cmd = NewCmd(SelectPrevContainerMsg{})
		} else if m.Key == tcell.KeyDown || m.Rune == 'j' {
			cmd = NewCmd(SelectNextContainerMsg{})
		} else if m.Key == tcell.KeyCR {
			cmd = NewCmd(EnterContainerMsg{})
		} else if m.Key == tcell.KeyCtrlD {
			cmd = NewCmd(ScrollDownMsg{})
		} else if m.Key == tcell.KeyCtrlU {
			cmd = NewCmd(ScrollUpMsg{})
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
		cmd = NewCmd(LoadContainerLogsMsg{})
	case LoadContainerLogsMsg:
		s.isDirty = true
		s.scrollOffset = 0
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
			cmd = NewCmd(LoadContainerLogsMsg{})
		}
	case SelectPrevContainerMsg:
		s.isDirty = true
		if s.selectedContainerIdx > 0 {
			s.selectedContainerIdx--
			cmd = NewCmd(LoadContainerLogsMsg{})
		}
	case EnterContainerMsg:
		s.isRunning = false
		s.next = s.containers[s.selectedContainerIdx].Id
	case ResizeMsg:
		s.isDirty = true
	case ScrollDownMsg:
		s.scrollOffset += 5
		s.isDirty = true
	case ScrollUpMsg:
		s.scrollOffset -= 5
		if s.scrollOffset < 0 {
			s.scrollOffset = 0
		}
		s.isDirty = true
	}

	return cmd
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
