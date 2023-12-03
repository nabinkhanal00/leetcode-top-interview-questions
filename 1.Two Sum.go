package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, element := range nums {
		if val, ok := m[target-element]; ok {
			return []int{val, index}
		}
		m[element] = index
	}
	return []int{}
}
