package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.com/problems/4sum/

type Key [4]int

func (k *Key) Less(i, j int) bool {
	return k[i] < k[j]
}

func (k *Key) Len() int {
	return 4
}

func (k *Key) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func fourSum(nums []int, target int) [][]int {
	var output [][]int
	answerMap := make(map[Key]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				for m := k + 1; m < len(nums); m++ {
					sum := nums[i] + nums[j] + nums[k] + nums[m]
					if sum != target {
						continue
					}

					newSlice := Key{nums[i], nums[j], nums[k], nums[m]}
					sort.Sort(&newSlice)
					if _, ok := answerMap[newSlice]; ok {
						continue
					}
					answerMap[newSlice] = struct{}{}
					output = append(output, []int{nums[i], nums[j], nums[k], nums[m]})
				}
			}
		}
	}
	return output
}

/* testing */

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
		},
		{
			nums:   []int{2, 2, 2, 2, 2, 2},
			target: 8,
		},
		{
			nums:   []int{2, 6},
			target: 8,
		},
	}
	for _, test := range tests {
		res := fourSum(test.nums, test.target)
		fmt.Println(res)
	}
}
