package earley

import(
	"errors"
	"github.com/maixianyu/parsing-method/common"
)

type item struct {
	dotPos int
	// at starting from 1
	at int
	nonTerm string
	rhSide []string
}

/* active set or predict set */
type subSet []item

/* itemSet consists of active set and predict set */
type itemSet [2]subSet

/* create an active set at the beginning */
func initActiveSet(gram common.Grammar) subSet {
	rhs := gram.Symb2NTerminal[gram.StartSymbol].RHSides
	sub := make([]item, len(rhs))
	for i, rh := range rhs {
		sub[i].dotPos = 0
		sub[i].at = 1
		sub[i].nonTerm = gram.StartSymbol
		sub[i].rhSide = rh
	}
	return sub
}

/* a depth-first function to get more predict items from one item input */
func predictHelper(it item, at int, symb2nt map[string]common.NonTerminal, exist *map[string]bool) subSet {
	predSet := []item{}
	symb := it.rhSide[it.dotPos]
	/* if the symb after dot is a non-terminal symb, and it did not exist in predSet */
	_, old := (*exist)[symb]
	if nt, found := symb2nt[symb]; found == true && old == false {
		/* add each of its right-hand side */
		for _, rh := range nt.RHSides {
			predItem := item{
				dotPos: 0,
				at: at,
				nonTerm: symb,
				rhSide: rh,					
			}
			/* add item to predSet */
			predSet = append(predSet, predItem)
			/* mark symb as existed*/
			(*exist)[symb] = true
			/* depth-first to find new non-terminal */
			sub := predictHelper(predItem, at, symb2nt, exist)
			predSet = append(predSet, sub...)
		}
	}
	return predSet
}

/* predictor generates predict set from active set */
func (active *subSet) predictor(at int, symb2nt map[string]common.NonTerminal) (subSet, error) {
	if len((*active)) == 0 {
		return nil, errors.New("subSet is null.")
	}

	predSet := make([]item, 0, len((*active)))
	/* loop active item */
	exist := map[string]bool{}
	for _, it := range (*active) {
		sub := predictHelper(it, at, symb2nt, &exist)
		predSet = append(predSet, sub...)
	}
	return predSet, nil
}

/* scanner looks at symb, goes through an item set, and generate (completed subSet, active subSet) */
func (its *itemSet) scanner(symb string, isNonTerminal bool, symb2nt map[string]common.NonTerminal) (completed subSet, active subSet) {
	completed, active = []item{}, []item{}
	/* loop subSet */
	for i := range *its {
		/* loop item */
		for _, it := range (*its)[i] {
			curSymb := it.rhSide[it.dotPos]
			/* if curSymb is a terminal and equals to symb */
			if _, found := symb2nt[curSymb]; curSymb == symb && found == isNonTerminal {
				/* curSymb is at tail, then recognized as completed */
				if it.dotPos == len(it.rhSide) - 1 {
					completed = append(completed, it)
				/* curSymb is not at tail, then recognized as active */
				} else {
					active = append(active, it)
				}
			}
		}
	}
	return
}

/* completer inspects completed set, which contains the items that have just been recognized and can now 
*  be reduced.
*/
func (completed *subSet) completer(allItemSet []itemSet, symb2nt map[string]common.NonTerminal) (newComp subSet, newActv subSet, err error) {
	newComp, newActv = []item{}, []item{}
	for _, it := range *completed {
		idxSet := it.at
		if idxSet > len(allItemSet) {
			return nil, nil, errors.New("idxSet >= len(allItemSet)")
		}
		searchSet := allItemSet[idxSet-1]
		c, a := searchSet.scanner(it.nonTerm, true, symb2nt)
		newComp = append(newComp, c...)
		newActv = append(newActv, a...)
	}
	return
}

/* parse */
func parse(gram common.Grammar, input string) ([]string, error) {
	in := common.SliceInput(input, " ")
	// make itemSet0
	actv0 := initActiveSet(gram)
	pred0, err := (&actv0).predictor(1, gram.Symb2NTerminal)
	if err != nil {
		return nil, err
	}
	// make allItemSet
	allItemSet := make([]itemSet, 0, len(in) + 1)
	allItemSet = append(allItemSet, itemSet{ actv0, pred0 })
	// make completed array
	compArr := make([]subSet, 0, len(in))

	/* goes through input, and construct itemSet */
	for idx, symb := range in {
		completed, active := (&allItemSet[idx]).scanner(symb, false, gram.Symb2NTerminal)
		moreComp, moreActv, err := (&completed).completer(allItemSet, gram.Symb2NTerminal)
		if err != nil {
			return nil, err
		}
		active = append(active, moreActv...)
		predict, err := (&active).predictor(idx + 1, gram.Symb2NTerminal)
		if err != nil {
			return nil, err
		}
		allItemSet = append(allItemSet, itemSet{ active, predict })
		compArr = append(compArr, completed)
	}

	res := identifyStartSymbolExpr(&compArr[len(compArr)-1], gram.StartSymbol)
	return res, nil
}

func identifyStartSymbolExpr(subset *subSet, startSymb string) []string {
	res := []string{}
	for _, it := range *subset {
		if it.nonTerm == startSymb {
			res = append(res, it.nonTerm)
		}
	}
	return res
}