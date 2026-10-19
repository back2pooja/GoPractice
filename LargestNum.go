package main

import (
	"errors"
)

func largestNum(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty number")
	}
	maxnum := nums[0]
	for _, n := range nums {
		if n > maxnum {
			maxnum = n
		}
	}
	return maxnum, nil
}
