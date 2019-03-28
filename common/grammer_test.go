package common

import (
	"testing"
	"reflect"
)

func TestGetSymbols(t *testing.T) {
	initKeyChar()

	sample := " a b c d "
	expect := []string{"a", "b", "c", "d"}
	res := getSymbols(sample)
	if StringSliceEqual(res, expect) == false {
		t.Errorf("want: %v, len(%d), get: %v, len(%d)\n",
			expect, len(expect), res, len(res))
	}

	sample = " a epsilon c epsilon "
	expect = []string{"a", "", "c", ""}
	res = getSymbols(sample)
	if StringSliceEqual(res, expect) == false {
		t.Errorf("want: %v, len(%d), get: %v, len(%d)\n",
			expect, len(expect), res, len(res))
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
	/*read from string*/
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
						RightHandSide([]string{
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
							"Factor",
						}),
						RightHandSide([]string{
							"Factor",
						}),
					},
				},
			"Factor": NonTerminal{
				Symbol: "Factor",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"(",
							"Expr",
							")",
						}),
						RightHandSide([]string{
							"i",
						}),
					},
				},
			},
		}
	input1 := "Expr -> Expr + Term | Term\nTerm -> Term x Factor | Factor\nFactor -> ( Expr ) | i"
	res1 := ReadGrammar(input1)
	eq := reflect.DeepEqual(expect1, res1)
	if eq == false {
		t.Errorf("want: %v\nget: %v",
			expect1, res1)
	}

	/* read from file */
	input2 := "../samples/arithmetic"
	res2 := ReadGrammarFromFile(input2)
	expect2 := expect1
	eq = reflect.DeepEqual(expect2, res2)
	if eq == false {
		t.Errorf("want: %v\nget: %v",
			expect2, res2)
	}
}