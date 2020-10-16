package models

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	var output int
	var outTMP string
	yTMP := 1
	xTMP := x
	if x < 0 {
		xTMP = x * -1
		yTMP = -1
	}

	strTMP := strconv.Itoa(xTMP)
	//lenTMP := len(strTMP)
	//if lenTMP%2 == 0 {

	for k := range strTMP {
		outTMP = string(strTMP[k]) + outTMP
	}
	outIntTMP, _ := strconv.Atoi(outTMP)
	output = outIntTMP * yTMP
	if output > 0 && output >= 2147483648 {
		output = 0
	}
	if output < 0 && output < -2147483648 {
		output = 0
	}
	return output
}

func reverse1(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		b := x
		x /= 10
		b = b / 10
		fmt.Println(b)
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}
