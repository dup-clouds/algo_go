package _567_permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	// 存储s1目标字符串map
	needMap := make(map[uint8]int)
	for i := range s1 {
		needMap[s1[i]]++
	}
	// 存储窗口字符串
	windowsMap := make(map[uint8]int)
	// 窗口左边界
	left := 0
	// 窗口右边界
	right := 0
	// 有效字符串个数
	valid := 0
	// 遍历s2
	for right < len(s2) {
		// 取出右边界元素
		curr := s2[right]
		right++
		// 判断右边界元素是否在目标字符串map中
		if needMap[curr] > 0 {
			// 窗口map增加
			windowsMap[curr]++
			// 判断是否有效(个数)
			if needMap[curr] == windowsMap[curr] {
				valid++
			}
		}
		// 当窗口字符串长度等于目标字符串长度时进行判断有效性
		for len(s1) == right-left {
			// 判断有效性
			if valid == len(needMap) {
				return true
			}
			// 无效时，说明left--right之间无s1的组合字符串
			// 移动left左边界
			leftVal := s2[left]
			left++
			// 判断是否删除有效字符串个数，窗口map是否进行更新操作
			if needMap[leftVal] > 0 {
				// 当left为有效字符个数组成时，valid-1
				if windowsMap[leftVal] == needMap[leftVal] {
					valid--
				}
				// 窗口map移除左边界值
				windowsMap[leftVal]--
			}
		}
	}
	return false
}

func checkInclusion2(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	// n > m 说明s2中不包含s1任意排列的子串
	if n > m {
		return false
	}
	// 定义两个数组 大小限定26, 数组下标即为元素的ASCII值-a的ASCII值，数组的值为该字符的个数值
	var cnt1, cnt2 [26]int
	// 遍历s1 将s1顺序放入数组中指定下标位置 将s2中0~len(s1)中元素放入数组中
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	// 判断两个数组是否相等 相等则条件成立
	if cnt1 == cnt2 {
		return true
	}
	// 到此则说明s2[0, len(s1)]中不包含s1排列字符串，则窗口需要往后移动一位，即s2最左边元素删除，s2最右边元素新增
	for i := n; i < m; i++ {
		// 增加窗口右边元素
		cnt2[s2[i]-'a']++
		// 去除窗口左边元素
		cnt2[s2[i-n]-'a']--
		// 判断是否相等
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}
