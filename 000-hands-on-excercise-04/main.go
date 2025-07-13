//Multiple values

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}


func main() {

	// multiple values - swap with multiple
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//add values -- function continues
	fmt.Println(add(43, 13))
}

