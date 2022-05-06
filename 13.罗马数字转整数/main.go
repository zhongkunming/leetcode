package main

func main() {
	println(romanToInt("XIV"))
}

/*

时间复杂度: O(n)
空间复杂度: O(1)
*/
func romanToInt(s string) (ans int) {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

	n := len(s)
	for i := range s {
		value := m[s[i]]
		if i < n-1 && value < m[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}
