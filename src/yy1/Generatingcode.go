package main

import (
	"fmt"
)

var count int

func sum() {
	var n int
	fmt.Print("Enter a positive number: ")
	fmt.Scan(&n)

	//create the `for` loop that starts with `0` till `n`
	//and each iteration makes `sum += i`
	count += n
	fmt.Print("Sum : ", count, "\n")
}

func main() {
	i := 1
	var p *int
	p = &i
	newUsers := []User{{"John", false}, {"Mike", true}, {"Sarah", true}}
	for i, user := range newUsers {
		// run `sum()` for users that have no admin rights
		fmt.Println(i, user)
	}
	fmt.Println(p)
	fmt.Println(*p)
	i = 10
	fmt.Println(p)
	fmt.Println(*p)

}

type User struct {
	name  string
	admin bool
}
