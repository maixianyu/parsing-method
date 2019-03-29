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
	/* difference between NoERule and HasERule lies in if epsilon exists in input */

	/* NoERule */
	/*read from string*/
	expect := Grammar{
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
		RhSide2NTSymb: map[string][]string{
			"Expr + Term": []string{
				"Expr",
			},
			"Term": []string{
				"Expr",
			},
			"Term x Factor": []string{
				"Term",
			},
			"Factor": []string{
				"Term",
			},
			"( Expr )": []string{
				"Factor",
			},
			"i": []string{
				"Factor",
			},
		},
	}
	input := "Expr -> Expr + Term | Term\nTerm -> Term x Factor | Factor\nFactor -> ( Expr ) | i"
	res := ReadGrammar(input)
	eq := reflect.DeepEqual(expect, res)
	if eq == false {
		t.Errorf("want: %v\nget: %v",
			expect, res)
	}

	/* read from file */
	input = "../samples/arithmetic"
	res = ReadGrammarFromFile(input)
	eq = reflect.DeepEqual(expect, res)
	if eq == false {
		t.Errorf("want: %v\nget: %v",
			expect, res)
	}


	/* HasERule */
	/*read from string*/
	expect = Grammar{
		StartSymbol: "S",
		Symb2NTerminal: map[string]NonTerminal{
			"S": NonTerminal{
				Symbol: "S",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"L",
							"S",
							"D",
						}),
						RightHandSide([]string{
							"",
						}),
					},
				},
			"L": NonTerminal{
				Symbol: "L",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"",
						}),
					},
				},
			"D": NonTerminal{
				Symbol: "D",
				RHSides: []RightHandSide{
						RightHandSide([]string{
							"d",
						}),
					},
				},
			},
		RhSide2NTSymb: map[string][]string{
			"L S D": []string{
				"S",
			},
			"": []string{
				"S", "L",
			},
			"d": []string{
				"D",
			},
		},
	}
	input = "S -> L S D | epsilon\nL -> epsilon\nD -> d"
	res = ReadGrammar(input)
	eq = reflect.DeepEqual(expect, res)
	if eq == false {
		t.Errorf("want: %v\nget: %v",
			expect, res)
	}
}