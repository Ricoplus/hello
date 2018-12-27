package main

import "fmt"

import "time"

func main() {
	i := 2
	fmt.Print("wirte", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("is is weekend")
	default:
		fmt.Println("it is weekday")
	}
	//t:=time.Now()
	a := 12
	switch {
	case a < 15:
		fmt.Println(a, "a < 15")
	case a < 13:
		fmt.Println(a, "a < 13")
	default:
		fmt.Println("it's after noon")
	}
}
