package main

import (
	"fmt"
)

func main() {
	test := "babad"
	fmt.Println(longestPalindrome(test))
}

/*
解1：暴力解法

时间复杂符：O(n^3)
空间复杂度：O(1)
*/
//func longestPalindrome(s string) string {
//	var res string
//	max := 0
//	length := len(s)
//
//	for i := 0; i < length; i++ {
//		for j := i + 1; j <= length; j++ {
//			test := s[i:j]
//			if isPalindrome(test) && len(test) > max {
//				res = test
//				max = len(test)
//			}
//		}
//	}
//	return res
//}
//func isPalindrome(s string) bool {
//	length := len(s)
//	for i := 0; i < length/2; i++ {
//		if s[i] != s[length-i-1] {
//			return false
//		}
//	}
//	return true
//}

/*
解2：动态规划，利用二维数组
	对于一个子串而言，如果它是回文串，并且长度大于2，那么去除掉两边的字母，它还是回文串
	用P(i,j)表示字符串s从i到j的子串是否为回文，则：
		如果子串s[i,j]是回文，P(i,j)=true
		其他情况，P(i,j)=false
	则状态转移方程为
		P(i,j) = P(i-1,j+1)^(s[i]==s[j])
	可得边界方程为
		P(i,i) = true
		P(i,i+1) = s[i]==s[i+1]
时间复杂度：O(n^2)
空间复杂度：O(n^2)
*/
//func longestPalindrome(s string) string {
//	length := len(s)
//	if length < 2 {
//		return s
//	}
//	maxLength := 1
//	index := 0
//
//	dpa := make([][]bool, length)
//	for i := 0; i < length; i++ {
//		dpa[i] = make([]bool, length)
//		dpa[i][i] = true
//	}
//
//	// 子串的长度
//	for L := 2; L <= length; L++ {
//		for i := 0; i < length; i++ {
//			// 右边界
//			j := i + L - 1
//			// 右越界
//			if j >= length {
//				break
//			}
//			if s[i] != s[j] {
//				dpa[i][j] = false
//			} else {
//				if j-i < 3 {
//					dpa[i][j] = true
//				} else {
//					dpa[i][j] = dpa[i+1][j-1]
//				}
//			}
//			// 只要 dpa[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
//			if dpa[i][j] && j-i+1 > maxLength {
//				maxLength = j - i + 1
//				index = i
//			}
//		}
//	}
//
//	return s[index : index+maxLength]
//}

/*
解3：中心扩展算法
	在动态规划中，P(i,j) ← P(i+1,j−1) ← P(i+2,j−2) ← ⋯ ← 某一边界情况
	如果从边界情况（子串为1、2）向外扩展
	本质就是，枚举所有的扩展中心并尝试扩展直到无法扩展

时间复杂度：O(n)
空间复杂度：O(1)
*/
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		// 边界情况1
		left1, right1 := spread(s, i, i)
		// 边界情况2
		left2, right2 := spread(s, i, i+1)
		if right1-left1 >= right-left {
			left, right = left1, right1
		}
		if right2-left2 >= right-left {
			left, right = left2, right2
		}
	}
	return s[left : right+1]
}
func spread(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
