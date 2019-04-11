package breadthFirst

import (
	"testing"
	"container/list"
	"github.com/maixianyu/parsing-method/common"
)

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

	res := true
	for res != false {
		res = predict(analysis, prediction, gram.Symb2NTerminal)
	}
	if res == true {
		t.Fatalf("res want:%v, got:%v", false, res)
	}
	// compare analysis
	t.Log("compare analysis")
	expect := [][]string{
		[]string{"S", "A"},
		[]string{"S", "A"},
		[]string{"S", "D"},
		[]string{"S", "D"},
	}
	for i, e := 0, analysis.Front(); e != nil; i, e = i+1, e.Next() {
		ana := e.Value.([]string)
		for j, symb := range ana {
			if symb != expect[i][j] {
				t.Errorf("symb want:%v, got:%v", expect[i][j], symb)
			}
		}
	}
	// compare prediction
	expect = [][]string{
		[]string{"a", "B"},
		[]string{"a", "A", "B"},
		[]string{"a", "b", "C"},
		[]string{"a", "D", "b", "C"},
	}
	t.Log("compare prediction")
	for i, e := 0, prediction.Front(); e != nil; i, e = i+1, e.Next() {
		pred := e.Value.([]string)
		for j, symb := range pred {
			if symb != expect[i][j] {
				t.Errorf("symb want:%v, got:%v", expect[i][j], symb)
			}
		}
	}

}