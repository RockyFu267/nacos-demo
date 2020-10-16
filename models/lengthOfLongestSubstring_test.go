package models

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	var input1 string
	input1 = "pwwkew"
	res := lengthOfLongestSubstring(input1)
	fmt.Println(res)
}
