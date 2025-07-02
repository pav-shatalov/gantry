package main

import (
	"fmt"
	"gantry/tui"
	"gantry/tui/widget/list"
	"gantry/tui/widget/paragraph"
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

	containerList := list.New(a.model.containersModel.Lines(), a.model.containersModel.SelectedContainerIdx)
	containerList.Title("Containers")
	containerList.Borders(tui.AllBorders)
	containerList.BorderStyle(tui.StyleDefault.Fg(tui.ColorBlack))

	containerInfo := paragraph.New(a.model.logsModel.Lines)
	containerInfo.Title("Logs")
	containerInfo.Borders(tui.AllBorders)
	containerInfo.BorderStyle(tui.StyleDefault.Fg(tui.ColorBlack))
	containerInfo.Padding(0, 1, 0, 1)
	containerInfo.Scroll(a.model.logsModel.Scroll)

	headerInfo.Render(buf, a.model.layoutModel.HeaderArea)
	debugInfo.Render(buf, a.model.layoutModel.BottomArea)
	containerList.Render(buf, a.model.layoutModel.ContainerListArea)
	containerInfo.Render(buf, a.model.layoutModel.LogsArea)
}
