package main

import (
	"fmt"
)

/*
解1：把两个数组合并到一个数组中, 再利用长度求中位数
时间复杂度：O(m+n) 不符合题目要求
空间复杂度：O(m+n)
*/
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n1Length := len(nums1)
//	n2Length := len(nums2)
//	resultArrayLength := n1Length + n2Length
//	resultArray := make([]int, resultArrayLength)
//
//	i, j, k := 0, 0, 0
//	for k != resultArrayLength {
//		if i == n1Length {
//			for j != n2Length {
//				resultArray[k] = nums2[j]
//				k++
//				j++
//			}
//			break
//		}
//		if j == n2Length {
//			for i != n1Length {
//				resultArray[k] = nums1[i]
//				k++
//				i++
//			}
//			break
//		}
//		if nums1[i] < nums2[j] {
//			resultArray[k] = nums1[i]
//			k++
//			i++
//		} else {
//			resultArray[k] = nums2[j]
//			k++
//			j++
//		}
//	}
//	if resultArrayLength%2 == 1 {
//		return float64(resultArray[(resultArrayLength / 2)])
//	} else {
//		return float64(resultArray[resultArrayLength/2]+resultArray[(resultArrayLength/2)-1]) / 2
//	}
//}

/*
解2：两个数组的长度相加得len,
	如果len为奇数，则中位数是len/2
	如果len为偶数，则求中位数需要知道len/2， (len/2)-1
	遍历 len/2 次

时间复杂度：O(m+n) 不符合题目要求
空间复杂度：O(1)
*/
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n1Length := len(nums1)
//	n2Length := len(nums2)
//	rLength := n1Length + n2Length
//	n1Start, n2Start, right, left := 0, 0, -1, -1
//	for i := 0; i <= rLength/2; i++ {
//		left = right
//		if n1Start < n1Length && (n2Start >= n2Length || nums1[n1Start] < nums2[n2Start]) {
//			right = nums1[n1Start]
//			n1Start++
//		} else {
//			right = nums2[n2Start]
//			n2Start++
//		}
//	}
//	if rLength%2 == 1 {
//		return float64(right)
//	} else {
//		return float64(right+left) / 2
//	}
//}

/*
解3：利用二分法，可以看作一个求最k小的数字

时间复杂度：O(log(m+n))
空间复杂度：O(1)
*/
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	n1Length := len(nums1)
//	n2Length := len(nums2)
//	left := (n1Length + n2Length + 1) / 2
//	right := (n1Length + n2Length + 2) / 2
//	return float64(getKth(nums1, 0, n1Length-1, nums2, 0, n2Length-1, left)+
//		getKth(nums1, 0, n1Length-1, nums2, 0, n2Length-1, right)) * 0.5
//}
//func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
//	n1Length := end1 - start1 + 1
//	n2Length := end2 - start2 + 1
//	// 保证nums1最短
//	if n1Length > n2Length {
//		return getKth(nums2, start2, end2, nums1, start1, end1, k)
//	}
//	// nums1 消耗完
//	if n1Length == 0 {
//		return nums2[start2+k-1]
//	}
//	if k == 1 {
//		return min(nums1[start1], nums2[start2])
//	}
//	i := start1 + min(n1Length, k/2) - 1
//	j := start2 + min(n2Length, k/2) - 1
//
//	if nums1[i] > nums2[j] {
//		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
//	} else {
//		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
//	}
//}
//func min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}

/*
解4：把两个数组分成平分的两部分，分别用i, j表示
	此时临界值分别为：
	nums1[i-1]  | nums1[i]
	nums2[j-1]  | nums2[j]

	理论上 max(nums1[i-1], nums2[j-1]) < min(nums1[i], nums2[j])
	但存在 nums2[j-1] > nums1[i] 移动i或者j
		  nums1[i-1] > nums2[j] 移动i或者j
	但需要注意的是，移动 i(j) 的同时需保证 j(i) 也变化，保证两边数量平衡

时间复杂度：O(log(min(m,n))
空间复杂度：O(1)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1Length := len(nums1)
	n2Length := len(nums2)
	// 保证 n1Length <= n2Length
	if n1Length > n2Length {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax := 0, n1Length
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (n1Length+n2Length+1)/2 - i
		if j != 0 && i != n1Length && nums2[j-1] > nums1[i] {
			// i 需要增大
			iMin = i + 1
		} else if i != 0 && j != n2Length && nums1[i-1] > nums2[j] {
			// i 需要减小
			iMax = i - 1
		} else {
			// 达到要求，并且将边界条件列出来单独考虑
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (n1Length+n2Length)%2 == 1 {
				// 奇数不需要考虑右半部分
				return float64(maxLeft)
			}
			minRight := 0
			if i == n1Length {
				minRight = nums2[j]
			} else if j == n2Length {
				minRight = nums1[i]
			} else {
				minRight = min(nums2[j], nums1[i])
			}
			return float64(maxLeft+minRight) * 0.5
		}
	}
	return 0
}
func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	num1 := []int{1, 3, 4}
	num2 := []int{2, 6, 7, 8}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
