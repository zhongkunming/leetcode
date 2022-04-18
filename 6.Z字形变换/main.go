package main

import (
	"fmt"
)

func main() {
	str := "PAYPALISHIRING"
	rows := 4
	fmt.Println(convert(str, rows))

}

/*
解1：利用二维矩阵模拟
	n = 字符串s长度， r = numRows
	Z字形的变换周期为: t = r + r - 2 = 2r - 2
	每个周期会占用矩阵的: 1 + r - 2 = r - 1 列
	假设有 n/t 个周期，则列为 c = n/t * r - 1

时间复杂度：O(r⋅n)
空间复杂度：O(r⋅n)。矩阵需要 O(r⋅n) 的空间
*/
//func convert(s string, numRows int) string {
//	n, r := len(s), numRows
//	if n == 1 || r >= n {
//		return s
//	}
//	t := 2*r - 2
//	c := (n + t - 1) / t * (r - 1)
//	mat := make([][]byte, r)
//	for i := range mat {
//		mat[i] = make([]byte, c)
//	}
//	x, y := 0, 0
//	for i, ch := range s {
//		mat[x][y] = byte(ch)
//		if i%t < r-1 {
//			x++ // 向下移动
//		} else {
//			x--
//			y++ // 向右上移动
//		}
//	}
//	ans := make([]byte, 0, n)
//	for _, row := range mat {
//		for _, ch := range row {
//			if ch > 0 {
//				ans = append(ans, ch)
//			}
//		}
//	}
//	return string(ans)
//}

/*
解2：压缩矩阵
	把每一行看成一个列表，这样每次添加都可以看作为添加到每一行列表的最后

时间复杂度：O(n)
空间负载度：O(n)
*/
//func convert(s string, numRows int) string {
//	r := numRows
//	if r == 1 || r >= len(s) {
//		return s
//	}
//	mat := make([][]byte, r)
//	t, x := 2*r-2, 0
//	for i, ch := range s {
//		mat[x] = append(mat[x], byte(ch))
//		if i%t < r-1 {
//			x++
//		} else {
//			x--
//		}
//	}
//	return string(bytes.Join(mat, nil))
//}

/*
解3：直接构造

时间复杂度: O(n)
空间复杂度: O(1)
*/
func convert(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := 2*r - 2
	res := make([]byte, 0, n)
	for i := 0; i < r; i++ {
		for j := 0; j+i < n; j += t {
			res = append(res, s[j+i])

			if 0 < i && i < r-1 && j+t-i < n {
				res = append(res, s[j+t-i])
			}
		}
	}
	return string(res)
}
