package unger

import (
	"log"

	. "github.com/maixianyu/parsing-method/common"
)

func parse(gram Grammar, input string) ([]string, bool) {
	// try to match start symbol
	trace, ok := matchNonTerminal(gram, gram.StartSymbol, input)
	res := []string{}
	if ok == true {
		for _, t := range trace {
			res = append(res, t + " ->")
		}
		return res, true
	} else {
		return []string{"fail to parse the input."}, false
	}
}

func matchTerminal(symb string, input string) (string, bool) {
	if symb == input {
		return input, true
	} else {
		return input, false
	}
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
		res, ok := matchRightHandSide(gram, rhSide, input)
		if ok == true {
			trace = append(trace, symb)
			trace = append(trace, res...)
			return trace, true
		}
	}

	return []string{}, false
}


func matchRightHandSide(gram Grammar, rhSide RightHandSide, input string) ([]string, bool) {
	numSymbol := len(rhSide)
	inputPartitions := GeneratePartitions(input, numSymbol)

	// step 1: eliminate some partitions unmatch with terminal
	inputPartitions = grepWithTerminal(gram, inputPartitions, rhSide)
	if len(inputPartitions) == 0 {
		return []string{}, false
	}


	// step 2: match terminals or non-terminals in right-hand side
	traces := [][]string{}
	for _, part := range inputPartitions {
		for idx, s := range rhSide {
			_, ok := gram.Symb2NTerminal[s]
			if ok == true {
				// compared with non-terminal
				matchTrace, matched := matchNonTerminal(gram, s, part[idx])
				if matched == true {
					traces = append(traces, matchTrace)
				} else {
					return []string{}, false
				}

			} else {
				// compared with terminal
				matchTrace, matched := matchTerminal(s, part[idx])
				if matched == true {
					traces = append(traces, []string{matchTrace})
				} else {
					return []string{}, false
				}
			}

		}
	}

	// step 3: generate []string{} from trace and rhSide
	res := CombTraceWithTemplate(traces, rhSide)
	return res, true
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
		if found == false {
			idxTerminal = append(idxTerminal, idx)
		}
	}

	// return if no terminals
	if len(idxTerminal) == 0 {
		return partitions
	}
	
	res := [][]string{}
	// match terminal
	Parts:
	for _, p := range partitions {
		for _, idxTerm := range idxTerminal {
			if p[idxTerm] != symbols[idxTerm] {
				continue Parts
			}
		}
		res = append(res, p)
	}
	return res
}
