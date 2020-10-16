package models

func lengthOfLongestSubstring(s string) int {
	MaxINT := 0
	//Output := 0
	Conut := 0
	var strTMP string
	mapTMP := make(map[string]int)
	mapStrTMP := make(map[string]int)
	for k := range s {
		strTMP = strTMP + string(s[k])
		for l := range strTMP {
			mapStrTMP[string(strTMP[l])] = l
		}
		
		_, ok := mapTMP[strTMP]
		if ok {
			continue
		} else {
			mapTMP[string(s[k])] = k
			Conut = Conut + 1
			if Conut >= MaxINT {
				MaxINT = Conut
			}
		}
	}
	return MaxINT
}
