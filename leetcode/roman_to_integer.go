package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	fmt.Print(romanToIntegerLast("MCMXCIV"))
}

// 罗马数字转整数
func romanToInteger(s string) int {
	m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
	sum := 0
	for i := 0; i < len(s); {

		if i+1 < len(s) {
			value, ok := m[s[i:i+2]]
			if ok {
				sum += value
				i += 2
			} else {
				sum += m[string(s[i:i+1])]
				i++
			}
		} else {
			sum += m[string(s[i:i+1])]
			i++
		}
	}
	return sum
}

var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

// 规则：正常是从大到小排列，如果出现后一个元素比前一个元素大，则减掉前面一个元素值即可
func romanToIntegerLast(s string) int {
	n := len(s)
	sum := 0
	for i := 0; i < n; i++ {
		if i+1 < n && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}
	return sum
}
