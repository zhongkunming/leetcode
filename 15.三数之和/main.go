package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%v", threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

/*
时间复杂度：O(N^2)
空间复杂度：O(logN)。
*/
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := make([][]int, 0)
	n := len(nums)
	for v1 := 0; v1 < n; v1++ {
		if v1 > 0 && nums[v1] == nums[v1-1] {
			continue
		}
		v3 := n - 1
		target := -1 * nums[v1]
		for v2 := v1 + 1; v2 < n; v2++ {
			if v2 > v1+1 && nums[v2] == nums[v2-1] {
				continue
			}
			for v2 < v3 && nums[v2]+nums[v3] > target {
				v3--
			}
			if v2 == v3 {
				break
			}
			if nums[v1]+nums[v2]+nums[v3] == 0 {
				res = append(res, []int{nums[v1], nums[v2], nums[v3]})
			}
		}

	}
	return res
}
