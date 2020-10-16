package models

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	var input1 int
	input1 = 110
	res := isPalindrome(input1)
	fmt.Println(res)
}
