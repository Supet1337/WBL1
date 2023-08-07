package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	var a, b int64 = 1 << 22, 1 << 22
	Biga := big.NewInt(a)
	Bigb := big.NewInt(b)
	res := big.NewInt(0)
	fmt.Println("Sum", res.Add(Biga, Bigb))
	fmt.Println("Substraction", res.Sub(Biga, Bigb))
	fmt.Println("Multiple", res.Mul(Biga, Bigb))
	fmt.Println("Division", res.Div(Biga, Bigb))

}
