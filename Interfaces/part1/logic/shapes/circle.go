package shapes

import "math"

type circle struct{
	radius float64
}

func NewCircle(r float64) *circle {
	return &circle{r}
}

func (c circle) GetAnArea() float64 {
	return math.Sqrt(c.radius) * math.Pi
}

func (c circle) GetAPerimeter() float64 {
	return c.radius * 2 * math.Pi
}
