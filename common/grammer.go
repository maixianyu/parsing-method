package common

import (
	"log"
	"io/ioutil"
	"strings"
)

type Symbol string

type RightHandSide struct {
	symbols []Symbol
}

type NonTerminal struct {
	symbol Symbol
	rightHandSides []RightHandSide
}

const (
	Left int8  = iota
	Right int8 = iota
) 

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSide(line string, i int8) string {
	sides := strings.Split(line, "->")
	if len(sides) != 2 {
		log.Fatal(line + "has sides unequal to 2")
	}
	return sides[i]
}

func getSymbols(line string) []Symbol {
	subs := strings.Split(line, " ")
	if len(subs) <= 0 {
		log.Fatal("no symbol in " + line)
	}
	res := []Symbol{}
	for s := range subs {
		res = append(res, Symbol(s))
	}
	return res
}

func ReadGrammar2NonTerminals(filename string) []NonTerminal {
	var res []NonTerminal

	// read grammar file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// get each line from grammar file
	lines := strings.Split(string(content), "\n")
	
	// construct non-terminal from each line
	for _, l := range lines {
		nt := NonTerminal{}

		// get non-terminal symbol from left-side
		nt.symbol = Symbol(getSide(l, Left))

		// get right-hand side
		rightSide := getSide(l, Right)
		opts := strings.Split(rightSide, "|")
		for _, opt := range opts {
			rhSide := RightHandSide{getSymbols(opt)}
			nt.rightHandSides = append(nt.rightHandSides, rhSide)
		}
	}

	return res
}