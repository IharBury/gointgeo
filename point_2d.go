package gointgeo

import "fmt"

type Point2D struct {
	X, Y int16
}

func (point Point2D) String() string {
	return fmt.Sprintf("(%v, %v)", point.X, point.Y)
}
