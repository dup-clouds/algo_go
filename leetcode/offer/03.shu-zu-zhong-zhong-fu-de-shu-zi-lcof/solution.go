package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

func findRepeatNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		if numMap[num] > 0 {
			return num
		} else {
			numMap[num]++
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
