package common

import (
	"log"
	"io/ioutil"
	"strings"
)

type RightHandSide struct {
	Symbols []string
}


type NonTerminal struct {
	Symbol string
	RHSides []RightHandSide
}

type Grammar struct {
	StartSymbol string
	Symb2NTerminal map[string]NonTerminal
}

const (
	Left uint8  = iota
	Right uint8 = iota
) 

func getSide(line string, i uint8) string {
	if i > 1 {
		log.Fatal("i > 1")
	}
	sides := strings.Split(line, "->")
	if len(sides) != 2 {
		log.Fatal(line + "has sides unequal to 2")
	}
	return strings.TrimRight(strings.TrimLeft(sides[i], " "), " ")
}

func getSymbols(line string) []string {
	subs := strings.Split(line, " ")
	if len(subs) <= 0 {
		log.Fatal("no symbol in " + line)
	}
	res := []string{}
	for _, s := range subs {
		res = append(res, strings.TrimSpace(s))
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
		nt.Symbol = getSide(l, Left)

		// identify start symbol
		if idx == 0 {
			gram.StartSymbol = nt.Symbol
		}

		// get right-hand side
		rightSide := getSide(l, Right)
		opts := strings.Split(rightSide, "|")
		for _, opt := range opts {
			rhSide := RightHandSide{getSymbols(opt)}
			nt.RHSides = append(nt.RHSides, rhSide)
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