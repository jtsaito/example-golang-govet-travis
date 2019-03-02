package main

import "fmt"

// Sum is the sum of two integers
func Sum(x int, y int) int {
	return x + y
}

func main() {
	d := Sum(-1, 1)
	fmt.Printf("%d\n", d)
}
