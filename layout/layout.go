package layout

import (
	"gantry/cassowary"
	"gantry/geometry"
	"math"
)

type Direction uint8

const (
	Horizontal Direction = iota
	Vertical
)

type Layout struct {
	area      geometry.Rect
	items     []Constraint
	direction Direction
}

func New(area geometry.Rect, d Direction) Layout {
	return Layout{
		area:      area,
		direction: d,
	}
}

func NewHorizontal(area geometry.Rect) Layout {
	return New(area, Horizontal)
}

func NewVertical(area geometry.Rect) Layout {
	return New(area, Vertical)
}

func (l Layout) Add(c Constraint) Layout {
	l.items = append(l.items, c)

	return l
}

func (l Layout) Constraints(constraints []Constraint) Layout {
	for _, c := range constraints {
		l = l.Add(c)
	}

	return l
}

func (l Layout) Areas() []geometry.Rect {
	var areas []geometry.Rect
	x := l.area.Col
	y := l.area.Row

	solver := cassowary.NewSolver()
	containerSize := l.area.Width
	if l.direction == Vertical {
		containerSize = l.area.Height
	}
	var totalSizeTerms []cassowary.Term
	var vars []cassowary.Symbol

	for _, item := range l.items {
		v := cassowary.New()
		switch item.Type() {
		case Percentage:
			target := containerSize * item.Value() / 100
			solver.AddConstraintWithPriority(
				cassowary.Medium,
				cassowary.NewConstraint(cassowary.EQ, float64(-1*target), v.T(1)),
			)
		case Min:
			target := item.Value()
			solver.AddConstraintWithPriority(
				cassowary.Strong,
				cassowary.NewConstraint(cassowary.GTE, float64(-1*target), v.T(1)),
			)
		case Length:
			target := item.Value()
			solver.AddConstraint(
				cassowary.NewConstraint(cassowary.EQ, float64(-1*target), v.T(1)),
			)
		}
		totalSizeTerms = append(totalSizeTerms, v.T(1))
		vars = append(vars, v)
	}

	solver.AddConstraintWithPriority(
		cassowary.Strong,
		cassowary.NewConstraint(cassowary.EQ, float64(-1*containerSize), totalSizeTerms...),
	)

	for idx := range l.items {
		v := vars[idx]
		val := int(math.Round(solver.Val(v)))
		width := val
		height := l.area.Height

		if l.direction == Vertical {
			width = l.area.Width
			height = val
		}

		newArea := geometry.Rect{
			Col:    x,
			Row:    y,
			Width:  width,
			Height: height,
		}

		areas = append(areas, newArea)
		if l.direction == Vertical {
			y += height
		} else {
			x += width
		}
	}

	return areas
}
