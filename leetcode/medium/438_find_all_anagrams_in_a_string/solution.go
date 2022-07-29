package _438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	var ans []int
	need := make(map[uint8]int)
	for i, _ := range p {
		need[p[i]]++
	}
	windows := make(map[uint8]int)
	left := 0
	right := 0
	valid := 0
	for right < len(s) {
		curr := s[right]
		right++
		if need[curr] > 0 {
			windows[curr]++
			if windows[curr] == need[curr] {
				valid++
			}
		}
		if right-left == len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}
			leftVal := s[left]
			left++
			if windows[leftVal] > 0 {
				if windows[leftVal] == need[leftVal] {
					valid--
				}
				windows[leftVal]--
			}
		}
	}
	return ans
}
