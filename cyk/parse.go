package cyk

import (
	//"fmt"
	. "github.com/maixianyu/parsing-method/common"
)

/* col-row-array of non-terminals */
type recogTable [][][]string

/* build a blank table */
func buildRecogTable(gram Grammar, input string) recogTable {
	symbs := []rune(input)

	/* generate a blank table */
	rcgTable := make([][][]string, len(symbs))
	for idxCol, _ := range symbs {
		rcgTable[idxCol] = make([][]string, len(symbs))
	}

	/* fill in the first row */
	for idxCol, s := range symbs {
		ntTerminals, found := gram.RhSide2NTSymb[string(s)]
		if found {
			rcgTable[idxCol][0] = append(rcgTable[idxCol][0], ntTerminals...)
		}
	}

	/* fill in other rows. Works the way up*/
	for idxRow := 1; idxRow < len(symbs); idxRow++ {
		for idxCol := 0; idxCol < len(symbs) - idxRow; idxCol++ {
			rcgTable[idxCol][idxRow] = calSetOfNterminals(gram, rcgTable, idxCol, idxRow)
		}
	}

	return rcgTable
}

func calSetOfNterminals(gram Grammar, rcgTable recogTable, idxCol int, idxRow int) []string {
	res := make([]string, 0)
	for i := 0; i < idxRow; i++ {
		vSet := rcgTable[idxCol][i]
		wSet := rcgTable[idxCol+i+1][idxRow-i-1]
		/* get possibleDerivations */
		possibleDeri := getPossibleDerivations(gram, vSet, wSet)
		res = append(res, possibleDeri...)
	}
	return res
}

func getPossibleDerivations(gram Grammar, vSet []string, wSet []string) []string {
	res := make([]string, 0)
	for _, v := range vSet {
		for _, w := range wSet {
			RhSide := v + " " + w
			deri, found := gram.RhSide2NTSymb[RhSide]
			if found {
				res = append(res, deri...)
			}
		}
	}
	return res
}


func matchTerminal(symb string, input string) (string, bool) {
	if symb == input {
		return input, true
	} else {
		return input, false
	}
}