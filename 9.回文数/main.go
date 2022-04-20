package main

import "strconv"

func main() {
	println(isPalindrome(121))
}

/*
把数字的后一半反转

时间复杂度：O(log(n))
空间复杂度：O(1)
*/
//func isPalindrome(x int) bool {
//	if x < 0 {
//		return false
//	}
//	if x < 10 {
//		return true
//	}
//	if x%10 == 0 {
//		return false
//	}
//	var x2 int
//	for x > x2 {
//		x2 = x2*10 + x%10
//		x = x / 10
//	}
//	return x == x2 || x == x2/10
//}

/*
先转成字符串

时间复杂度：O(n)
空间复杂度：O(n)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	itoa := strconv.Itoa(x)
	n := len(itoa)
	j := n - 1
	for i := 0; i < n/2; i++ {
		if itoa[i] != itoa[j] {
			return false
		}
		j--
	}
	return true
}
