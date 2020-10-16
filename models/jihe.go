package models

import (
	"encoding/json"
	"fmt"
)

func findErrorNums(nums []int) []int {
	var silceTMP map[int]int
	var output []int
	silceTMP = make(map[int]int)

	max := 0
	lenTMP := len(nums)
	fmt.Println(lenTMP)
	for k := range nums {
		if k == 0 {
			max = nums[k]
		}
		if max < nums[k] {
			max = nums[k]
		}
		if _, ok := silceTMP[nums[k]]; ok {
			output = append(output, nums[k])
		} else {
			silceTMP[nums[k]] = k
		}
	}
	res2B, _ := json.Marshal(silceTMP)
	fmt.Println(string(res2B))
	fmt.Println(max)
	fmt.Println(output)
	if len(silceTMP) == 1 {
		output = append(output, max-1)
		return output
	}
	sum := 1 + lenTMP
	for i := 1; i < sum; i++ {
		if _, ok := silceTMP[i]; ok {
			continue
		} else {
			output = append(output, i)
		}
	}
	return output
}
