/*
* Regular expression implementation.
* Supports only ( | ) * + ?.  No escapes.
* Compiles to NFA and then simulates NFA
* using Thompson's algorithm.
*
* Caching the NFA to build a DFA 
*
* See also http://swtch.com/~rsc/regexp/ and
* Thompson, Ken.  Regular Expression Search Algorithm,
* Communications of the ACM 11(6) (June 1968), pp. 419-422.
 */

package regularGrammar

import (
	"log"
)

/*
* Represent a DFA state: a cached NFA state list.
*/
type DState struct {
	l List
	next [256]*DState
	left *DState
	right *DState
}

func listcomp(l1 *List, l2 *List) bool {
	equal := true
	for idx := range l1.s {
		equal = equal && (l1.s[idx].stateIdx == l2.s[idx].stateIdx)
	}
	return equal
}

func fingerPrint(l *List) int {
	res := len(l.s)
	for idx := range l.s {
		res += l.s[idx].stateIdx
		return res
	}
	return res
}

/* Return the cached DState for list l, creating a new one if needed. */
var allDState *DState
func dstate(l *List) *DState {
	dp := &allDState
	finIn := fingerPrint(l)
	for *dp != nil {
		finDP := fingerPrint(&(*dp).l)
		if finIn < finDP {
			dp = &(*dp).left
		} else if finIn > finDP {
			dp = &(*dp).right
		} else {
			if (listcomp(l, &(*dp).l)) {
				return *dp
			} else {
				d := new(DState)
				d.l.s = make([]*State, len(l.s), len(l.s))
				copy(d.l.s, l.s)
				*dp = d
				return d
			}
		}
	}
	d := new(DState)
	d.l.s = make([]*State, len(l.s), len(l.s))
	copy(d.l.s, l.s)
	*dp = d
	return d
}

func startDState(start *State) *DState {
	return dstate(startlist(start, &l1))
}

func nextDState(d *DState, c int) *DState {
	step(&d.l, c, &l1)
	d.next[c] = dstate(&l1)
	return d.next[c]
}

func matchDState(start *DState, input []rune) bool {
	d := start
	for _, c := range input {
		next := d.next[int(c)]
		if next == nil {
			next = nextDState(d, int(c))
		}
		d = next
	}
	return ismatch(&d.l)
}

/* parse string in DFA style */
func ParseDFA(regexp string, input []string) []string {
	res := []string{}
	/* re2post */
	post, err := re2post([]rune(regexp))
	if err != nil {
		log.Fatal(err)
	}
	/* post2nfa */
	nstate = 0
	allDState = nil
	start, err := post2nfa(post)
	if err != nil {
		log.Fatal(err)
	}
	/* parse */
	for _, s := range input {
		if matchDState(startDState(start), []rune(s)) == true {
			res = append(res, s)
		}
	}
	return res
}