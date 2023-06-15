package senthil

import "fmt"

func ZeroVal() {
	var a int
	var b float64
	var c string
	var d bool

	fmt.Printf("var a %T [%v]\n", a, a)
	fmt.Printf("var b %T [%v]\n", b, b)
	fmt.Printf("var c %T [%v]\n", c, c)
	fmt.Printf("var d %T [%v]\n", d, d)

	fmt.Printf("var aa %T [%v]\n", int32(a), a)
}
