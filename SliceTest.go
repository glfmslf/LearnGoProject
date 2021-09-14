package main

import "fmt"

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
	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println(n)
	for _, i := range n {
		fmt.Println(i)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
