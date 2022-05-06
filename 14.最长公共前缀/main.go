package main

import "sort"

func main() {
	println(longestCommonPrefix([]string{""}))
}

/*

 */
func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	start := strs[0]
	end := strs[len(strs)-1]

	n1 := len(start)
	n2 := len(end)
	loc := 0
	for loc < n1 && loc < n2 {
		if start[loc] == end[loc] {
			loc += 1
			continue
		} else {
			break
		}
	}
	return start[0:loc]
}
