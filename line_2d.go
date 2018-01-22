package gointgeo

type Line2D struct {
	isHorizontal     bool
	isVertical       bool
	horizontalY      int16
	verticalX        int16
	slope            fraction64
	xOfXAxisCrossing fraction64
}
