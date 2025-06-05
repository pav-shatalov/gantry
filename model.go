package main

import (
	"fmt"
	"gantry/docker"
	"time"
)

type ApplicationModel struct {
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
	shouldRedraw          bool
	counter               int
	scrollOffset          int
}

func NewModel() (ApplicationModel, error) {
	state := ApplicationModel{shouldRedraw: true}
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

func (s *ApplicationModel) Update(msg Msg) Cmd {
	var cmd Cmd
	switch msg.(type) {
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
		s.shouldRedraw = true
		s.scrollOffset = 0
		logs, err := s.client.ContainerLogs(s.containers[s.selectedContainerIdx].Id)
		if err != nil {
			s.debug = fmt.Sprint(err)
		}
		s.selectedContainerLogs = logs
		s.debug = fmt.Sprintf("Loaded container logs. %d", s.counter)
		s.counter++
	case SelectNextContainerMsg:
		s.shouldRedraw = true
		if len(s.containers)-1 > s.selectedContainerIdx {
			s.selectedContainerIdx++
			cmd = NewCmd(LoadContainerLogsMsg{})
		}
	case SelectPrevContainerMsg:
		s.shouldRedraw = true
		if s.selectedContainerIdx > 0 {
			s.selectedContainerIdx--
			cmd = NewCmd(LoadContainerLogsMsg{})
		}
	case EnterContainerMsg:
		s.isRunning = false
		s.next = s.containers[s.selectedContainerIdx].Id
	case ResizeMsg:
		s.shouldRedraw = true
	case ScrollDownMsg:
		s.scrollOffset += 5
		s.shouldRedraw = true
	case ScrollUpMsg:
		s.scrollOffset -= 5
		if s.scrollOffset < 0 {
			s.scrollOffset = 0
		}
		s.shouldRedraw = true
	}

	return cmd
}

func (s *ApplicationModel) ContainerNames() []string {
	var names []string
	for _, ctr := range s.containers {
		names = append(names, ctr.Name)
	}

	return names
}

func (s *ApplicationModel) ContainerTableData() [][]string {
	var rows [][]string
	for _, ctr := range s.containers {
		row := []string{ctr.Name, ctr.Image, ctr.Id}
		rows = append(rows, row)
	}

	return rows
}
