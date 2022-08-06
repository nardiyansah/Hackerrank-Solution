package main

func thirdMax(nums []int) int {
	if len(nums) < 3 {
		temp := 0
		for _, v := range nums {
			if v > temp {
				temp = v
			}
		}
		return temp
	}
	return 0
}