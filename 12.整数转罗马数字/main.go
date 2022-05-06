package main

func main() {

	println(intToRoman(3))
}

/*

时间复杂度: O(1)
空间复杂度: O(1)
*/
func intToRoman(num int) string {
	m := []struct {
		values int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var sum = ""
	for _, v := range m {
		for num >= v.values {
			num -= v.values
			sum += v.symbol
		}
		if num == 0 {
			break
		}
	}

	return sum
}
