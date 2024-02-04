package main

/*
....................
[2,1,3,0,4]
.....M..............
*/
func mergeSort(nums []int) (res []int) {
	if len(nums) < 2 {
		return nums
	}

	m := len(nums) / 2
	left := mergeSort(nums[m:])
	right := mergeSort(nums[:m])

	l, r := 0, 0
	for l < len(left) || r < len(right) {
		switch {
		case l == len(left):
			res = append(res, right[r:]...)
			r = len(right)
		case r == len(right):
			res = append(res, left[l:]...)
			l = len(left)
		case left[l] < right[r]:
			res = append(res, left[l])
			l++
		case right[r] <= left[l]:
			res = append(res, right[r])
			r++
		}
	}

	return res
}

func main() {}
