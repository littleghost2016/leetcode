package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if index, ok := hashMap[target-nums[i]]; ok {

			// 以下两种都可以
			// return []int{i, index}
			return []int{index, i}
		} else {
			hashMap[nums[i]] = i
		}
	}

	return nil
}
