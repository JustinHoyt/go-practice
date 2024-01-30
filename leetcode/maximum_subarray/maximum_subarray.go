package main

import (
	"math"
)

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	high := math.MinInt
	leftIdx := 0 
	sum := 0
	for rightIdx := 0; rightIdx < len(nums); rightIdx++ {
		sum += nums[rightIdx]
		high = max(high, sum)

		if sum < 0 {
			for ; leftIdx <= rightIdx; leftIdx++ {
				sum -= nums[leftIdx]
			}
		}

	}

	return high
}

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	high := math.MinInt
	sum := 0
	for rightIdx := 0; rightIdx < len(nums); rightIdx++ {
		sum += nums[rightIdx]
		high = max(high, sum)

		sum = max(sum, 0)
		if sum < 0 {
			sum = 0
		}
	}

	return high
}

func maxSubArray3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	high := math.MinInt
	sum := 0
	for rightIdx := 0; rightIdx < len(nums); rightIdx++ {
		sum += nums[rightIdx]
		high = max(high, sum)
		sum = max(sum, 0)
	}

	return high
}


func main() {
	maxSubArray1([]int{-2,1,-3,4,-1,2,1,-5,4})
}
