package models

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	var input1 int
	input1 = -1234123123
	res := reverse(input1)
	fmt.Println(res)
}

func Test_reverse1(t *testing.T) {
	var input1 int
	input1 = -123412
	res := reverse1(input1)
	fmt.Println(res)
}
