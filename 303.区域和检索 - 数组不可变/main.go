package main

type NumArray struct {
	MyArray []int
}

// 自己写的
// func Constructor(nums []int) NumArray {
// 	return NumArray{MyArray: nums}
// }

// 自己写的
// func (this *NumArray) SumRange(i int, j int) int {
// 	mySum := 0
// 	sumFlag := false
// 	for index, eachValue := range this.MyArray {
// 		if index == i {
// 			sumFlag = true
// 		}
// 		if index == j {
// 			sumFlag = false
// 			mySum += eachValue
// 		}
// 		if sumFlag {
// 			mySum += eachValue
// 		}
// 	}

// 	return mySum
// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

//
func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for index, value := range nums {
		preSum[index+1] = preSum[index] + value
	}
	return NumArray{MyArray: preSum}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.MyArray[j+1] - this.MyArray[i]
}
