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
	"sort"
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

/*
* Compare lists: first by length, then by members.
*/
func listcomp(l1 *List, l2 *List) int {
	len1, len2 := l1.numS, l2.numS
	if len1 > len2 {
		return 1
	}
	if len1 < len2 {
		return -1
	}

	for idx := range l1.s {
		// is there another way without fmt ?
		p1 := l1.s[idx].stateIdx
		p2 := l2.s[idx].stateIdx
		if p1 < p2 {
			return -1
		} else if p1 > p2 {
			return 1
		}
	}
	return 0
}

var allDState *DState
/* Return the cached DState for list l, creating a new one if needed. */
func dstate(l *List) *DState {
	sort.SliceStable(l.s,
		func(i, j int) bool {
			return l.s[i].stateIdx < l.s[j].stateIdx
		})

	dp := &allDState
	d := *dp
	for ;d != nil; d = *dp {
		i := listcomp(l, &d.l)
		if i < 0 {
			dp = &d.left
		} else if i > 0 {
			dp = &d.right
		} else {
			return d
		}
	}
	d = new(DState)
	d.l.s = make([]*State, len(l.s))
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