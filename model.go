package main

import (
	"fmt"
	"gantry/docker"
	"gantry/model"
	"gantry/tui"
	"time"
)

type ApplicationModel struct {
	layoutModel         model.LayoutModel
	logsModel           model.LogsViewModel
	containersModel     model.ContainersViewModel
	debug               string
	isRunning           bool
	startTime           time.Time
	lastKey             string
	client              *docker.Client
	next                string
	dockerClientVersion string
	dockerServerVersion string
	shouldRedraw        bool
	counter             int
}

func NewModel(area tui.Rect) (ApplicationModel, error) {
	layoutModel := model.NewLayoutModel(area)
	state := ApplicationModel{
		shouldRedraw:    true,
		layoutModel:     layoutModel,
		logsModel:       model.NewLogsViewModel([]string{}, layoutModel.LogsArea),
		containersModel: model.NewContainersViewModel([]docker.Container{}, layoutModel.ContainerListArea),
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
	switch m := msg.(type) {
	case ExitMsg:
		s.isRunning = false
	case LoadContainerListMsg:
		containers, err := s.client.LoadContainerList()
		if err != nil {
			s.debug = fmt.Sprintf("%s", err)
			break
		}

		s.containersModel.ReplaceContainers(containers)
		cmd = NewCmd(LoadContainerLogsMsg{})
	case LoadContainerLogsMsg:
		prevLogs := s.logsModel.Logs
		logs, err := s.client.ContainerLogs(s.containersModel.GetSelectedContainer().Id)
		if err != nil {
			s.debug = fmt.Sprint(err)
		}
		s.logsModel.Logs = logs
		s.logsModel.SetLines(s.logsModel.Logs)
		s.debug = fmt.Sprintf("Loaded container logs. %d", s.counter)
		s.counter++
		if len(prevLogs) > 0 && len(s.logsModel.Logs) > 0 {
			if prevLogs[len(prevLogs)-1] != s.logsModel.Logs[len(s.logsModel.Logs)-1] {
				s.shouldRedraw = true
			}
		} else {
			s.shouldRedraw = true
		}
	case SelectNextContainerMsg:
		s.shouldRedraw = true
		if s.containersModel.CanSelectNext() {
			s.containersModel.SelectNext()
			cmd = NewCmd(LoadContainerLogsMsg{})
			s.logsModel.AutoScroll = true
			s.logsModel.Scroll = 0
		}
	case SelectPrevContainerMsg:
		s.shouldRedraw = true
		if s.containersModel.CanSelectPrev() {
			s.containersModel.SelectPrev()
			cmd = NewCmd(LoadContainerLogsMsg{})
			s.logsModel.AutoScroll = true
			s.logsModel.Scroll = 0
		}
	case EnterContainerMsg:
		s.isRunning = false
		s.next = s.containersModel.GetSelectedContainer().Id
	case ResizeMsg:
		s.layoutModel.Resize(m.area)
		s.logsModel.Reflow(s.layoutModel.LogsArea)
		s.shouldRedraw = true
	case ScrollDownMsg:
		s.logsModel.Scroll += s.logsModel.Area.Height / 2
		s.logsModel.AutoScroll = false
		if s.logsModel.Scroll > len(s.logsModel.Lines)-1 {
			s.logsModel.Scroll = len(s.logsModel.Lines) - 1

		}
		s.shouldRedraw = true
	case ScrollUpMsg:
		s.logsModel.Scroll -= s.logsModel.Area.Height
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

// func (s *ApplicationModel) ContainerNames() []string {
// 	var names []string
// 	for _, ctr := range s.containers {
// 		names = append(names, ctr.Name)
// 	}
//
// 	return names
// }
