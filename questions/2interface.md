# Что такое интерфейсы, как они применяются в Go?

### В Go интерфейсы представляют собой контракты, описывающие набор методов, которые должны быть реализованы определенными типами данных.
### Интерфейсы позволяют абстрагироваться от конкретных типов данных и работать с объектами, которые реализуют определенный набор методов, независимо от их конкретной реализации.
```
package main

import "fmt"

// Определение интерфейса
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Реализация интерфейса для структуры Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Реализация интерфейса для структуры Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Функции, принимающие интерфейс
func printPerimeter(s Shape) {
	fmt.Println(s.Perimeter())
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	// Создание объекта типа Circle
	c := Circle{Radius: 5}
	// Создание объекта типа Rectangle
	r := Rectangle{
		Width:  5,
		Height: 5,
	}
	// Вывод
	printPerimeter(&c)
	printArea(&c)
	printPerimeter(&r)
	printArea(&r)
}

```
