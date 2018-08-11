// Source: https://blog.alexellis.io/golang-writing-unit-tests/
package main

import (
	"fmt"
)

// Sum : 썸ㅋ
func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(Sum(5, 5))
}
