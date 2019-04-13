package breadthFirst

import(
	"container/list"
	"strings"
	"errors"
	"fmt"
	"github.com/maixianyu/parsing-method/common"
)

func predict(analysis, prediction *list.List, sym2nt map[string]common.NonTerminal) bool {
	noNonterminal := true
	// ep for element of predition stack
	// ea for element of analysis stack
	for ep, ea := prediction.Front(), analysis.Front(); ep != nil && ea != nil; ep, ea = ep.Next(), ea.Next() {
		pStack := ep.Value.([]string)
		if nt, found := sym2nt[pStack[0]]; found {
			noNonterminal = false
			// nt to r-h sides
			rhSides := nt.RHSides
			aStack := ea.Value.([]string)
			for _, rh := range rhSides {
				// move non-terminal from prediction stack to analysis stack
				ntStack := make([]string, len(aStack), 2 * len(aStack))
				copy(ntStack, aStack)
				ntStack = append(ntStack, nt.Symbol)
				analysis.InsertBefore(ntStack, ea)

				// push each rh to prediction stack
				rhStack := make([]string, len(pStack), 2 * len(pStack))
				copy(rhStack, pStack)
				rhStack = append(rh, pStack[1:]...)
				prediction.InsertBefore(rhStack, ep)
			}

			// remove original stack in analysis/prediction List
			markA, markP := ea.Prev(), ep.Prev()
			analysis.Remove(ea)
			prediction.Remove(ep)
			ea, ep = markA, markP

		}
	}
	return noNonterminal
}

func match(matched, rest *[]string, analysis, prediction *list.List, sym2nt map[string]common.NonTerminal) bool {
	if len(*rest) == 0 {
		return true
	}

	// ready to predict
	ready := false

	// test if all symbols are teminals
	for ep := prediction.Front(); ep != nil; ep = ep.Next() {
		pStack := ep.Value.([]string)
		_, found := sym2nt[pStack[0]]
		ready = ready || found
	}
	if ready == true {
		return ready
	}

	// remove mismatch stack
	for ep, ea := prediction.Front(), analysis.Front(); ep != nil; {
		pStack := ep.Value.([]string)
		if pStack[0] != (*rest)[0] {
			// mismatch: remove stack
			markA, markP := ea, ep
			ep, ea = ep.Next(), ea.Next()
			analysis.Remove(markA)
			prediction.Remove(markP)
		} else {
			// match: move symbol from prediction to analysis
			ea.Value = append(ea.Value.([]string), pStack[0])
			ep.Value = pStack[1:]
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

	return ready
}

func parse(gram common.Grammar, input string) ([]string, error) {
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
		for predict(analysis, prediction, gram.Symb2NTerminal) == false {}
		for match(&matched, &rest, analysis, prediction, gram.Symb2NTerminal) == false {}
	}

	var err error
	if analysis.Len() == 0 {
		info := "fail to parse " + input
		err = errors.New(info)
	}

	return analysis.Front().Value.([]string), err
}

func printList(l1 *list.List, l2 *list.List) {
	for e1, e2 := l1.Front(), l2.Front(); e1 != nil && e2 != nil; e1, e2 = e1.Next(), e2.Next() {
		fmt.Printf("e1:%v, e2:%v\n", e1.Value, e2.Value)
	}
}