package main

import (
	"fmt"
	"math"

	"github.com/lithdew/casso"
)

func main() {
	s := casso.NewSolver()

	child1 := casso.New()
	// child2 := casso.New()
	child3 := casso.New()
	
	// width := casso.New()

	totalWidth := 120

	c1 := casso.NewConstraint(casso.EQ, -1 * float64(totalWidth) * 0.5, child1.T(1));
	c3 := casso.NewConstraint(casso.EQ, -1 * float64(totalWidth) * 0.5, child3.T(1));

	c4 := casso.NewConstraint(casso.EQ, -1 * float64(totalWidth), child1.T(1), child3.T(1))

	s.AddConstraintWithPriority(casso.Weak, c1)
	// s.AddConstraintWithPriority(casso.Strong, c2)
	s.AddConstraintWithPriority(casso.Weak, c3)
	s.AddConstraintWithPriority(casso.Strong, c4)

	fmt.Printf("child 1 width %d\n", int(math.Round(s.Val(child1))));
	// fmt.Printf("child 2 width %d\n", int(math.Round(s.Val(child2))));
	fmt.Printf("child 3 width %d\n", int(math.Round(s.Val(child3))));
	// fmt.Printf("total width %f\n", s.Val(width));
}
