package models

import (
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var output *ListNode
	var strL1, strL2 string

	for true {
		strTMP := strconv.Itoa(l1.Val)
		strL1 = strL1 + strTMP
		if l1.Next == nil {
			break
		}
	}
	for true {
		strTMP := strconv.Itoa(l2.Val)
		strL2 = strL2 + strTMP
		if l2.Next == nil {
			break
		}
	}
	intL1, _ := strconv.Atoi(strL1)
	intL2, _ := strconv.Atoi(strL2)
	intSum := intL1 + intL2
	strSum := strconv.Itoa(intSum)
	lenStr := len(strSum)
	for k := range strSum {
		output.Val, _ = strconv.Atoi(string(strSum[k]))
		if k == lenStr {
			return output
		}
		output.Next.Val, _ = strconv.Atoi(string(strSum[k+1]))
	}
	return output
}
