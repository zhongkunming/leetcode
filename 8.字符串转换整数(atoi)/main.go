package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("   +0 123"))
}

/*
 有限状态机 dfa 算法

 时间复杂度：O(n)
 空间复杂度：O(1)
*/
func myAtoi(s string) int {
	var (
		result   = 0
		sign     = 1
		flag     = true
		maxInt32 = math.MaxInt32
		minInt32 = math.MinInt32
	)
	for i := 0; i < len(s); i++ {
		if flag {
			if flag && s[i] == ' ' {
				continue
			}
			if s[i] == '-' {
				sign = -1
				flag = false
				continue
			}
			if s[i] == '+' {
				sign = 1
				flag = false
				continue
			}
		}

		if s[i] >= '0' && s[i] <= '9' {
			flag = false
			result = result*10 + int(s[i]-'0')

			if sign == -1 && result > maxInt32 {
				return minInt32
			}
			if sign == 1 && result > maxInt32 {
				return maxInt32
			}
		} else {
			break
		}
	}
	return result * sign
}
