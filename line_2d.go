package gointgeo

// Represents a 2D line.
// Only allows to compare the represented lines.
// Use Line2DDefinition for all the other purposes.
type Line2D struct {
	isHorizontal     bool
	isVertical       bool
	horizontalY      int16
	verticalX        int16
	slope            fraction64
	xOfXAxisCrossing fraction64
}
