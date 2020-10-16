package models

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	var input1 string
	input1 = "MCMXCIV"
	res := romanToInt(input1)
	fmt.Println(res)
}
