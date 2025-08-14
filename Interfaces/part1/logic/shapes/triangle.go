package shapes

import (
	"errors"
	"math"
)

// тип треугольник
type triangle struct{
	side1 float64
	side2 float64
	side3 float64
}

func NewTriangle(s1, s2, s3 float64) (*triangle, error) {
	// getOwner
	if s1 + s2 > s3 && s1 + s3 > s2 && s2 + s3 > s1 {
		return &triangle{s1, s2, s3}, nil
	}
	return nil, errors.New("неверно введены стороны")

}

func (t triangle) GetAnArea() float64 {
	// площадь
	s1, s2, s3 := t.side1, t.side2, t.side3
	p := (s1 + s2 + s3) / 2.0
	return math.Sqrt(p * (p - s1) * (p - s2) * (p - s3))
}

func (t triangle) GetAPerimeter() float64 {
	// периметр
	return t.side1 + t.side2 + t.side3
}

