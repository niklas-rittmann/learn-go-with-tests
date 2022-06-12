package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Calculate the perimeter of rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)

}

// Calculate the area of rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

// Calculate the perimeter of circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius

}

// Calculate the area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}

// Calculate the Perimeter of triangle
func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + math.Sqrt((t.Base*t.Base)+(t.Height*t.Height))

}

// Calculate the area of triangle
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2

}
