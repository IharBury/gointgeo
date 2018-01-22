package gointgeo

// Defines a 2D line by two of its points.
// Two definitions of the same line
// which do not have points matching in the same order
// are not considered equal.
// Use the Line method to get a structure with equality
// which represents the equality of the defined lines.
type Line2DDefinition struct {
	point1, point2 Point2D
}

func NewLine2DDefinition(point1, point2 Point2D) Line2DDefinition {
	if point1 == point2 {
		panic("Points are equal")
	}

	return Line2DDefinition{point1, point2}
}

func (lineDefinition Line2DDefinition) Point1() Point2D {
	return lineDefinition.point1
}

func (lineDefinition Line2DDefinition) Point2() Point2D {
	return lineDefinition.point2
}

func (lineDefinition Line2DDefinition) IsLineHorizontal() bool {
	return lineDefinition.point1.Y == lineDefinition.point2.Y
}

func (lineDefinition Line2DDefinition) IsLineVertical() bool {
	return lineDefinition.point1.X == lineDefinition.point2.X
}

func (lineDefinition Line2DDefinition) DoesLineHavePoint(point Point2D) bool {
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
	xDifference1 := int32(point.X) - int32(lineDefinition.point1.X)
	yDifference1 := int32(point.Y) - int32(lineDefinition.point1.Y)
	xDifference2 := int32(point.X) - int32(lineDefinition.point2.X)
	yDifference2 := int32(point.Y) - int32(lineDefinition.point2.Y)

	// Using int64 to avoid overflow when multiplying 17-bit values
	// resulting in 34-bit values.
	return int64(xDifference1)*int64(yDifference2) ==
		int64(yDifference1)*int64(xDifference2)
}

func (lineDefinition Line2DDefinition) lineSlope() fraction64 {
	if lineDefinition.IsLineVertical() {
		panic("The line is vertical.")
	}

	// Using int32 to avoid overflow when subtracting 16-bit values
	// resulting in 17-bit values.
	xDifference := int32(lineDefinition.point2.X) - int32(lineDefinition.point1.X)
	yDifference := int32(lineDefinition.point2.Y) - int32(lineDefinition.point1.Y)

	return newFraction64FromNonCanonical(int64(yDifference), int64(xDifference))
}

func (lineDefinition Line2DDefinition) xOfLineCrossingXAxis() fraction64 {
	if lineDefinition.IsLineHorizontal() {
		panic("The line is horizontal.")
	}

	// Using int32 to avoid overflow when subtracting 16-bit values
	// resulting in 17-bit values.
	xDifference := int32(lineDefinition.point2.X) - int32(lineDefinition.point1.X)
	yDifference := int32(lineDefinition.point2.Y) - int32(lineDefinition.point1.Y)

	return newFraction64FromNonCanonical(
		int64(lineDefinition.point1.X)*int64(yDifference)-
			int64(lineDefinition.point1.Y)*int64(xDifference),
		int64(yDifference))
}

func (line1Definition Line2DDefinition) LineEqualsTo(line2Definition Line2DDefinition) bool {
	return line1Definition.DoesLineHavePoint(line2Definition.point1) &&
		line1Definition.DoesLineHavePoint(line2Definition.point2)
}

func (lineDefinition Line2DDefinition) Line() Line2D {
	if lineDefinition.IsLineVertical() {
		// A vertical line is completely described by its x coordinate.
		return Line2D{
			isVertical:   true,
			isHorizontal: false,
			verticalX:    lineDefinition.point1.X}
	}

	if lineDefinition.IsLineHorizontal() {
		// A horizontal line is completely described by its y coordinate.
		return Line2D{
			isVertical:   false,
			isHorizontal: true,
			horizontalY:  lineDefinition.point1.Y}
	}

	// The slope together with the x axis crossing point
	// form a canonical representation of the line.
	return Line2D{
		isVertical:       false,
		isHorizontal:     false,
		slope:            lineDefinition.lineSlope(),
		xOfXAxisCrossing: lineDefinition.xOfLineCrossingXAxis()}
}
