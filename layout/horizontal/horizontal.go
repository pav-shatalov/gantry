package horizontal

import (
	"gantry/cassowary"
	"gantry/geometry"
	"gantry/layout"
	"math"
)

type Horizontal struct {
	area geometry.Rect
	items []layout.Constraint
}

func New(area geometry.Rect) Horizontal {
	return Horizontal{ area: area }
}

// func (h *Horizontal) Items() []layout.Constraint {
// 	return h.items
// }

func (h *Horizontal) Add(c layout.Constraint) {
	h.items = append(h.items, c)
}

func (h *Horizontal) Areas() []geometry.Rect {
	var areas []geometry.Rect;
	x := h.area.X;
	y := h.area.Y;

	solver := cassowary.NewSolver()
	containerWidth := h.area.Width
	var totalWidthTerms []cassowary.Term
	var vars []cassowary.Symbol

	for _, item := range h.items {
		v := cassowary.New()
		switch item.Type() {
		case layout.Percentage:
			target := h.area.Width * item.Value() / 100
			// log.Printf("!!!!target w:%+v, val:%+v, target: %+v", h.area.Width, item.Value() / 100.0, target)
			solver.AddConstraintWithPriority(
				cassowary.Medium, 
				cassowary.NewConstraint(cassowary.EQ, float64(-1 * target), v.T(1)),
			)
			totalWidthTerms = append(totalWidthTerms, v.T(1))
		}
		vars = append(vars, v)
	}


	solver.AddConstraintWithPriority(
		cassowary.Strong,
		cassowary.NewConstraint(cassowary.EQ, float64(-1 * containerWidth), totalWidthTerms...),
	)

	for idx := range h.items {
		v := vars[idx]
		width := int(math.Round(solver.Val(v)))
		newArea := geometry.Rect{
			X: x,
			Y: y,
			Width: width,
			Height: h.area.Height,
		}

		areas = append(areas, newArea)
		x += width
	}

	return areas;
}
