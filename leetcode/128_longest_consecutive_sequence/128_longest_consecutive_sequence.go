package main

func longestConsecutive(nums []int) int {
	sequences := map[int]bool{}
	high := 0

	countConsecutive := func(num int) int {
		count := 1
		for i := num + 1; sequences[i]; i++ {
			sequences[i] = false
			count++
		}
	
		for i := num - 1; sequences[i]; i-- {
			sequences[i] = false
			count++
		}
		return count
	}

	for _, num := range nums {
		sequences[num] = true
	}

	for key, val := range sequences {
		if val == true {
			high = max(high, countConsecutive(key))
		}
	}

	return high
}

func main() {}
