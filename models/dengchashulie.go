package models

func canMakeArithmeticProgression(arr []int) bool {
	// var arrTMP []int
	//	var intTMP int
	//	intTMP = 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}

	sum := arr[0] - arr[1]
	for j := range arr {
		if j == 0 {
			continue
		}
		if sum != arr[j-1]-arr[j] {
			return false
		}
	}
	return true
}
