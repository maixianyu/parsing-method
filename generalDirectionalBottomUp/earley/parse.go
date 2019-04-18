package earley

import(
	"errors"
	"log"
	"fmt"
	"strconv"
	"strings"

	"github.com/maixianyu/parsing-method/common"
)

type item struct {
	dotPos int
	// at starting from 0
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
		sub[i].at = 0
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
					it.dotPos += 1
					completed = append(completed, it)
				/* curSymb is not at tail, then recognized as active */
				} else {
					it.dotPos += 1
					active = append(active, it)
				}
			}
		}
	}
	return
}

/* completer inspects completed set, which contains the items that have just been recognized
and can now be reduced.
*/
func (completed *subSet) completer(allItemSet []itemSet, symb2nt map[string]common.NonTerminal) (newComp subSet, newActv subSet, err error) {
	newComp, newActv = []item{}, []item{}
	if len(*completed) == 0 {
		return 
	}

	// a loop for current items in completed
	for _, it := range *completed {
		idxSet := it.at
		if idxSet >= len(allItemSet) {
			info := fmt.Sprintf("idxSet=%d >= len(allItemSet)=%d", idxSet, len(allItemSet))
			return nil, nil, errors.New(info)
		}
		searchSet := allItemSet[idxSet]
		c, a := searchSet.scanner(it.nonTerm, true, symb2nt)
		newComp = append(newComp, c...)
		newActv = append(newActv, a...)
	}

	// a recursive for new items added before
	recurComp, recurActv, err := newComp.completer(allItemSet, symb2nt)
	if err != nil {
		return
	}
	newComp = append(newComp, recurComp...)
	newActv = append(newActv, recurActv...)
	
	return
}

/* parse */
func parse(gram common.Grammar, input string) ([]string, error) {
	in := common.SliceInput(input, " ")
	// make itemSet0
	actv0 := initActiveSet(gram)
	pred0, err := (&actv0).predictor(0, gram.Symb2NTerminal)
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
		completed = append(completed, moreComp...)
		compArr = append(compArr, completed)
		curAt := idx
		predict, err := (&active).predictor(curAt + 1, gram.Symb2NTerminal)
		if err != nil {
			return nil, err
		}
		allItemSet = append(allItemSet, itemSet{ active, predict })
	}

	// make sure the startSymb is in the last completed array
	startSymb, err := identifyStartSymbolExpr(&compArr[len(compArr)-1], gram.StartSymbol)
	if err != nil {
		return nil, err
	}

	// construct parse tree in completed array
	res, err := constructTree(startSymb, compArr, gram.Symb2NTerminal)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/* search for item at specific completed subset */
func searchItem(symb string, at int, completedArr []subSet) (item, error) {
	if at >= len(completedArr) {
		return item{}, errors.New("at=" + strconv.Itoa(at) + " >= len(completedArr)=" + strconv.Itoa(len(completedArr)))
	}
	compSet := completedArr[at]
	for i := range compSet {
		if compSet[i].nonTerm == symb {
			return compSet[i], nil
		}
	}
	return item{}, errors.New("searchItem fails in " + symb + ",at=" + string(at))
}

/* a depth-first style function to combine item */
func constructHelper(curItem item, curAt int, completedArr []subSet, symb2nt map[string]common.NonTerminal) []string {
	// conditon to terminate recursive
	if _, found := symb2nt[curItem.rhSide[0]]; len(curItem.rhSide) == 1 && found == false {
		return []string{curItem.rhSide[0]}
	}

	// loop symbols in the curItem backward
	rhSide := curItem.rhSide
	res := []string{ strings.Join(rhSide, "") }
	branchRes := make([][]string, len(rhSide))
	for idx := len(rhSide)-1; idx >= 0; idx-- {
		symb := rhSide[idx]
		at := curAt - (len(rhSide)-1-idx)
		// if symb is a terminal
		if _, found := symb2nt[symb]; found == false {
			branchRes[idx] = []string{symb}
		
		// if symb is a non-terminal
		} else {
			it, err := searchItem(symb, at, completedArr)
			//fmt.Println(it)
			if err != nil {
				log.Fatal(err)
			}
			branchRes[idx] = constructHelper(it, at, completedArr, symb2nt)
		}
	}
	res = append(res, common.CombTraceWithTemplate(branchRes, rhSide)...)
	return res
}

/* construct parse tree from completed sets*/
func constructTree(startSymb string, completedArr []subSet, symb2nt map[string]common.NonTerminal) ([]string, error) {
	res := make([]string, 0)
	for _, it := range completedArr[len(completedArr)-1] {
		if it.nonTerm == startSymb {
			res = append(res, constructHelper(it, len(completedArr)-1, completedArr, symb2nt)...)
			return res, nil
		}
	}
	return nil, errors.New("Fail to find startSymbol in the last completed set.")
}

func identifyStartSymbolExpr(subset *subSet, startSymb string) (string, error) {
	for _, it := range *subset {
		if it.nonTerm == startSymb {
			return startSymb, nil
		}
	}
	return "", errors.New("Fail to parse input.")
}