package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//k := 6
	//switch k {
	//case 4:
	//	fmt.Println("was <= 4")
	//	fallthrough
	//case 5:
	//	fmt.Println("was <= 5")
	//	fallthrough
	//case 6:
	//	fmt.Println("was <= 6")
	//	fallthrough
	//case 7:
	//	fmt.Println("was <= 7")
	//	fallthrough
	//case 8:
	//	fmt.Println("was <= 8")
	//	fallthrough
	//default:
	//	fmt.Println("default case")
	//}
	//	i := 0
	//sss:
	//	fmt.Println(i)
	//	if i < 15 {
	//		i++
	//		goto sss
	//
	//
	//	}

	//for i := 0; i < 100; i++ {
	//	switch {
	//	case i%3 == 0 && i%5 == 0:
	//
	//		fmt.Println("FizzBuzz")
	//	case i % 3 == 0:
	//		fmt.Println("Fizz")
	//	case i % 5 == 0:
	//		fmt.Println("Buzz")
	//	default:
	//		fmt.Println(i)
	//
	//
	//	}
	//}

	//s := "中文"
	//fmt.Println(len(s))

	//for i := 0; i < 5; i++ {
	//	var v int
	//	fmt.Printf("%d ", v)
	//	v = 5
	//}

	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}

	//i := 0
	//for { //since there are no checks, this is an infinite loop
	//	if i >= 3 { break }
	//	//break out of this for loop when this condition is met
	//	fmt.Println("Value of i is:", i)
	//	i++
	//}
	//fmt.Println("A statement just after for loop.")

	//for i := 0; i<7 ; i++ {
	//	if i%2 == 0 { continue }
	//	fmt.Println("Odd:", i)
	//}
	//fmt.Println(sqr(-1))
	//fmt.Println(sqr2(-1))

	//var reply *int
	//*reply = 0
	//
	//a := 10
	//b := 5
	//Multiply(a, b, reply)
	//fmt.Printf("Multiply:%d,a:%d,b:%d\n", *reply, a, b) // Multiply: 50
	//a := 1
	//
	//b := 1.0
	//c := "sss"
	//d := true
	//typecheck(a,b,c,d)
	//function1()
	//for i := 0; i < 10; i++ {
	//	result := sur(i)
	//	fmt.Println(result)
	//
	//}
	//f()

	//var g int
	//go func(i int) {
	//	s := 0
	//	for j := 0; j < i; j++ { s += j }
	//	g = s
	//	fmt.Println(g)
	//}(1000)
	//log.SetFlags()
	//start := time.Now()
	//time.Sleep(2 * 1e9)
	//
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	//var result uint64 = 0
	//start := time.Now()
	//for i := 0; i < LIM; i++ {
	//	result = fibonacci(i)
	//	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	//}
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Printf("longCalculation took this amount of time: %s\n", float64(delta)/float64(1e9))

	//var buffer bytes.Buffer
	//for {
	//	if s, ok := getNextString(); ok { //method getNextString() not shown here
	//		buffer.WriteString(s)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Print(buffer.String(), "\n")
	//var a  = [5] int{1,2,3,4, 5}
	//b := a[4:5]
	//fmt.Println(len(b))
	//s := "abcdef"
	//r := []byte(s)
	//for _, i := range r {
	//	fmt.Printf("result = %b\n", i)
	//
	//}
	//fmt.Println(r)

	//目标字符串
	//searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	//pat := "[0-9]+.[0-9]+" //正则
	//
	//f := func(s string) string {
	//	v, _ := strconv.ParseFloat(s, 32)
	//	return strconv.FormatFloat(v*2, 'f', 2, 32)
	//}
	//
	//if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
	//	fmt.Println("Match Found!")
	//}
	//
	//re, _ := regexp.Compile(pat)
	////将匹配到的部分替换为"##.#"
	//str := re.ReplaceAllString(searchIn, "##.#")
	//fmt.Println(str)
	////参数为函数时
	//str2 := re.ReplaceAllStringFunc(searchIn, f)
	//fmt.Println(str2)
	//fmt.Println(yy.A)

	// OK
	y := new(Bar)
	(y).thingOne = "hello"
	y.thingTwo = 1
	fmt.Println(y)

	// NOT OK
	//z := make(Bar{thingOne:"222"}) // 编译错误：cannot make type Bar
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	//// NOT OK
	//u := new(Foo)
	//(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	//(*u)["y"] = "world"
}

const LIM = 41

var fibs [LIM]uint64

//fibonacci
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
func sur(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = sur(n-1) + sur(n-2)
	}
	return
}
func function1() {
	fmt.Printf("In function1 at the top\n")
	function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!\n")
}
func opea(a int, b int) (int, int, int) {
	x := a + b
	y := a * b
	z := a - b
	return x, y, z
}

func opeaNo(a int, b int) (x int, y int, z int) {
	x = a + b
	y = a * b
	z = a - b
	return
}

func sqr(a float64) (float64, error) {
	if a < 0 {
		return math.NaN(), errors.New("I won't be able to do a sqrt of negative number!")

	}
	return math.Sqrt(a), nil

}

func sqr2(a float64) (x float64, err error) {
	if a < 0 {
		x = math.NaN()
		err = errors.New("I won't be able to do a sqrt of negative number!")
	} else {
		x = math.Sqrt(a)
		err = nil
	}

	return
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
	a = 11
	b = 6
}

func typecheck(values ...interface{}) {
	for _, value := range values {

		switch value.(type) {
		case int:
			fmt.Printf("int")
		case float32:
			fmt.Printf("float")
		case string:
			fmt.Printf("string")
		case bool:
			fmt.Printf("bool")
		default:
			fmt.Printf("int")
		}
	}
}
