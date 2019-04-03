/*
* Regular expression implementation.
* Supports only ( | ) * + ?.  No escapes.
* Compiles to NFA and then simulates NFA
* using Thompson's algorithm.
*
* See also http://swtch.com/~rsc/regexp/ and
* Thompson, Ken.  Regular Expression Search Algorithm,
* Communications of the ACM 11(6) (June 1968), pp. 419-422.
*/

package regularGrammar

import (
	"errors"
)

/*
* Convert infix regexp re to postfix notation.
* Insert . as explicit concatenation operator.
* Similar rules in conversion can be viewd in:
* http://csis.pace.edu/~wolf/CS122/infix-postfix.htm
*/

/* a struct to record nalt and natom in a pair of parentheses */
type paren struct {
	nalt int
	natom int
}

func re2post(re []rune) ([]rune, error) {
	/* number of alternative, number of atom */
	nalt, natom := 0, 0
	res := []rune{}

	/* used to record the context */
	p := []paren{}

	for _, r := range re {
		switch r {
		case '(':
			/* if clause, not for clause. because ( is not an end symbol */
			if natom > 1 {
				natom--
				res = append(res, '.')
			}
			/* push current context */
			curP := paren {
				nalt: nalt,
				natom: natom,
			}
			p = append(p, curP)
			nalt, natom = 0, 0
		case '|':
			if natom == 0 {
				return nil, errors.New("natom == 0 befor |")
			}
			/* for clause. because | is an end symbol */
			for natom--; natom > 0; natom-- {
				res = append(res, '.')
			}
			nalt++
		case ')':
			if len(p) == 0 {
				return nil, errors.New("len(p)==0 befor )")
			}
			if natom == 0 {
				return nil, errors.New("natom == 0 before )")
			}
			/* for clause. because ) is an end symbol */
			for natom--; natom > 0; natom-- {
				res = append(res, '.')
			}
			for ; nalt > 0; nalt-- {
				res = append(res, '|')
			}
			/* recover the context befor, and pop the current paren */
			nalt, natom = p[len(p)-1].nalt, p[len(p)-1].natom
			p = p[:len(p)-1]
			natom++
		case '*', '+', '?':
			if natom == 0 {
				return nil, errors.New("natom == 0 before */+/?")
			}
			res = append(res, r)
		default:
			/* if clause, not for clause. because ( is not an end symbol */
			if natom > 1 {
				natom--
				res = append(res, '.')
			}
			res = append(res, r)
			natom++
		}
	}

	if len(p) != 0 {
		return nil, errors.New("len(p)==0 at end")
	}
	/* for clause, because it is going to end. And notice that natom must be 0 */
	for natom--; natom > 0; natom-- {
		res = append(res, '.')
	}
	for ; nalt > 0; nalt-- {
		res = append(res, '|')
	}
	return res, nil
}

/*
* Represents an NFA state plus zero or one or two arrows exiting.
* if c == Match, no arrows out; matching state.
* If c == Split, unlabeled arrows to out and out1 (if != NULL).
* If c < 256, labeled arrow with character c to out.
*/

const (
	Match int = 256
	Split int = 257
)

type State struct {
	c int
	out *State
	out1 *State
	lastlist int
}

/* matching state */
var matchState State = State{ c: Match }
var nstate int

/* Allocate and initialize State */
func state(c int, out *State, out1 *State) *State {
	return &State{
		c: c,
		out: out,
		out1: out1,
		lastlist: 0,
	}
}

/*
* A partially built NFA without the matching state filled in.
* Frag.start points at the start state.
* Frag.out is a list of places that neet to be set to be
* next state for this fragment.
*/
type Ptrlist struct {
	next *Ptrlist
	s *State
}

type Frag struct {
	start *State
	out *Ptrlist
}

/* Initialize Frag struct. */
func frag(start *State, out *Ptrlist) Frag {
	return Frag{
		start: start,
		out: out,
	}
}

/* Create singleton list containing just outp. */
func listl(outp *State) *Ptrlist {
	var l *Ptrlist
	l.s = outp
	l.next = nil
	return l
}

/* Patch the list of states at out to point to start. */
func patch(l *Ptrlist, s *State) {
	var next *Ptrlist
	for ; l != nil; l=next {
		next = l.next
		l.s = s
	}
}

/* Join the two lists l1 and l2, returning the combination. */
func append(l1 *Ptrlist, l2 *Ptrlist) *Ptrlist {
	oldl1 := l1
	for ; l1.next != nil;  {
		l1 = l1.next
	}
	l1.next = l2
	return oldl1
}

