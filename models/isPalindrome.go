package models

import (
	"strconv"
)

func isPalindrome(x int) bool {
	strTMP := strconv.Itoa(x)
	resTMP := ""
	for k := range strTMP {
		resTMP = string(strTMP[k]) + resTMP
	}
	if resTMP == strTMP {
		return true
	}
	return false
}
