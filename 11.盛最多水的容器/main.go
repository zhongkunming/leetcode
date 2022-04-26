package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/*
双向指针

时间复杂度：O(n)
空间复杂度：O(1)

*/
func maxArea(height []int) int {
	i, j, max, maxTemp := 0, len(height)-1, 0, 0
	for i < j {
		if height[i] < height[j] {
			maxTemp = height[i] * (j - i)
			i++
		} else {
			maxTemp = height[j] * (j - i)
			j--
		}

		if max < maxTemp {
			max = maxTemp
		}
	}

	return max
}
