package main

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)

	a, b := 0, 0
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		a, b = nums1[i], nums2[j]
		if a == b {
			if len(res) == 0 || a > res[len(res)-1] {
				res = append(res, a)
			}
			i++
			j++
		} else if a < b {
			i++
		} else {
			j++
		}
	}

	return res
}
