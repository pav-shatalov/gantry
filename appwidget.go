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
	model *ApplicationModel
}

func (a AppWidget) Render(buf *tui.OutputBuffer, area geometry.Rect) {
	if area.Width == 0 {
		return
	}
	// constraints := []layout.Constraint{
	// 	layout.NewMin(2),
	// 	layout.NewPercentage(100),
	// 	layout.NewMin(1),
	// }
	// verticalAreas := layout.NewVertical(area).Constraints(constraints).Areas()

	topArea := paragraph.New("Select container")
	bottomArea := paragraph.New(
		fmt.Sprintf(
			"Client v%s, Server v%s, Last KeyPress: %s; Debug: %s",
			a.model.dockerClientVersion,
			a.model.dockerServerVersion,
			a.model.lastKey,
			a.model.debug,
		),
	)
	topArea.Render(buf, a.model.layoutModel.HeaderArea)
	bottomArea.Render(buf, a.model.layoutModel.BottomArea)
	midAreaSplit := layout.NewHorizontal(a.model.layoutModel.MidArea).Constraints([]layout.Constraint{
		layout.NewPercentage(30),
		layout.NewLength(1),
		layout.NewPercentage(70),
	}).Areas()

	containerListBlock := container.New(midAreaSplit[2], geometry.Position{X: 0, Y: 0}).WithPadding(0, 0, 0, 1)
	containerList := list.New(a.model.ContainerNames(), a.model.selectedContainerIdx)
	containerInfo := paragraph.New(strings.Join(a.model.selectedContainerLogs, "\n")).Scroll(a.model.scrollOffset)
	divider := divider.NewVertical()

	topArea.Render(buf, a.model.layoutModel.HeaderArea)
	bottomArea.Render(buf, a.model.layoutModel.BottomArea)
	containerList.Render(buf, midAreaSplit[0])
	divider.Render(buf, midAreaSplit[1])
	containerInfo.Render(buf, containerListBlock.InnerArea())
}
