package layout

type ConstraintType int

const (
	Min ConstraintType = iota
	Max
	Percentage
	Length
)

type Constraint struct {
	constraintType ConstraintType
	value int
}

func NewMin(value int) Constraint {
	return Constraint{constraintType: Min, value: value}
}

func NewMax(value int) Constraint {
	return Constraint{constraintType: Max, value: value}
}

func NewPercentage(value int) Constraint {
	return Constraint{constraintType: Percentage, value: value}
}

func NewLength(value int) Constraint {
	return Constraint{constraintType: Length, value: value}
}

func (c *Constraint) Value() int {
	return c.value
}

func (c *Constraint) Type() ConstraintType {
	return c.constraintType
}
