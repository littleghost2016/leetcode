package main

func maxProfit(prices []int) int {

	currentMin, currentMax := 10000, 0
	currentProfit := 0
	myMaxProfit := 0

	for _, eachPrice := range prices {
		if eachPrice < currentMin {
			currentMax = eachPrice
			currentMin = eachPrice
		} else if eachPrice > currentMax {
			currentMax = eachPrice
			currentProfit = currentMax - currentMin
			if currentProfit > myMaxProfit {
				myMaxProfit = currentProfit
			}
		}
	}

	return myMaxProfit
}

// 执行用时：176 ms, 在所有 Go 提交中击败了13.46%的用户
// 内存消耗：8.1 MB, 在所有 Go 提交中击败了58.42%的用户
// func maxProfit(prices []int) int {
//
// 	currentMin := 10000
// 	maxProfit := 0
// 	for _, eachPrice := range prices {
// 		if currentProfit := eachPrice - currentMin; currentProfit > maxProfit {
// 			maxProfit = currentProfit
// 		} else if eachPrice < currentMin {
// 			currentMin = eachPrice
// 		}
// 	}

// 	return maxProfit
// }
