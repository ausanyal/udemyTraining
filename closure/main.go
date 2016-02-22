package main

import "fmt"

func wrapper2() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func wrapper() func() int {
	y := wrapper2()
	return y
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
