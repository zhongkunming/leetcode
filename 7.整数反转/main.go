package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1463847412
	fmt.Println(reverse(a))
}

/*

时间复杂度: O(log|x|)
空间复杂度: O(1)
*/

func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		y := x % 10
		x /= 10
		res = 10*res + y
	}
	return res
}
