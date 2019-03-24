package unger

import (
	"log"

	. "github.com/maixianyu/parsing-method/common"
)

func parse(gram Grammar, input string) []string {
	trace := []string{}
	// try to match start symbol
	trace, ok := matchNonTerminal(gram, gram.StartSymbol, input)
	return trace
}

func matchNonTerminal(gram Grammar, symb string, input string) ([]string, bool) {
	trace := []string{}
	// map NonTerminal from Symbol
	NTerminal, ok := gram.Symb2NTerminal[symb]
	if ok == false {
		log.Fatal("Symol cannot be found in NonTerminal: " + symb)
	}

	// match every right-hand side
	for _, rhSide := range NTerminal.RHSides {
		res, ok := matchRightHandSide(gram, rhSide, input, []string{})
		if ok == true {
			trace = append(trace, res...)
			return trace, true
		}
	}

	return []string{}, false
}

func matchTerminal(symb string, input string) bool {
	if symb == input {
		return true
	} else {
		return false
	}
}

func matchRightHandSide(gram Grammar, rhSide RightHandSide, input string, trace []string) ([]string, bool) {
	numSymbol := len(rhSide.Symbols)
	inputPartitions := generatePartitions(input, numSymbol)

	// eliminate some partitions unmatch with terminal
	inputPartitions = grepWithTerminal(gram, inputPartitions, rhSide.Symbols)

	res := trace

	for _, part := range inputPartitions {
		// each symbol in right-hand side expression
		for idx, s := range rhSide.Symbols {
			nt, ok := gram.Symb2NTerminal[s]
			if ok == true {
				// compared with non-terminal
				matchStr, matched := matchNonTerminal(gram, s, p)
				if matched == true {
					res = append(res, matchStr...)
				} else {
					return []string{}, false
				}

			} else {
				// compared with terminal
				resTerminal := matchTerminal(s, part[idx])
				if resTerminal == true {
					res = append(res, s)
				} else {
					return []string{}, false
				}
			}

		}
	}

	return []string{}, false
}

func genPartitionsHelper(preRes []string, input string, num int) [][]string {
	if num == 0 || len(input) < num {
		return [][]string{}
	}
	inputSlice := []rune(input)
	var res = [][]string{}
	for i := 1; i<=num; i++ {
		curRes := inputSlice[0:i]
		res = append(res, genPartitionsHelper(append(preRes, string(curRes)), string(inputSlice[i:]), num-1)...)
	}
	return res
}

func generatePartitions(input string, num int) [][]string{
	if num == 0 || len(input) < num {
		return [][]string{}
	}
	return genPartitionsHelper([]string{}, input, num)
}

func grepWithTerminal(gram Grammar, partitions [][]string, symbols []string) [][]string {
	if len(partitions) == 0 {
		return [][]string{}
	} else if len(partitions[0]) != len(symbols) {
		return [][]string{}
	}

	// get terminal idx
	idxTerminal := []int{}
	for idx, s := range symbols {
		_, found := gram.Symb2NTerminal[s]
		if found == true {
			idxTerminal = append(idxTerminal, idx)
		}
	}
	
	res := [][]string{}
	// match terminal
	for _, p := range partitions {
		for _, idxTerm := range idxTerminal {
			if p[idxTerm] != symbols[idxTerm] {
				break
			}
		}
		res = append(res, p)
	}
	return res
}
