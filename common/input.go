package common

import (
	//"fmt"
	"strings"
)

func SliceInput(input string, sep string) []string {
	in := strings.Split(input, sep)
	res := make([]string, 0, len(in))
	for _, s := range in {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}

/* Generate n partitions from input */
func GeneratePartitions(input string, num int, erule ERule) ([][]string, bool) {
	inputSlice := []rune(input)

	// set condition to finish recursive
	// true: generated successfully
	// false: cannot be partitoined
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
		nxtRes, ok := GeneratePartitions(string(inputSlice[i:N]), num-1, erule)
		if ok == true {
			// nxtRes length is 0
			if len(nxtRes) == 0 {
				combRes := []string{string(curRes)}
				res = append(res, combRes)
			// nxtRes length > 0
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

/* Generate M partitions from []string */
func AssemblyPartitions(text []string, M int, erule ERule) ([][]string, bool) {
	// string in []string is not allowed to include space
	for _, s := range text {
		if strings.Contains(s, " ") == true {
			panic("string in []string is not allowed to include space")
		}
	}

	// set conditions to teminate recursive
	N := len(text)
	if M == N {
		return [][]string{text}, true
	} else if M < N {
		return nil, false
	}

	// loop from 1 to M-(N-1) for first string text[0]
	res := make([][]string, 0, M-(N-1))
	for i := 1; i <= M-(N-1); i++ {
		curRes, curV := GeneratePartitions(text[0], i, NoERule)
		nxtRes, nxtV := [][]string{}, false
		if N > 1 {
			nxtRes, nxtV = AssemblyPartitions(text[1:], M-i, NoERule)
		}
		// be careful of N of the last element in text
		// when no more text need to assembly (M-i==0)
		if N == 1 && M - i == 0 {
			nxtRes, nxtV = [][]string{}, true
		}

		// both of them are valid
		if (curV && nxtV) == true {
			if len(nxtRes) == 0 {
				res = append(res, curRes...)
			} else {
				for _, c := range curRes {
					for _, n := range nxtRes {
						newRes := append(c, n...)
						res = append(res, newRes)
					}
				}
			}
		}
	}

	// if epslon-rule is applied
	if erule == HasERule {
		// cur is "" because of epslon, so M-1 is passed to next
		nxtRes, v := AssemblyPartitions(text, M-1, NoERule)
		if v == true {
			for _, n := range nxtRes {
				newRes := make([]string, 0, 1 + len(n))
				// cur is "" because of epslon
				newRes = append(newRes, "")
				newRes = append(newRes, n...)
				res = append(res, newRes)
			}
		}
	}

	return res, true
}