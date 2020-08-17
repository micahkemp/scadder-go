package scad

import "math"

type regularPolygonAttributes struct {
	sides  int
	radius float64
}

func (p regularPolygonAttributes) angleRadiansPerPoint() float64 {
	return (math.Pi * 2) / float64(p.sides)
}

func (p regularPolygonAttributes) angleForPointNumber(pointNumber int) float64 {
	return p.angleRadiansPerPoint() * float64(pointNumber)
}

func (p regularPolygonAttributes) pointForPointNumber(pointNumber int) Point2D {
	x := math.Cos(p.angleForPointNumber(pointNumber))
	y := math.Sin(p.angleForPointNumber(pointNumber))

	return Point2D{x, y}
}

func (p regularPolygonAttributes) points() Points2D {
	points := make(Points2D, p.sides)

	for i := 0; i < p.sides; i++ {
		points[i] = p.pointForPointNumber(i)
	}

	return points
}

func NewRegularPolygon(name string, sides int, radius float64) scadTemplate {
	attrs := regularPolygonAttributes{sides, radius}

	return NewPolygon(name, attrs.points())
}

func NewHexagon(name string, radius float64) scadTemplate {
	return NewRegularPolygon(name, 6, radius)
}
