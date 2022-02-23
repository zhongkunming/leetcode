package main

import (
	"fmt"
)

/*
解1：把两个数组合并到一个数组中
时间复杂度：O(m+n)
空间复杂度：O(m+n)
*/
/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1Length := len(nums1)
	n2Length := len(nums2)
	resultArrayLength := n1Length + n2Length
	resultArray := make([]int, resultArrayLength)

	i, j, k := 0, 0, 0
	for k != resultArrayLength {
		if i == n1Length {
			for j != n2Length {
				resultArray[k] = nums2[j]
				k++
				j++
			}
			break
		}
		if j == n2Length {
			for i != n1Length {
				resultArray[k] = nums1[i]
				k++
				i++
			}
			break
		}
		if nums1[i] < nums2[j] {
			resultArray[k] = nums1[i]
			k++
			i++
		} else {
			resultArray[k] = nums2[j]
			k++
			j++
		}
	}
	if resultArrayLength%2 == 1 {
		return float64(resultArray[(resultArrayLength / 2)])
	} else {
		return float64(resultArray[resultArrayLength/2]+resultArray[(resultArrayLength/2)-1]) / 2
	}
}*/

/*
解2：只需要记住中位数的下标
时间复杂度：O(m+n)
空间复杂度：O(1)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	return 0
}

func main() {
	num1 := []int{1, 3, 4}
	num2 := []int{2, 6, 7, 8}
	fmt.Println(findMedianSortedArrays(num1, num2))
}
