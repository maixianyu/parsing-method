package breadthFirst

import(
	"container/list"
	"strings"

	"github.com/maixianyu/parsing-method/common"
)

func predict(analysis, prediction *list.List, sym2nt map[string]common.NonTerminal) bool {
	remaining := false
	// ep for element of predition stack
	// ea for element of analysis stack
	for ep, ea := prediction.Front(), analysis.Front(); ep != nil; ep, ea = ep.Next(), ea.Next() {
		pStack := ep.Value.([]string)
		if nt, found := sym2nt[pStack[0]]; found {
			remaining = true
			// nt to r-h sides
			rhSides := nt.RHSides
			aStack := ea.Value.([]string)
			for _, rh := range rhSides {
				// move non-terminal from prediction stack to analysis stack
				ntStack := []string{}
				copy(ntStack, aStack)
				ntStack = append(ntStack, nt.Symbol)
				analysis.InsertBefore(ntStack, ea)

				// push each rh to prediction stack
				rhStack := []string{}
				copy(rhStack, pStack)
				rhStack = append(rh, pStack[1:]...)
				prediction.InsertBefore(rhStack, ep)
			}

			// remove original stack in analysis/prediction List
			markA, markP := ea.Prev(), ep.Prev()
			analysis.Remove(ea)
			analysis.Remove(ep)
			ea, ep = markA, markP

		}
	}
	return remaining
}

func match(matched, rest *[]string, analysis, prediction *list.List, sym2nt map[string]common.NonTerminal) bool {
	if len(*rest) == 0 {
		return false
	}

	remaining := true

	// remove mismatch stack
	for ep, ea := prediction.Front(), analysis.Front(); ep != nil; {
		pStack := ep.Value.([]string)
		_, found := sym2nt[pStack[0]]
		remaining = remaining && found
		if pStack[0] != (*rest)[0] {
			// mismatch: remove stack
			markA, markP := ep, ea
			ep, ea = ep.Next(), ea.Next()
			analysis.Remove(markA)
			prediction.Remove(markP)
		} else {
			// match: move symbol from prediction to analysis
			aStack := ea.Value.([]string)
			aStack = append(aStack, pStack[0])
			ea.Value = pStack[1:]
			ep, ea = ep.Next(), ea.Next()
		}
	}

	// move symbol from rest to matched
	*matched = append(*matched, (*rest)[0])
	if len(*rest) == 1 {
		*rest = []string{}
	} else {
		*rest = (*rest)[1:]
	}

	return remaining
}

func parse(gram common.Grammar, input string) ([]string, bool) {
	split := strings.Split(input, " ")
	split = append(split, "#")
	rest := []string{}
	for _, s := range split {
		if s != "" {
			rest = append(rest, s)
		}
	}
	matched := []string{}
	// prediction List
	preStack := []string{gram.StartSymbol, "#"}
	prediction := list.New()
	prediction.PushFront(preStack)
	// analysis List
	anaStack := []string{}
	analysis := list.New()
	analysis.PushFront(anaStack)

	for len(rest) != 0 {
		for predict(analysis, prediction, gram.Symb2NTerminal) == true {}
		for match(&matched, &rest, analysis, prediction, gram.Symb2NTerminal) == true {}
	}

	return analysis.Front().Value.([]string), analysis.Len() > 0
}