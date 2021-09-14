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
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//ret := max(100, 200)
	//fmt.Println(ret)

	//testCallBack(1, callBack)
	//testCallBack(2, func(x int) int {
	//	fmt.Printf("我是回调，x：%d\n", x)
	//	return x
	//})
	//a := 1690           // 表示1.69
	//b := 1700           // 表示1.70
	//c := a * b          // 结果应该是2873000表示 2.873
	//fmt.Println(c)      // 内部编码
	//fmt.Println(float64(c) / 1000000) // 显示
	//fmt.Println(float32(c) / 1000000) // 显示

	//a := 100
	//var ptr *int
	//var pptr **int
	//ptr = &a
	//pptr = &ptr
	//fmt.Printf("a的值%d,地址%d\n",a,&a)
	//fmt.Printf("ptr的值=%d，地址%d\n", *ptr,ptr)
	//fmt.Printf("pptr的值%d,地址%d\n",**pptr,*pptr)
	//p := Person{"yy", 1}
	//fmt.Println(&p)
	//if result,err:=json.Marshal(&p);err==nil{
	//	fmt.Println(string(result))
	//
	//}

	//var n [5]int
	//for i := 0; i < 5; i++ {
	//	n[i] = i
	//}
	//fmt.Println(n)
	//n := make([]int)
	var n []int
	var name = 1
	fmt.Println(name)
	for i := 0; i < 10; i++ {
		n[i] = i
	}
	fmt.Println(n)
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
