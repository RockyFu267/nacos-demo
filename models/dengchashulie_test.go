package models

import (
	"fmt"
	"testing"
)

func Test_canMakeArithmeticProgression(t *testing.T) {
	var input1 []int
	input1 = append(input1, 1, 2, 4)
	res := canMakeArithmeticProgression(input1)
	fmt.Println(res)
}
