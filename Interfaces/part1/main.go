package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"app/logic"
	"app/logic/shapes"
)


func main() {
	var NameOfShape string
	var method logic.GeometricShapes
	fmt.Print("Введите название фигуры: ")
	fmt.Scanln(&NameOfShape)
	fmt.Printf("Вы ввели: '%s'\n", NameOfShape)

	switch NameOfShape {
	case "треугольник":
		floatSlice, err := setSides(3)
		if err != nil {
			fmt.Println(err)
		}
		method, err = shapes.NewTriangle(floatSlice[0], floatSlice[1], floatSlice[2])
		if err != nil {
			log.Fatal(err)
		}
	case "прямоугольник", "квадрат":
		floatSlice, err := setSides(2)
		if err != nil {
			fmt.Println(err)
		}	
		method = shapes.NewSquare(floatSlice[0], floatSlice[1])	
	case "круг":
		floatSlice, err := setSides(1)
		if err != nil {
			fmt.Println(err)
		}
		method = shapes.NewCircle(floatSlice[0])
	default:
		log.Fatal("такой фигуры нет")
	}

	module := logic.NewGeometricModule(method)
	perimeter := module.GetAPerimeter()
	fmt.Println(perimeter)
	area := module.GetAnArea()
	fmt.Println(area)
}


func setSides(x int) ([]float64, error) {
	if x == 1 {
		fmt.Print("Введите значение радиуса круга: ")
	} else {
		fmt.Printf("Введите значения %dх сторон через пробел: ", x)
	}
	scanner := bufio.NewScanner(os.Stdin)

	// Считываем строку
	if !scanner.Scan() {
		return nil, fmt.Errorf("ошибка чтения ввода: %v", scanner.Err())
	}

	stringSlice := strings.Fields(scanner.Text())
	floatSlice := make([]float64, 0, len(stringSlice))
	for _, v := range stringSlice {
		floatValue, err  := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, fmt.Errorf("неверно введены данные, %s не число", v)
		}
		floatSlice = append(floatSlice, floatValue)
	}
	return floatSlice, nil
}
