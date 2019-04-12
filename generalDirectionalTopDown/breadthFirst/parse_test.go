package breadthFirst

import (
	"testing"
	"fmt"
	"container/list"
	"github.com/maixianyu/parsing-method/common"
)

func listComp(t *testing.T, expect [][]string, l *list.List) {
	for i, e := 0, l.Front(); e != nil; i, e = i+1, e.Next() {
		stack := e.Value.([]string)
		if len(expect[i]) != len(stack) {
			t.Fatalf("expect[i] is %v, stack is %v, len unequal! i=%d", expect[i], stack, i)
		}
		for j := range expect[i] {
			if stack[j] != expect[i][j] {
				t.Errorf("symb want:%v, got:%v", expect[i][j], stack[j])
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
	for predict(analysis, prediction, gram.Symb2NTerminal) == false {}

	// match
	t.Log("begin match")
	fmt.Println("begin match")
	matched, rest := []string{}, []string{"a", "a", "b", "c"}
	for match(&matched, &rest, analysis, prediction, gram.Symb2NTerminal) == false { }

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

func TestParse(t *testing.T) {
	fpath := "../../samples/GreibachNormalForm"
	gram := common.ReadGrammarFromFile(fpath)
	input := "a a b c"
	res, err := parse(gram, input)
	if err != nil {
		t.Log(err)
	}
	expectStr := []string{"S", "A", "a", "A", "a", "B", "b", "c", "#"}
	if common.StringSliceEqual(res, expectStr) == false {
		t.Errorf("rest want:%v, got:%v", expectStr, res)
	}

}