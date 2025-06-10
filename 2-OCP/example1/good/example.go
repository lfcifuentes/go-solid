package good

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// New shapes can be added without modifying existing code
type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
