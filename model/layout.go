package model

import (
	"gantry/layout"
	"gantry/tui"
)

type LayoutModel struct {
	HeaderArea        tui.Rect
	BottomArea        tui.Rect
	ContainerListArea tui.Rect
	LogsArea          tui.Rect
}

func NewLayoutModel(area tui.Rect) LayoutModel {
	vAreas := verticalAreas(area)
	midAreas := midAreas(vAreas[1])

	return LayoutModel{
		HeaderArea:        vAreas[0],
		BottomArea:        vAreas[2],
		ContainerListArea: midAreas[0],
		LogsArea:          midAreas[1],
	}
}

func (m *LayoutModel) Resize(area tui.Rect) {
	vAreas := verticalAreas(area)
	midAreas := midAreas(vAreas[1])

	m.HeaderArea = vAreas[0]
	m.BottomArea = vAreas[2]
	m.ContainerListArea = midAreas[0]
	m.LogsArea = midAreas[1]
}

func verticalAreas(area tui.Rect) []tui.Rect {
	constraints := []layout.Constraint{
		layout.NewMin(2),
		layout.NewPercentage(100),
		layout.NewMin(1),
	}
	return layout.NewVertical(area).Constraints(constraints).Areas()
}

func midAreas(area tui.Rect) []tui.Rect {
	return layout.NewHorizontal(area).Constraints([]layout.Constraint{
		layout.NewPercentage(30),
		layout.NewPercentage(70),
	}).Areas()
}
