package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	fmt.Println(t)
	fmt.Println(t.Format(time.RFC850))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	var week time.Duration = 60 * 60 * 24 * 7 * 1e9
	fmt.Println(t.Add(0 - week).Format("2006-01-02 15:04:05"))
}
