package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes11(nums)
	fmt.Println(nums)
}
func moveZeroes11(nums []int) {
	low := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[low], nums[i] = nums[i], nums[low]
			low++
		}
	}
}
