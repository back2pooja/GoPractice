package main

func summonArrayElement(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
