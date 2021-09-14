package main

import (
	"fmt"
)

func main() {
	//var n [] int
	//printSlice(n)
	//n = append(n, 0)
	//printSlice(n)
	//n = append(n, 1,2,3)
	//printSlice(n)
	//n = append(n,2)
	//printSlice(n)
	//
	//n = append(n,3)
	//printSlice(n)
	//
	//n = append(n,4)
	//printSlice(n)
	//
	//n = append(n,5)
	//printSlice(n)
	//
	//n = append(n,6)
	//printSlice(n)
	//n := [5]int{1, 2, 3, 4, 5}
	//fmt.Println(n)
	//for _, i := range n {
	//	fmt.Println(i)
	//}

	//var goos = runtime.GOOS
	//fmt.Printf("The operating system is: %s\n", goos)
	//path := os.Getenv("PATH")
	//fmt.Printf("Path is %s\n", path)
	//N()
	//M()
	//N()
	//a := 0xFF
	b := 1.0234
	//fmt.Printf("result %x\n",a)
	//fmt.Printf("result %X\n",a)
	//fmt.Printf("result %d\n",a)
	//fmt.Printf("result %g\n",b)
	//fmt.Printf("result %f\n",b)
	//fmt.Printf("result %e\n",b)
	//fmt.Printf("result %02d\n",a)
	//fmt.Printf("result %1.2g\n",a)
	fmt.Printf("result %2.2f\n", b)

	var c1 = 5 + 10i
	fmt.Printf("result = %v\n", c1)
}

var t = "G"

func N() { print(t) }

func M() {
	t = "O"
	print(t)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
