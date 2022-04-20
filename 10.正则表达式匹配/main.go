package main

func main() {
	println(isMatch("aaa", "ab*a*c*a"))
}

/*
动态规划

	f[i][j]表示s的前i个字符与p中的前j个字符是否能够匹配
	考虑p的第j个字符的匹配情况：
		1.如果p的第j个字符是一个小写字母，那么必须在s中匹配一个相同的小写字母，即
			f[i][j] = f[i - 1][j - 1], s[i] = p[j]
			f[i][j] = false, s[i] != p[j]
			也就是说，如果s的第i个字符与p的第j个字符不相同，那么无法进行匹配
		2.如果p的第j个字符是 *，那么就表示可以对p的第j−1个字符匹配任意自然数次。
			字母+星号的组合在匹配的过程中，本质上只会有两种情况：
				匹配s末尾的一个字符，将该字符扔掉，而该组合还可以继续进行匹配；
				不匹配字符，将该组合扔掉，不再进行匹配。
			f[i][j] = f[i-1][j] or f[i][j-2], s[i] = p[j-1]
			f[i][j] = f[i][j-2], s[i] != p[j-1]
		3.在任意情况下，只要p[j]是.，那么p[j]一定成功匹配s中的任意一个小写字母。
	最终的状态转移方程->:


时间复杂度: O(mn)
空间复杂度: O(mn)
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	states := make([][]bool, m+1)
	matchers := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}
	for i := 0; i < len(states); i++ {
		states[i] = make([]bool, n+1)
	}
	states[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				states[i][j] = states[i][j-2]
				if matchers(i, j-1) {
					states[i][j] = states[i][j] || states[i-1][j]
				}
			} else if matchers(i, j) {
				states[i][j] = states[i-1][j-1]
			}
		}
	}
	return states[m][n]
}
