package common

import (
	//"fmt"
)

func genPartitionsHelper(input string, num int, erule ERule) ([][]string, bool) {
	inputSlice := []rune(input)

	// set condition to finish recursive
	if num == 0 {
		if len(inputSlice) == 0 {
			return [][]string{}, true
		} else {
			return [][]string{}, false
		}
	} else {
		if erule == HasERule && num == 1 && input == "" {
			return [][]string{
				[]string{""},
				}, true
		} else if erule == NoERule && num > len(inputSlice) {
			return [][]string{}, false
		}
	}

	// startIdx depends on e-rule
	var startIdx int
	if erule == HasERule && num != 1 {
		startIdx = 0
	} else {
		startIdx = 1
	}

	/* combine current string and next sub string into result*/
	/* current string = "xx" */
	/* next string = [][]string{
		[]string{"a", "bc"},
		[]string{"ab", "c"},
	}
	   res string = [][]string{
		[]string{"xx", "a", "bc"},
		[]string{"xx", "ab", "c"},
	   }
	*/
	var res = [][]string{}
	N := len(inputSlice)
	for i := startIdx; i<=N; i++ {
		curRes := inputSlice[0:i]
		nxtRes, ok := genPartitionsHelper(string(inputSlice[i:N]), num-1, erule)
		if ok == true {
			if len(nxtRes) == 0 {
				combRes := []string{string(curRes)}
				res = append(res, combRes)
			} else {
				for _, nxt := range nxtRes {
					combRes := []string{string(curRes)}
					combRes = append(combRes, nxt...)
					res = append(res, combRes)
				}
			}
		}
	}
	return res, true
}

func GeneratePartitions(input string, num int, erule ERule) [][]string{
	if num <= 0 || len(input) < num {
		return [][]string{}
	}
	res, _ := genPartitionsHelper(input, num, erule)
	return res
}