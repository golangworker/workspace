package logic

import ("fmt")

type GeometricShapes interface {
	GetAnArea() float64
	GetAPerimeter() float64
}

type GeometricModule struct {
	geometricShapes GeometricShapes
}

func NewGeometricModule(geometricShapes GeometricShapes) *GeometricModule {
	// функция для создания модуля фигуры
	return &GeometricModule{
		geometricShapes: geometricShapes,
	}
}
func (p *GeometricModule) GetAnArea() string {
	// функция для вывода площади фигуры
	return fmt.Sprintf("Площадь: %.2f", p.geometricShapes.GetAnArea())
}

func (p *GeometricModule) GetAPerimeter() string {
	// функция для вывода периметра фигуры
	return fmt.Sprintf("Периметр: %.2f", p.geometricShapes.GetAPerimeter())
}
