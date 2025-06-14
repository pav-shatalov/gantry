package model

import (
	"fmt"
	"gantry/docker"
	"gantry/tui"
)

type ContainersViewModel struct {
	Area                 tui.Rect
	Containers           []docker.Container
	SelectedContainerIdx int
}

func NewContainersViewModel(containers []docker.Container, area tui.Rect) ContainersViewModel {
	return ContainersViewModel{
		Area:       area,
		Containers: containers,
	}
}

func (m *ContainersViewModel) Lines() []string {
	var lines []string
	for _, ctr := range m.Containers {
		line := fmt.Sprintf("%s[%s]", ctr.Name, ctr.State)
		lines = append(lines, line)
	}

	return lines
}

func (m *ContainersViewModel) ReplaceContainers(containers []docker.Container) {
	m.Containers = containers
}

func (m *ContainersViewModel) GetSelectedContainer() docker.Container {
	return m.Containers[m.SelectedContainerIdx]
}

func (m *ContainersViewModel) CountContainers() int {
	return len(m.Containers)
}

func (m *ContainersViewModel) Select(idx int) {
	m.SelectedContainerIdx = idx
}

func (m *ContainersViewModel) CanSelectNext() bool {
	nextIdx := m.SelectedContainerIdx
	return nextIdx < len(m.Containers)-1
}

func (m *ContainersViewModel) CanSelectPrev() bool {
	return m.SelectedContainerIdx > 0
}

func (m *ContainersViewModel) SelectNext() {
	m.SelectedContainerIdx++
}

func (m *ContainersViewModel) SelectPrev() {
	m.SelectedContainerIdx--
}
