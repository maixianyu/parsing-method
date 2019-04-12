package breadthFirst

import (
	"testing"
	"fmt"
	"container/list"
	"github.com/maixianyu/parsing-method/common"
)

func listComp(t *testing.T, expect [][]string, l *list.List) {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		ana := e.Value.([]string)
		for j, symb := range ana {
			if symb != expect[i][j] {
				t.Errorf("symb want:%v, got:%v", expect[i][j], symb)
			}
		}
	}
}

func TestPredict(t *testing.T) {
	fpath := "../../samples/GreibachNormalForm"
	gram := common.ReadGrammarFromFile(fpath)
	// prediction
	preStack := []string{gram.StartSymbol}
	prediction := list.New()
	prediction.PushFront(preStack)
	// analysis
	anaStack := []string{}
	analysis := list.New()
	analysis.PushFront(anaStack)

	res := false
	for res != true {
		res = predict(analysis, prediction, gram.Symb2NTerminal)
	}
	if res == false {
		t.Fatalf("res want:%v, got:%v", true, res)
	}
	// compare analysis
	t.Log("compare analysis")
	expect := [][]string{
		[]string{"S", "A"},
		[]string{"S", "A"},
		[]string{"S", "D"},
		[]string{"S", "D"},
	}
	listComp(t, expect, analysis)

	// compare prediction
	t.Log("compare prediction")
	expect = [][]string{
		[]string{"a", "B"},
		[]string{"a", "A", "B"},
		[]string{"a", "b", "C"},
		[]string{"a", "D", "b", "C"},
	}
	listComp(t, expect, prediction)
}


func TestMatch(t *testing.T) {
	fpath := "../../samples/GreibachNormalForm"
	gram := common.ReadGrammarFromFile(fpath)
	// prediction
	preStack := []string{gram.StartSymbol}
	prediction := list.New()
	prediction.PushFront(preStack)
	// analysis
	anaStack := []string{}
	analysis := list.New()
	analysis.PushFront(anaStack)

	// predict first
	t.Log("begin predict")
	res := false
	for res != true {
		res = predict(analysis, prediction, gram.Symb2NTerminal)
	}
	if res == false {
		t.Fatalf("res want:%v, got:%v", true, res)
	}

	// match
	t.Log("begin match")
	fmt.Println("begin match")
	matched, rest := []string{}, []string{"a", "a", "b", "c"}
	res = false
	for res != true {
		res = match(&matched, &rest, analysis, prediction, gram.Symb2NTerminal)
	}

	fmt.Println("begin compare")
	// compare matched
	expectStr := []string{"a"}
	if common.StringSliceEqual(matched, expectStr) == false {
		t.Errorf("matched want:%v, got:%v", expectStr, matched)
	}

	// compare rest
	expectStr = []string{"a", "b", "c"}
	if common.StringSliceEqual(rest, expectStr) == false {
		t.Errorf("rest want:%v, got:%v", expectStr, rest)
	}

	// compare analysis
	t.Log("compare analysis")
	expect := [][]string{
		[]string{"S", "A", "a"},
		[]string{"S", "A", "a"},
		[]string{"S", "D", "a"},
		[]string{"S", "D", "a"},
	}
	listComp(t, expect, analysis)

	// compare prediction
	t.Log("compare prediction")
	expect = [][]string{
		[]string{"B"},
		[]string{"A", "B"},
		[]string{"b", "C"},
		[]string{"D", "b", "C"},
	}
	listComp(t, expect, prediction)
}