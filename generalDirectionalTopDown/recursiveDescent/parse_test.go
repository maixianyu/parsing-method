package recursiveDescent

import(
	"testing"
	"github.com/maixianyu/parsing-method/common"
)

func TestParse(t *testing.T) {
	// a a b c
	input := "a a b c"
	res := parse(input)
	expectStr := []string{"S -> AB", "A -> aA", "A -> a", "B -> bc"}
	if len(res) != 1 || common.StringSliceEqual(res[0], expectStr) == false {
		t.Error(res)
	}
	t.Log(input, res)

	// a b c c
	input = "a b c c"
	res = parse(input)
	expectStr = []string{"S -> DC", "D -> ab", "C -> cC", "C -> c"}
	if len(res) != 1 || common.StringSliceEqual(res[0], expectStr) == false {
		t.Error(res)
	}
	t.Log(input, res)
}