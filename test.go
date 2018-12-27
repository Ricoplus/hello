package main

import "fmt"

func Judge(x int) int {
	if x == 0 {
		return 1
	} else {
		return x
	}
}
func main() {
	c := Judge(1)

	fmt.Println("c:", c)
}
