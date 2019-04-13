package backtracking

import (
	"testing"
	"github.com/maixianyu/parsing-method/common"
)

func BenchmarkParse(b *testing.B) {
	fpath := "../../samples/GreibachNormalForm"
	gram := common.ReadGrammarFromFile(fpath)
	input := "a a b c"
	for i := 0; i < b.N; i++ {
		res, err := parse(gram, input)
		if err != nil {
			b.Log(err)
			b.Log(res)
		}
	}
}

func TestPerformance(t *testing.T) {
	res := testing.Benchmark(BenchmarkParse)
	t.Log(res)
}