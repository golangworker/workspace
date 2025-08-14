package shapes

type square struct{
	side1 float64
	side2 float64
}
func NewSquare(s1, s2 float64) *square { // getOwner
	return &square{s1, s2}
}

func (s square) GetAnArea() float64 {
	return s.side1 * s.side2
}

func (s square) GetAPerimeter() float64 {
	return s.side1 * 2 + s.side2 * 2
}
