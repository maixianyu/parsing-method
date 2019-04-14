package backtracking

import(
	"strings"
	"github.com/maixianyu/parsing-method/common"
)

// notations below are explained in Parsing Technique
// input = matched + rest
// matched: matched parts of input
// rest: parts of input, about to be matched
// analysis: used to record parsing result
// prediction: used to predict parsing
func depthFirst(matched, rest, analysis, prediction []string, symb2nt map[string]common.NonTerminal, res *[][]string) {
	nt, found := symb2nt[prediction[0]]

	if len(rest) < len(prediction) {
		return
	}
	if len(rest) == 1 && found == false && len(prediction) == 1 && rest[0] == prediction[0] {
		analysis = append(analysis, rest[0])
		*res = append(*res, analysis)
		return
	} 

	// current symbol is non-terminal
	if found {
		// construct new analysis
		newAly := append(analysis, prediction[0])
		// iter right-hand side
		for _, rh := range nt.RHSides {
			// construct new predition
			newPred := make([]string, 0, len(rh) + len(prediction) - 1)
			newPred = append(newPred, rh...)
			if len(prediction) != 1 {
				newPred = append(newPred, prediction[1:]...)
			}
			depthFirst(matched, rest, newAly, newPred, symb2nt, res)
		}

	// current symbol is terminal
	} else {
		// matched
		if rest[0] == prediction[0] {
			newRest := rest[1:]
			var newPred []string
			if len(prediction) == 1 {
				newPred = prediction[0:0]
			} else {
				newPred = prediction[1:]
			}
			newMch := append(matched, rest[0])
			newAly := append(analysis, rest[0])
			depthFirst(newMch, newRest, newAly, newPred, symb2nt, res)

		// mismatched
		} else {
			return
		}
	}
}

func parse(gram common.Grammar, input string) ([][]string, error) {
	split := strings.Split(input, " ")
	split = append(split, "#")
	rest := []string{}
	for _, s := range split {
		if s != "" {
			rest = append(rest, s)
		}
	}
	prediction := []string{gram.StartSymbol, "#"}
	matched, analysis := []string{}, []string{}
	res := [][]string{}
	depthFirst(matched, rest, analysis, prediction, gram.Symb2NTerminal, &res)
	return res, nil
}