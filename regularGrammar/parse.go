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
			for ; natom > 1; natom-- {
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
			for ; natom > 1; natom-- {
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

	if len(p) == 0 {
		return nil, errors.New("len(p)==0 at end")
	}
	/* for clause, because it is going to end. And notice that natom must be 0 */
	for ; natom > 0; natom-- {
		res = append(res, '.')
	}
	for ; nalt > 0; nalt-- {
		res = append(res, '|')
	}
	return res, nil
}

