package common

import (
	//"fmt"
)

func genPartitionsHelper(preRes []string, input string, num int) ([][]string, bool) {
	inputSlice := []rune(input)

	// set condition to finish
	if num == 0 {
		if len(inputSlice) == 0 {
			return [][]string{preRes}, true
		} else {
			return [][]string{}, false
		}
	} else {
		if num > len(inputSlice) {
			return [][]string{}, false
		}
	}

	// loop and recursive
	var res = [][]string{}
	N := len(inputSlice)

	// allocate a new string slice and copy preRes into it
	fixed := make([]string, len(preRes))
	copy(fixed, preRes)

	for i := 1; i<=N; i++ {
		curRes := inputSlice[0:i]
		nxtRes := append(fixed, string(curRes))

		subRes, ok := genPartitionsHelper(nxtRes, string(inputSlice[i:N]), num-1)
		if ok == true {
			res = append(res, subRes...)
		}
	}
	return res, true
}

func generatePartitions(input string, num int) [][]string{
	if num <= 0 || len(input) < num {
		return [][]string{}
	}
	res, _ := genPartitionsHelper([]string{}, input, num)
	return res
}