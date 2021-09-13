package main

import (
	"fmt"
)

const (
	z = iota
	x
	v
)

var a, b, c = 11, false, "test11111"

const yy = "yy"

func main() {
	//fmt.Println(a,b,c,yy)
	//fmt.Println(unsafe.Sizeof(c))
	//fmt.Println(z,x,v)
	//fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}
