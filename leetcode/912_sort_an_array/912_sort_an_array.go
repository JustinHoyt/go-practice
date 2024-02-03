package main

/*
....................
[2,1,3,0,4]
.....M..............
*/
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	midIdx := len(nums) / 2
	left := mergeSort(nums[midIdx:])
	right := mergeSort(nums[:midIdx])

	temp := []int{}
	leftIdx, rightIdx := 0, 0
	for leftIdx < len(left) && rightIdx < len(right) {
		if left[leftIdx] < right[rightIdx] {
			temp = append(temp, left[leftIdx])
			leftIdx++
		} else {
			temp = append(temp, right[rightIdx])
			rightIdx++
		}
	}
	if leftIdx < len(left) {
		temp = append(temp, left[leftIdx:]...)
	} else {
		temp = append(temp, right[rightIdx:]...)
	}

	return temp
}

func main() {}
