package main

func canCompleteCircuit(gas []int, cost []int) int {
	lowestSum := 0
	sum := 0
	for i := range gas {
		sum += gas[i] - cost[i]
		lowestSum = min(sum, lowestSum)
	}
	if lowestSum >= 0 {
		return 0
	}

	for i := len(gas) - 1; i > -1; i-- {
		lowestSum += gas[i] - cost[i]

		if lowestSum >= 0 {
			return i
		}
	}
	return -1
}
