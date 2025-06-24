package main

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

type Circle struct {
	Diameter float64
}

func (c *Circle) Area() float64 {
	return c.Diameter * 3.14
}

func (c *Circle) Perimeter() float64 {
	return 3.14 * c.Diameter * c.Diameter / 4
}
