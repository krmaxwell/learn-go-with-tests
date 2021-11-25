package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

type Point struct {
	X float64
	Y float64
}

// SecondHand represents the unit vector of the second hand of an analog clock at time t
// represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale to length of hand
	p = Point{p.X, -p.Y}                                      // flip over the axis
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate past center
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
