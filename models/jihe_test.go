package models

import (
	"fmt"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	var input1 []int
	input1 = append(input1, 2, 2)
	fmt.Println(input1)
	res := findErrorNums(input1)
	fmt.Println(res)
}
