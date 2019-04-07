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
	"log"
)

/*
* Convert infix regexp re to postfix notation.
* Insert . as explicit concatenation operator.
* Similar rules in conversion can be viewd in:
* http://csis.pace.edu/~wolf/CS122/infix-postfix.htm
 */

/* a struct to record nalt and natom in a pair of parentheses */
type paren struct {
	nalt  int
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
			curP := paren{
				nalt:  nalt,
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
	c        int
	out      *State
	out1     *State
	lastlist int
	stateIdx int
}

/* matching state */
var matchState State = State{c: Match}
var nstate int = 0

/* Allocate and initialize State */
func state(c int, out *State, out1 *State) *State {
	nstate++
	return &State{
		c:        c,
		out:      out,
		out1:     out1,
		lastlist: 0,
		stateIdx: nstate,
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
	s    **State
}

type Frag struct {
	start *State
	out   *Ptrlist
}

/* Initialize Frag struct. */
func frag(start *State, out *Ptrlist) Frag {
	return Frag{
		start: start,
		out:   out,
	}
}

/* Create singleton list containing just outp. */
func listl(outp **State) *Ptrlist {
	if outp == nil {
		log.Fatal("outp cannot be nil.")
	}
	return &Ptrlist{
		next: nil,
		s:    outp,
	}
}

/* Patch the list of states at out to point to start. */
func patch(l *Ptrlist, s *State) {
	var next *Ptrlist
	for ; l != nil; l = next {
		next = l.next
		*l.s = s
	}
}

/* Join the two lists l1 and l2, returning the combination. */
func appendList(l1 *Ptrlist, l2 *Ptrlist) *Ptrlist {
	oldl1 := l1
	for l1.next != nil {
		l1 = l1.next
	}
	l1.next = l2
	return oldl1
}

/*
* Convert postfix regular expression to NFA.
* Return start state.
 */
func post2nfa(postfix []rune) (*State, error) {
	stack := []Frag{}
	var e1, e2, e Frag
	var s *State

	for _, p := range postfix {
		switch p {
		case '.': /* catenate */
			/* pop stack */
			e2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			e1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			patch(e1.out, e2.start)
			/* push stack */
			stack = append(stack, Frag{e1.start, e2.out})
		case '|': /* alternate */
			e2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			e1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = state(Split, e1.start, e2.start)
			stack = append(stack, Frag{s, appendList(e1.out, e2.out)})
		case '?': /* zero or one */
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = state(Split, e.start, nil)
			stack = append(stack, Frag{s, appendList(e.out, listl(&s.out1))})
		case '*':
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = state(Split, e.start, nil)
			patch(e.out, s)
			stack = append(stack, Frag{s, listl(&s.out1)})
		case '+':
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			s = state(Split, e.start, nil)
			patch(e.out, s)
			stack = append(stack, Frag{e.start, listl(&s.out1)})
		default:
			s = state(int(p), nil, nil)
			stack = append(stack, Frag{s, listl(&s.out)})
		}
	}

	e, stack = stack[len(stack)-1], stack[:len(stack)-1]
	if len(stack) != 0 {
		return nil, errors.New("len(stack) != 0 at the end.")
	}
	patch(e.out, &matchState)
	return e.start, nil
}

type List struct {
	s []*State
	numS int
}

var listid int = 0

/* Add s to 1, following unlabeled arrows. */
func addstate(l *List, s *State) {
	if s == nil || s.lastlist == listid {
		return
	}
	s.lastlist = listid
	if s.c == Split {
		addstate(l, s.out)
		addstate(l, s.out1)
		return
	}
	l.numS++
	l.s = append(l.s, s)
}

/* Compute initial state list */
func startlist(start *State, l *List) *List {
	listid++
	l.s = l.s[0:0]
	addstate(l, start)
	return l
}

/* Check whether state list contains a match */
func ismatch(l *List) bool {
	for _, s := range l.s {
		if s == &matchState {
			return true
		}
	}
	return false
}

/*
* Step the NFA from the states in clist
* past the character c,
* to create next NFA state set nlist.
 */
func step(clist *List, c int, nlist *List) {
	listid++
	/* clear nlist */
	nlist.s = nlist.s[0:0]
	/* transfer state */
	for _, s := range clist.s {
		if s.c == c {
			/* out is for c matched, out1 is for pass */
			addstate(nlist, s.out)
		}
	}
}

var l1, l2 List = List{[]*State{}, 0}, List{[]*State{}, 0}

/* Run NFA to determine whether it matches s. */
func match(start *State, input []rune) bool {
	clist := startlist(start, &l1)
	nlist := &l2
	for _, c := range input {
		step(clist, int(c), nlist)
		clist, nlist = nlist, clist
	}
	return ismatch(clist)
}

/* parse string in NFA style */
func ParseNFA(regexp string, input []string) []string {
	res := []string{}
	post, err := re2post([]rune(regexp))
	if err != nil {
		log.Fatal(err)
	}
	start, err := post2nfa(post)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range input {
		if match(start, []rune(s)) == true {
			res = append(res, s)
		}
	}
	return res
}
