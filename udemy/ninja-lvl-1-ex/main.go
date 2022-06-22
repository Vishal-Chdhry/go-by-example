package ninjalvl1ex

import "fmt"

type T int

func Task1() {
	var a T = 42
	var b string = "James Bond"
	var c bool = true

	fmt.Println("VALUES: ", a, b, c)
	fmt.Printf("TYPES: %T %T %T\n", a, b, c)

	x := 52
	a = T(x)
	b = "Fifty Two"
	c = false
	fmt.Println("VALUES: ", a, b, c)
	fmt.Printf("TYPES: %T %T %T\n", a, b, c)
}
