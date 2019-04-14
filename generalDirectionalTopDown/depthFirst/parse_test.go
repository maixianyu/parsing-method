package backtracking

import (
	"testing"
	"github.com/maixianyu/parsing-method/common"
)

func TestParse(t *testing.T) {
	fpath := "../../samples/GreibachNormalForm"
	gram := common.ReadGrammarFromFile(fpath)
	input := "a a b c"
	res, err := parse(gram, input)
	if err != nil {
		t.Log(err)
	}
	expectStr := []string{"S", "A", "a", "A", "a", "B", "b", "c", "#"}
	if len(res) != 1 || common.StringSliceEqual(res[0], expectStr) == false {
		t.Errorf("rest want:%v, got:%v", expectStr, res)
	} else {
		t.Log(res)
	}
}