package main

import (
	"fmt"
	"math/cmplx"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func add(x int, y int) int{
	return x + y
}

// Возврат 2
func swap(x, y string)(string, string){
	return y, x
}

//Голый возврат
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Типы
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	Calculate(2)

	var a, b, c = 3, 4, "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)

	fmt.Println(add(42, 13))

	a1, b1 := swap("hello", "world")
	fmt.Println(a1, b1)

	fmt.Println(split(17))

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}


