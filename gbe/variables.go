package main

import "fmt"

func Variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 21, 12
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "walrus"
	fmt.Println(f)
}
