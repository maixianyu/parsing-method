package unger

import (
	"testing"
	"strings"
	"github.com/maixianyu/parsing-method/common"
)

func TestGrepWithTerminal(t *testing.T) {
	gramfile := "../samples/arithmetic"
	gram := common.ReadGrammarFromFile(gramfile)

	/* teminal in symbols*/
	symbols := []string{"Term", "x", "Factor"}
	partitions, _ := common.GeneratePartitions("(ixi)", len(symbols), common.NoERule)
	expect := [][]string{
		[]string{"(i", "x", "i)"},
	}
	res := grepWithTerminal(gram, partitions, symbols)

	if common.StringSSEqual(expect, res) == false {
		t.Errorf("want: %v, got: %v", expect, res)
	}


	/* no teminal in symbols*/
	symbols = []string{"Term"}
	partitions, _ = common.GeneratePartitions("(ixi)", len(symbols), common.NoERule)
	expect = partitions
	res = grepWithTerminal(gram, partitions, symbols)

	if common.StringSSEqual(expect, res) == false {
		t.Errorf("want: %v, got: %v", expect, res)
	}

}

func TestParseNoERule(t *testing.T) {
	gramfile := "../samples/arithmetic"
	gram := common.ReadGrammarFromFile(gramfile)

	input := "i+i"
	expect := []string{
		"Expr ->",
		"Expr+Term ->",
		"Term+Term ->",
		"Factor+Term ->",
		"i+Term ->",
		"i+Factor ->",
		"i+i ->",
	}
	rule := common.NoERule
	res, ok := parse(gram, input, rule)
	t.Log(strings.Join(res, "\n"))
	if ok != true || common.StringSliceEqual(res, expect) == false {
		t.Errorf("want: %v\nget: %v", expect, res)
	}


	input = "(i+i)xi"
	expect = []string{
		"Expr ->",
		"Term ->",
		"TermxFactor ->",
		"FactorxFactor ->",
		"(Expr)xFactor ->",
		"(Expr+Term)xFactor ->",
		"(Term+Term)xFactor ->",
		"(Factor+Term)xFactor ->",
		"(i+Term)xFactor ->",
		"(i+Factor)xFactor ->",
		"(i+i)xFactor ->",
		"(i+i)xi ->",
	}
	res, ok = parse(gram, input, rule)
	t.Log(strings.Join(res, "\n"))
	if ok != true || common.StringSliceEqual(res, expect) == false {
		t.Errorf("want: %v\nget: %v", expect, res)
	}
}


func TestParseHasERule(t *testing.T) {
	gramfile := "../samples/LSD"
	gram := common.ReadGrammarFromFile(gramfile)
	rule := common.HasERule
	t.Log(gram)

	input := "d"
	expect := []string{
		"S ->",
		"LSD ->",
		"SD ->",
		"D ->",
		"d ->",
	}
	res, ok := parse(gram, input, rule)
	t.Log(strings.Join(res, "\n"))
	if ok != true || common.StringSliceEqual(res, expect) == false {
		t.Errorf("want: %v\nget: %v", expect, res)
	}
}
