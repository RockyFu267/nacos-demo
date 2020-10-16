package models

import (
	"fmt"
)

func TwoSum() {
	var nums, nums1 []int
	var num, num1 int
	nums1 = []int{2, 7, 5, 9, 4, 8}
	num1 = 17
	nums = []int{3, 2, 4}
	num = 6
	res := twoSum(nums, num)
	res1 := twoSum1(nums1, num1)
	fmt.Println(res)
	fmt.Println(res1)
	fmt.Println("OK")
}

func twoSum(nums []int, target int) []int {
	var Output []int
	for k := range nums {
		if k == 0 {
			continue
		}
		// var tmpS []int
		// tmpS = nums[k:]
		for l := range nums[k:] {
			if target == nums[k-1]+nums[k:][l] {
				Output = append(Output, k-1)
				Output = append(Output, l+k)
				return Output
			}
		}
	}
	return Output
}

func twoSum1(nums []int, target int) []int {
	//var mapTMP map[int]int
	mapTMP := make(map[int]int)
	var Output []int
	for k := range nums {
		intTMP := target - nums[k]
		_, m := mapTMP[intTMP]
		if m {
			Output = append(Output, mapTMP[intTMP])
			Output = append(Output, k)
		} else {
			mapTMP[nums[k]] = k
		}
	}
	return Output
}

func twoSum2(nums []int, target int) []int {
	mapTMP := make(map[int]int)
	var Output []int
	for k := range nums {
		intTMP := target - nums[k]
		_, m := mapTMP[intTMP]
		if m {
			Output = append(Output, mapTMP[intTMP], k)
		} else {
			mapTMP[nums[k]] = k
		}
	}
	return Output
}
