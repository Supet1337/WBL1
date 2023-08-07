package main

import (
	"fmt"
	"math"
	"point/point"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
func main() {

	point1 := point.PointNew(5, 6)
	point2 := point.PointNew(1, 2)
	fmt.Println("Distance: ", distance(point1, point2))
}

func distance(point1 *point.Point, point2 *point.Point) float64 {
	px1, py1 := point1.GetX(), point1.GetY()
	px2, py2 := point2.GetX(), point2.GetY()
	return math.Sqrt(math.Abs(math.Pow(px1, 2) - math.Pow(px2, 2) + math.Pow(py1, 2) - math.Pow(py2, 2)))
}
