package pubilc

import (
	"fmt"
	"testing"
)

func Test_ReJsonNacosCon(t *testing.T) {
	strin := "{\"url_a\":\"AAA\",\"url_b\":\"BBB\",\"url_c\":\"CCC1\"}"
	abc, err := ReJsonNacosCon(strin)
	if err != nil {
		fmt.Println("fuck")
		fmt.Println("err")
	}
	fmt.Println("ok")
	fmt.Println(abc.URLA)
	fmt.Println(abc.URLB)
	fmt.Println(abc.URLC)
}
