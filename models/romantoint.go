package models

func romanToInt(s string) int {
	var output int
	if len(s) == 1 {
		output = baseCase(s)
		return output
	}
	for k := range s {
		if k == 0 {
			output = baseCase(string(s[k]))
			continue
		}
		secINT := baseCase(string(s[k]))
		firINT := baseCase(string(s[k-1]))
		if firINT < secINT {
			output = secINT - firINT*2 + output
			continue
		}
		output = output + secINT
	}
	return output
}

func baseCase(input string) (output int) {
	switch input {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}
