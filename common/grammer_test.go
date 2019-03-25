package common

import (
	"testing"
	"reflect"
)

func TestGetSymbols(t *testing.T) {
	sp1 := " a b c d "
	expect1 := []string{"a", "b", "c", "d"}
	res1 := getSymbols(sp1)
	if StringSliceEqual(res1, expect1) == false {
		t.Errorf("want: %v, len(%d), get: %v, len(%d)\n",
			expect1, len(expect1), res1, len(res1))
	}
}

func TestGetSide(t *testing.T) {
	sp1 := "A -> B + C"
	expect1 := "A"
	res1 := getSide(sp1, Left, "->")
	if res1 != expect1 {
		t.Errorf("want: %v, get: %v",
			expect1, res1)
	}

	expect2 := "B + C"
	res2 := getSide(sp1, Right, "->")
	if res2 != expect2 {
		t.Errorf("want: %v, get: %v",
			expect2, res2)
	}
}

func TestReadGrammar(t *testing.T) {
	expect1 := Grammar{
		StartSymbol: "Expr",
		Symb2NTerminal: map[string]NonTerminal{
			"Expr": NonTerminal{
				Symbol: "Expr",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"Expr",
							"+",
							"Term",
						}),
					},
				},
			"Term": NonTerminal{
				Symbol: "Term",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"Term",
							"x",
							"i",
						}),
						RightHandSide([]string{
							"i",
						}),
					},
				},
			},
		}
	input1 := "Expr -> Expr + Term\nTerm -> Term x i | i"
	res1 := ReadGrammar(input1)
	eq := reflect.DeepEqual(expect1, res1)
	if eq == false {
		t.Errorf("want: %v, get: %v",
			expect1, res1)
	}
}