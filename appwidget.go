package main

import (
	"fmt"
	"gantry/geometry"
	"gantry/tui"
	"gantry/widget/container"
	"gantry/widget/divider"
	"gantry/widget/list"
	"gantry/widget/paragraph"
	"strings"
)

type AppWidget struct {
	model *ApplicationModel
}

func (a AppWidget) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	topArea := paragraph.New("Select container")
	bottomArea := paragraph.New(
		fmt.Sprintf(
			"Client v%s, Server v%s, Debug: %s",
			a.model.dockerClientVersion,
			a.model.dockerServerVersion,
			a.model.debug,
		),
	)
	topArea.Render(buf, a.model.layoutModel.HeaderArea)
	bottomArea.Render(buf, a.model.layoutModel.BottomArea)

	containerLogsBlock := container.New(a.model.layoutModel.LogsArea, geometry.Position{X: 0, Y: 0}).WithPadding(0, 0, 0, 1)
	containerList := list.New(a.model.ContainerNames(), a.model.selectedContainerIdx)
	containerInfo := paragraph.New(strings.Join(a.model.selectedContainerLogs, "\n")).Scroll(a.model.scrollOffset)
	divider := divider.NewVertical()

	topArea.Render(buf, a.model.layoutModel.HeaderArea)
	bottomArea.Render(buf, a.model.layoutModel.BottomArea)
	containerList.Render(buf, a.model.layoutModel.ContainerListArea)
	divider.Render(buf, a.model.layoutModel.VerticalDividerArea)
	containerInfo.Render(buf, containerLogsBlock.InnerArea())
}
