package structs

import "testing"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Tests for perimeter of various shapes
func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, want: 40.0},
		{name: "Circle", shape: Circle{10}, want: 62.83185307179586},
		{name: "Triangle", shape: Triangle{10, 20}, want: 52.3606797749979},
	}

	for _, per := range perimeterTests {
		t.Run(per.name, func(t *testing.T) {
			got := per.shape.Perimeter()
			if got != per.want {
				t.Errorf("%#v Expected perimeter %g but got %g", per.shape, got, per.want)
			}
		})
	}

}

// Tests for area of various shapes
func TestArea(t *testing.T) {
	perimeterTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, want: 100.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10, 20}, want: 100.0},
	}

	for _, area := range perimeterTests {
		t.Run(area.name, func(t *testing.T) {

			got := area.shape.Area()
			if got != area.want {
				t.Errorf("%#v Expected area %g but got %g", area.shape, got, area.want)
			}

		})

	}

}
