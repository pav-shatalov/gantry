package model

import (
	"gantry/geometry"
	"gantry/layout"
)

type LayoutModel struct {
	HeaderArea geometry.Rect
	MidArea    geometry.Rect
	BottomArea geometry.Rect
}

func NewLayoutModel(area geometry.Rect) LayoutModel {
	constraints := []layout.Constraint{
		layout.NewMin(2),
		layout.NewPercentage(100),
		layout.NewMin(1),
	}
	verticalAreas := layout.NewVertical(area).Constraints(constraints).Areas()

	return LayoutModel{
		HeaderArea: verticalAreas[0],
		MidArea:    verticalAreas[1],
		BottomArea: verticalAreas[2],
	}
}
