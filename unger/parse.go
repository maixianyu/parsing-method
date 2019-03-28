package unger

import (
	"log"
	"fmt"
	. "github.com/maixianyu/parsing-method/common"
)

type answer struct {
	tf bool
	rec []string
}
var cutOff map[string]answer

func parse(gram Grammar, input string, erule ERule) ([]string, bool) {
	cutOff = map[string]answer{}

	// try to match start symbol
	trace, ok := matchNonTerminal(gram, gram.StartSymbol, input, erule)
	if ok == true {
		res := AppendString2StrSlice(trace, " ->")
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

func matchNonTerminal(gram Grammar, symb string, input string, erule ERule) ([]string, bool) {
	// map NonTerminal from Symbol
	NTerminal, ok := gram.Symb2NTerminal[symb]
	if ok == false {
		log.Fatal("Symol cannot be found in NonTerminal: " + symb)
	}

	ques := symb + "->" + input
	fmt.Println(ques)
	if erule == HasERule {
		/* search for question symb->input first  */
		asw, found := cutOff[ques]
		if found {
			return asw.rec, asw.tf
		} else {
			/* update cutOff map*/
			cutOff[ques] = answer{
				tf: false,
				rec: []string{},
			}
		}
	}

	// match every right-hand side
	trace := []string{}
	fmt.Printf("begin symb:%s, rhSides:%v, input:%s\n", symb, NTerminal.RHSides, input)
	for _, rhSide := range NTerminal.RHSides {
		res, ok := matchRightHandSide(gram, rhSide, input, erule)
		if ok == true {
			fmt.Printf("finish symb:%s, rhSide:%v, input:%s, res:%v\n", symb, rhSide, input, ok)
			trace = append(trace, symb)
			trace = append(trace, res...)
			/* update cutOff map*/
			if erule == HasERule {
				cutOff[ques] = answer{
					tf: true,
					rec: trace,
				}
			}
			return trace, true
		}
	}
	fmt.Printf("finish symb:%s, input:%s, res:%v\n", symb, input, false)

	return []string{}, false
}


func matchRightHandSide(gram Grammar, rhSide RightHandSide, input string, erule ERule) ([]string, bool) {
	numSymbol := len(rhSide)
	inputPartitions, ok := GeneratePartitions(input, numSymbol, erule)
	if ok == false {
		return []string{}, false
	}

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
				matchTrace, matched := matchNonTerminal(gram, s, part[idx], erule)
				if matched == true {
					traces = append(traces, matchTrace)
				} else {
					//fmt.Printf("false symb:%s, input:%s\n", s, input)
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

		/* once a rhSide matches input, then break */
		break
	}

	/* step 3: generate []string{} from trace and rhSide */
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
