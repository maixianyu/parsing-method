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
	StartSymbol string
	Symb2NTerminal map[string]NonTerminal
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

/* get symbols in a expression, like "a + b" -> "a", "+", "b" */
func getSymbols(line string) []string {
	subs := strings.Split(strings.TrimSpace(line), " ")
	if len(subs) <= 0 {
		log.Fatal("no symbol in " + line)
	}
	res := []string{}
	for _, s := range subs {
		res = append(res, s)
	}
	return res
}


func ReadGrammar(content string) Grammar {
	var gram Grammar

	// get each line from grammar file
	lines := strings.Split(string(content), "\n")
	gram.Symb2NTerminal = make(map[string]NonTerminal, len(lines))
	
	// construct non-terminal from each line
	for idx, l := range lines {
		nt := NonTerminal{}

		// get non-terminal symbol from left-side
		nt.Symbol = getSide(l, Left, "->")

		// identify start symbol
		if idx == 0 {
			gram.StartSymbol = nt.Symbol
		}

		// get right-hand side
		rightSide := getSide(l, Right, "->")
		opts := strings.Split(rightSide, "|")
		for _, opt := range opts {
			rhSide := getSymbols(opt)
			nt.RHSides = append(nt.RHSides, RightHandSide(rhSide))
		}

		// construct map
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