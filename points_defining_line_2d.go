package gointgeo

type PointsDefiningLine2D struct {
	point1, point2 Point2D
}

func NewPointsDefiningLine2D(point1, point2 Point2D) PointsDefiningLine2D {
	if point1 == point2 {
		panic("Points are equal")
	}

	return PointsDefiningLine2D{point1, point2}
}

func (linePoints PointsDefiningLine2D) Point1() Point2D {
	return linePoints.point1
}

func (linePoints PointsDefiningLine2D) Point2() Point2D {
	return linePoints.point2
}

func (linePoints PointsDefiningLine2D) IsLineHorizontal() bool {
	return linePoints.point1.Y == linePoints.point2.Y
}

func (linePoints PointsDefiningLine2D) IsLineVertical() bool {
	return linePoints.point1.X == linePoints.point2.X
}

func (linePoints PointsDefiningLine2D) DoesLineHavePoint(point Point2D) bool {
	// Ignoring the case when the point is equal
	// to either the point1 or the point2,
	// the point lies on the line defined by the point1 and the point2,
	// iff the line defined by the point and the point1 has the same slope
	// as the line defined by the point and the point2.
	// To avoid division (which is slower than multiplication)
	// and especially division by zero for vertical lines,
	// we rewrite the slope equality equation using multiplication.
	// That also takes care of the case when the point is equal
	// to either the point1 or the point2.
	// Both sides of the equation become zero in that case.

	// Using int32 to avoid overflow when subtracting 16-bit values
	// resulting in 17-bit values.
	xDifference1 := int32(point.X) - int32(linePoints.point1.X)
	yDifference1 := int32(point.Y) - int32(linePoints.point1.Y)
	xDifference2 := int32(point.X) - int32(linePoints.point2.X)
	yDifference2 := int32(point.Y) - int32(linePoints.point2.Y)

	// Using int64 to avoid overflow when multiplying 17-bit values
	// resulting in 34-bit values.
	return int64(xDifference1)*int64(yDifference2) ==
		int64(yDifference1)*int64(xDifference2)
}

func (linePoints PointsDefiningLine2D) lineSlope() fraction64 {
	if linePoints.IsLineVertical() {
		panic("The line is vertical.")
	}

	// Using int32 to avoid overflow when subtracting 16-bit values
	// resulting in 17-bit values.
	xDifference := int32(linePoints.point2.X) - int32(linePoints.point1.X)
	yDifference := int32(linePoints.point2.Y) - int32(linePoints.point1.Y)

	return newFraction64FromNonCanonical(int64(yDifference), int64(xDifference))
}

func (linePoints PointsDefiningLine2D) xOfLineCrossingXAxis() fraction64 {
	if linePoints.IsLineHorizontal() {
		panic("The line is horizontal.")
	}

	// Using int32 to avoid overflow when subtracting 16-bit values
	// resulting in 17-bit values.
	xDifference := int32(linePoints.point2.X) - int32(linePoints.point1.X)
	yDifference := int32(linePoints.point2.Y) - int32(linePoints.point1.Y)

	return newFraction64FromNonCanonical(
		int64(linePoints.point1.X)*int64(yDifference)-
			int64(linePoints.point1.Y)*int64(xDifference),
		int64(yDifference))
}

func (linePoints PointsDefiningLine2D) Line() Line2D {
	if linePoints.IsLineVertical() {
		// A vertical line is completely described by its x coordinate.
		return Line2D{
			isVertical:   true,
			isHorizontal: false,
			verticalX:    linePoints.point1.X}
	}

	if linePoints.IsLineHorizontal() {
		// A horizontal line is completely described by its y coordinate.
		return Line2D{
			isVertical:   false,
			isHorizontal: true,
			horizontalY:  linePoints.point1.Y}
	}

	// The slope together with the x axis crossing point
	// form a canonical representation of the line.
	return Line2D{
		isVertical:       false,
		isHorizontal:     false,
		slope:            linePoints.lineSlope(),
		xOfXAxisCrossing: linePoints.xOfLineCrossingXAxis()}
}
