package main

import "fmt"

/*
利用左右指针取一个滑动窗口
利用hashmap判断是否有重复的字符
map['k'] != 0 即有
时间复杂度为 O(n)
空间复杂度为 O(Σ)
*/
//func lengthOfLongestSubstring(s string) int {
//	m := make(map[byte]int)
//	length := len(s)
//	// 利用左右指针，取一个滑动窗口
//	// pl 左指针, pr 右指针
//	pl, pr, max := 0, 0, 0
//	for pl < length {
//		if pl != 0 {
//			// 删除前一个字符
//			delete(m, s[pl-1])
//		}
//		// 移动右指针
//		for pr < length && m[s[pr]] == 0 {
//			m[s[pr]]++
//			pr++
//		}
//		if max < pr-pl {
//			max = pr - pl
//		}
//		pl++
//	}
//	return max
//}

/*
用map来存字符和字符所在的下标
只需要移动右指针，当发现重复的时候，开始移动左指针
需要注意的是，不能直接把左右指针距离与max直接划等号
例如abcbb，当右指针移动到第四个字符b的时候，左指针移动到第三个字符c，
	左右指针的距离变成2，因此当出现重复时，即时保存max的值

时间复杂度为 O(n)
空间复杂度为 O(Σ)
*/
func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	pl, max, tempMax := 0, 0, 0
	for pr, v := range s {
		idx, ok := m[v]
		if ok {
			if tempMax > max {
				max = tempMax
			}
			for pl <= idx {
				pl++
				tempMax--
			}
			tempMax += 1
			m[v] = pr
		} else {
			m[v] = pr
			tempMax++
		}
		tempMax = pr - pl + 1
	}
	if tempMax > max {
		max = tempMax
	}
	return max
}

func main() {
	s := "abcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
