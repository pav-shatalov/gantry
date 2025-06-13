package main

import (
	"fmt"
	"gantry/tui"
	"gantry/widget/divider"
	"gantry/widget/list"
	"gantry/widget/paragraph"
)

type AppWidget struct {
	model *ApplicationModel
}

func (a AppWidget) Render(buf *tui.OutputBuffer, area tui.Rect) {
	headerInfo := paragraph.New([]string{"Select container"})
	debugInfo := paragraph.New([]string{
		fmt.Sprintf(
			"Client v%s, Server v%s, Debug: %s, Scroll: %d, Lines: %d",
			a.model.dockerClientVersion,
			a.model.dockerServerVersion,
			a.model.debug,
			a.model.logsModel.Scroll,
			len(a.model.logsModel.Lines),
		),
	})

	// containerLogsBlock := container.New(a.model.layoutModel.LogsArea, geometry.Position{X: 0, Y: 0}).WithPadding(0, 0, 0, 1)
	containerList := list.New(a.model.ContainerNames(), a.model.selectedContainerIdx)
	containerInfo := paragraph.New(a.model.logsModel.Lines).Scroll(a.model.logsModel.Scroll)

	divider := divider.NewVertical()

	headerInfo.Render(buf, a.model.layoutModel.HeaderArea)
	debugInfo.Render(buf, a.model.layoutModel.BottomArea)
	containerList.Render(buf, a.model.layoutModel.ContainerListArea)
	divider.Render(buf, a.model.layoutModel.VerticalDividerArea)
	containerInfo.Render(buf, a.model.layoutModel.LogsArea)
}
