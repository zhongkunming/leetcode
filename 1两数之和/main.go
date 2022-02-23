package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}

// 利用hashmap降低时间复杂度
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for index, value := range nums {
		if existIndex, flag := hashmap[target-value]; flag {
			return []int{existIndex, index}
		} else {
			hashmap[value] = index
		}
	}
	return nil
}
