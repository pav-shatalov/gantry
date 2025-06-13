package model

import (
	"gantry/layout"
	"gantry/tui"
)

type LayoutModel struct {
	HeaderArea          tui.Rect
	BottomArea          tui.Rect
	ContainerListArea   tui.Rect
	VerticalDividerArea tui.Rect
	LogsArea            tui.Rect
}

func NewLayoutModel(area tui.Rect) LayoutModel {
	constraints := []layout.Constraint{
		layout.NewMin(2),
		layout.NewPercentage(100),
		layout.NewMin(1),
	}
	verticalAreas := layout.NewVertical(area).Constraints(constraints).Areas()

	midAreaSplit := layout.NewHorizontal(verticalAreas[1]).Constraints([]layout.Constraint{
		layout.NewPercentage(30),
		layout.NewLength(1),
		layout.NewPercentage(70),
	}).Areas()

	return LayoutModel{
		HeaderArea:          verticalAreas[0],
		BottomArea:          verticalAreas[2],
		ContainerListArea:   midAreaSplit[0],
		VerticalDividerArea: midAreaSplit[1],
		LogsArea:            midAreaSplit[2],
	}
}
