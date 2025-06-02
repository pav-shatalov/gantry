package main

import (
	"fmt"
	"gantry/geometry"
	"gantry/layout"
	"gantry/tui"
	"gantry/widget/container"
	"gantry/widget/divider"
	"gantry/widget/list"
	"gantry/widget/paragraph"
	"strings"
)

type AppWidget struct {
	state *ApplicationState
}

func (a AppWidget) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	constraints := []layout.Constraint{
		layout.NewMin(2),
		layout.NewPercentage(100),
		layout.NewMin(1),
	}
	verticalAreas := layout.NewVertical(area).Constraints(constraints).Areas()

	topArea := paragraph.New("Top")
	bottomArea := paragraph.New(
		fmt.Sprintf(
			"Client v%s, Server v%s, Last KeyPress: %s; Debug: %s",
			a.state.dockerClientVersion,
			a.state.dockerServerVersion,
			a.state.lastKey,
			a.state.debug,
		),
	)
	topArea.Render(buf, verticalAreas[0])
	bottomArea.Render(buf, verticalAreas[2])
	midAreaSplit := layout.NewHorizontal(verticalAreas[1]).Constraints([]layout.Constraint{
		layout.NewPercentage(30),
		layout.NewLength(1),
		layout.NewPercentage(70),
	}).Areas()

	containerListBlock := container.New(midAreaSplit[2], geometry.Position{X: 0, Y: 0}).WithPadding(0, 0, 0, 1)
	containerList := list.New(a.state.ContainerNames(), a.state.selectedContainerIdx)
	containerInfo := paragraph.New(strings.Join(a.state.selectedContainerLogs, "\n"))
	divider := divider.NewVertical()

	topArea.Render(buf, verticalAreas[0])
	bottomArea.Render(buf, verticalAreas[2])
	containerList.Render(buf, midAreaSplit[0])
	divider.Render(buf, midAreaSplit[1])
	containerInfo.Render(buf, containerListBlock.InnerArea())
}
