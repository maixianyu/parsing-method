package breadthFirst

import(
	. "container/list"

	"github.com/maixianyu/parsing-method/common"
)

func predict(analysis *List, predition *List, sym2nt map[string]common.NonTerminal) {
	// ep for element of predition stack
	// ea for element of analysis stack
	for ep, ea := predition.Front(), analysis.Front(); ep != nil; ep, ea = ep.Next(), ea.Next() {
		pStack := ep.Value.([]string)
		if nt, found := sym2nt[pStack[0]]; found {
			// nt to r-h sides
			rhSides := nt.RHSides
			aStack := ea.Value.([]string)
			// append each rh to analysis List
			for _, rh := range rhSides {
				appendStack := []string{}
				copy(appendStack, aStack)
				appendStack = append(appendStack, rh...)
				analysis.InsertBefore(appendStack, ea)
			}
			// remove original stack in analysis List
			mark := ea.Prev()
			analysis.Remove(ea)
			ea = mark
		}
	}
}

//func parse(gram common.Grammar, input string) ([]string, bool) {
//
//}
