package common

import (
	"log"
	"io/ioutil"
	"strings"
)

type RightHandSide []string

type NonTerminal struct {
	Symbol string
	RHSides []RightHandSide
}

type Grammar struct {
	/*
	Expr -> Expr + Term | Term
	Term -> Term x Factor | Factor
	Factor -> ( Expr ) | i
	*/

	/* first symble on left-hand side of the first-line expression */
	/* example: Expr is the start symbol */
	StartSymbol string

	/* symbol -> non-terminal struct */
	/* example: Expr, Term and Factor are non-terminal symbol */
	Symb2NTerminal map[string]NonTerminal

	/* right-hand side string -> non-termial symbol slice*/
	/* example: Term -> Expr, Factor -> Term, i -> Factor */
	RhSide2NTSymb map[string][]string
}

type Side uint8

const (
	Left Side  = iota
	Right Side = iota
) 

type ERule uint8
const (
	HasERule ERule = iota
	NoERule ERule = iota
) 

var keyChar map[string]string

func getSide(line string, i Side, sep string) string {
	if i > 1 {
		log.Fatal("i > 1")
	}
	sides := strings.Split(line, sep)
	if len(sides) != 2 {
		log.Fatal(line + "has sides unequal to 2")
	}
	return strings.TrimSpace(sides[i])
}

func initKeyChar() {
	keyChar = map[string]string{}
	keyChar["epsilon"] = ""
}

func transKeyChar(input string) string {
	ch, found := keyChar[input]
	if found {
		return ch
	} else {
		return input
	}
}

/* get symbols in a expression, like "a + b" -> "a", "+", "b" */
func getSymbols(line string) []string {
	subs := strings.Split(strings.TrimSpace(line), " ")
	if len(subs) <= 0 {
		log.Fatal("no symbol in " + line)
	}
	res := []string{}
	for _, s := range subs {
		res = append(res, transKeyChar(s))
	}
	return res
}



func ReadGrammar(content string) Grammar {
	var gram Grammar

	/* init keyChar map*/
	initKeyChar()

	// get each line from grammar file
	lines := strings.Split(string(content), "\n")
	gram.Symb2NTerminal = make(map[string]NonTerminal, len(lines))
	
	/* process each line */ 
	for idx, l := range lines {
		nt := NonTerminal{}

		/* left-hand side: get non-terminal symbol */
		nt.Symbol = getSide(l, Left, "->")
		if idx == 0 {
			gram.StartSymbol = nt.Symbol
		}

		/* right-hand side:  */
		rightSide := getSide(l, Right, "->")
		opts := strings.Split(rightSide, "|")
		for _, opt := range opts {
			rhSide := getSymbols(opt)

			/* construct RhSide2NTSymb map */
			strRhSide := strings.Join(rhSide, "")
			ntSymb, found := gram.RhSide2NTSymb[strRhSide]
			if found {
				ntSymb = append(ntSymb, strRhSide)
			} else {
				ntSymb = []string{strRhSide}
			}

			/* gether rhSides */
			nt.RHSides = append(nt.RHSides, RightHandSide(rhSide))
		}
		/* construct Symb2NTerminal map */
		gram.Symb2NTerminal[nt.Symbol] = nt
	}

	return gram
}

func ReadGrammarFromFile(filename string) Grammar {
	// read grammar file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return ReadGrammar(string(content))
}