package pubilc

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

func ReJsonNacosCon(input string) (output ConfigTest, error error) {
	var datTag ConfigTest
	if resjsonTag := json.Unmarshal([]byte(input), &datTag); resjsonTag != nil {
		beego.Error(resjsonTag, input)
		return output, resjsonTag
	}
	output = datTag
	return output, nil
}
