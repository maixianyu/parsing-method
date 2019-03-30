package cyk

import (
	//"fmt"
	"strings"
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

/* v and w are refered in the book Parsing Technique */
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

func genPossiblePartitions(vSet []string, wSet []string) []string {
	res := make([]string, 0)
	for _, v := range vSet {
		for _, w := range wSet {
			RhSide := v + " " + w
			res = append(res, RhSide)
		}
	}
	return res
}

func getPossibleDerivations(gram Grammar, vSet []string, wSet []string) []string {
	res := make([]string, 0)
	for _, p := range genPossiblePartitions(vSet, wSet) {
		deri, found := gram.RhSide2NTSymb[p]
		if found {
			res = append(res, deri...)
		}
	}
	return res
}

func parse(gram Grammar, input string) ([]string, bool) {
	rcgTable := buildRecogTable(gram, input)
	N := len(rcgTable)

	/* make sure input has been recognized in recognitionTable */
	if strInSlice(gram.StartSymbol, rcgTable[0][N-1]) == false {
		return nil, false
	} 

	trace, ok := matchNonTerminal(input, gram, rcgTable, gram.StartSymbol, 0, N-1)
	if ok == true {
		res := AppendString2StrSlice(trace, " ->")
		return res, true
	} else {
		return []string{"fail to parse the input."}, false
	}
}

func matchNonTerminal(input string, gram Grammar, rcgTable recogTable, symb string, idxCol int, idxRow int) ([]string, bool){
	/* condition to finish recursive */
	if idxRow == 0 {
		inSlice := []rune(input)
		traces := []string{symb, string(inSlice[idxCol])}
		return traces, true
	}

	/* get possible partitions with the help of rcgTable */
	allPartitions, positions := genPartitions(rcgTable, idxCol, idxRow)

	/* grep matched partitions */
	matchedPartitions := []string{}
	matchedPositions := [][4]int{}
	for rhIdx, rh := range allPartitions {
		if strInSlice(symb, gram.RhSide2NTSymb[rh]) == true {
			matchedPartitions = append(matchedPartitions, rh)
			matchedPositions = append(matchedPositions, positions[rhIdx])
		}
	}

	/* depth-first search, unger-style */
	var allTraces [][]string
	var template []string
	Loop:
	for pIdx, p := range matchedPartitions {
		NTermsInP := strings.Split(p, " ")
		ntTrace := [][]string{}
		for ntIdx, ntSymb := range NTermsInP {
			t, ok := matchNonTerminal(input, gram, rcgTable, ntSymb, matchedPositions[pIdx][ntIdx * 2], matchedPositions[pIdx][ntIdx * 2 + 1])
			if ok {
				ntTrace = append(ntTrace, t)
			} else {
				continue Loop
			}
		}
		allTraces = ntTrace
		template = NTermsInP
		break
	}

	if len(allTraces) == 0 {
		return nil, false
	}

	/* combine traces in template */
	combTraces := CombTraceWithTemplate(allTraces, RightHandSide(template))
	res := []string{symb}
	res = append(res, combTraces...)
	return res, true

}


func genPartitions(rcgTable recogTable, idxCol int, idxRow int) ([]string, [][4]int) {
	res := []string{}
	/* [4]int{ v idxCol, v idxRow, w idxCol, w idxRow } */
	positoins := [][4]int{}

	for i := 0; i < idxRow; i++ {
		vIdxCol, vIdxRow, wIdxCol, wIdxRow := idxCol, i, idxCol+i+1, idxRow-i-1
		vSet := rcgTable[vIdxCol][vIdxRow]
		wSet := rcgTable[wIdxCol][wIdxRow]
		partitions := genPossiblePartitions(vSet, wSet)
		for _, p := range partitions {
			res = append(res, p)
			positoins = append(positoins, [4]int{vIdxCol, vIdxRow, wIdxCol, wIdxRow})
		}
	}
	return res, positoins
}


func strInSlice(s string, strSlice []string) bool {
	if strSlice == nil || len(strSlice) == 0 {
		return false
	} else {
		for _, v := range strSlice {
			if s == v {
				return true
			}
		}
		return false
	}
}

func matchTerminal(symb string, input string) (string, bool) {
	if symb == input {
		return input, true
	} else {
		return input, false
	}
}