package medium

func longestPalindrome(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		ret1 := palindrome(s, i, i)
		ret2 := palindrome(s, i, i+1)
		if len(ret) < len(ret1) {
			ret = ret1
		}
		if len(ret) < len(ret2) {
			ret = ret2
		}
	}
	return ret
}

func palindrome(s string, left int, right int) string {
	for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
