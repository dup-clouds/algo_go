package _380_insert_delete_getrandom_o1

import "math/rand"

type RandomizedSet struct {
	nums        []int
	valIndexMap map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{nums: []int{}, valIndexMap: map[int]int{}}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.valIndexMap[val]; ok {
		return false
	}
	r.valIndexMap[val] = len(r.nums)
	r.nums = append(r.nums, val)
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	index, ok := r.valIndexMap[val]
	if !ok {
		return false
	}
	// 注意数组下标越界行为
	if index != len(r.nums)-1 {
		r.nums[index] = r.nums[len(r.nums)-1]
		r.valIndexMap[r.nums[index]] = index
	}
	r.nums = r.nums[:len(r.nums)-1]
	delete(r.valIndexMap, val)
	return true
}

func (r *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(r.nums))
	return r.nums[index]
}
