package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	// Hello World
	fmt.Println("Hello World")
	// Values
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	//Variables
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
	//Constant
	fmt.Println(s)

	const n = 500000000
	const g = 3e20 / n
	fmt.Println(g)
	fmt.Println(int64(g))
	fmt.Println(math.Sin(n))
}
