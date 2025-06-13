package main

import (
	"fmt"
	"gantry/docker"
	"gantry/model"
	"gantry/tui"
	"time"
)

type ApplicationModel struct {
	layoutModel           model.LayoutModel
	logsModel             model.LogsViewModel
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
}

func NewModel(area tui.Rect) (ApplicationModel, error) {
	layoutModel := model.NewLayoutModel(area)
	state := ApplicationModel{
		shouldRedraw: true,
		layoutModel:  layoutModel,
		logsModel:    model.NewLogsViewModel([]string{}, layoutModel.LogsArea),
	}
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
		prevLogs := s.selectedContainerLogs
		logs, err := s.client.ContainerLogs(s.containers[s.selectedContainerIdx].Id)
		if err != nil {
			s.debug = fmt.Sprint(err)
		}
		s.selectedContainerLogs = logs
		s.logsModel.SetLines(s.selectedContainerLogs)
		s.debug = fmt.Sprintf("Loaded container logs. %d", s.counter)
		s.counter++
		if len(prevLogs) > 0 && len(s.selectedContainerLogs) > 0 {
			if prevLogs[len(prevLogs)-1] != s.selectedContainerLogs[len(s.selectedContainerLogs)-1] {
				s.shouldRedraw = true
			}
		} else {
			s.shouldRedraw = true
		}
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
		s.logsModel.Reflow()
		s.shouldRedraw = true
	case ScrollDownMsg:
		s.logsModel.Scroll += 5
		s.logsModel.AutoScroll = false
		if s.logsModel.Scroll > len(s.logsModel.Lines)-1 {
			s.logsModel.Scroll = len(s.logsModel.Lines) + 10

		}
		s.shouldRedraw = true
	case ScrollUpMsg:
		s.logsModel.Scroll -= 5
		s.logsModel.AutoScroll = false
		if s.logsModel.Scroll < 0 {
			s.logsModel.Scroll = 0
		}
		s.shouldRedraw = true
	case TickMsg:
		cmd = NewCmd(LoadContainerLogsMsg{})
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
