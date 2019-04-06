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
	"fmt"
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
	if len(l1.s) < len(l2.s) {
		return -1
	}

	if len(l1.s) > len(l2.s) {
		return 1
	}

	for idx := range l1.s {
		// is there another way without fmt ?
		p1 := fmt.Sprintf("%p", l1.s[idx])
		p2 := fmt.Sprintf("%p", l2.s[idx])
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
			return fmt.Sprintf("%p", l.s[i]) < fmt.Sprintf("%p", l.s[j])
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