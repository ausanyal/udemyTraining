package main

import "fmt"

func main() {
	a := 10
	b := "golong"
	c := 2.3
	d := true
	e := "Hello world!"
	f := 'm'
	g := `whats ip ?`

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

	fmt.Println("")

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)

	fmt.Println("")

	var h int
	var i string
	var j float64
	var k bool

	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)
	fmt.Printf("%v \n", j)
	fmt.Printf("%v \n", k)

}
