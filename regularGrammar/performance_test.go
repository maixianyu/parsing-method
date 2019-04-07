package regularGrammar

import (
	"strings"
	"testing"
)

func repeatedString(rep string, N int) string {
	res := make([]string, N, N)
	for idx := range res {
		res[idx] = rep
	}
	return strings.Join(res, "")
}

func benchParse(textSize int, parse func (regexp string, input []string) []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp := repeatedString("a?", textSize) + repeatedString("a", textSize)
		input := repeatedString("a", textSize)
		parse(regexp, []string{input})
	}
}

var textSize int
var parse func (regexp string, input []string) []string

func BenchmarkParse(b *testing.B) {
	i := textSize
	f := parse
	benchParse(i, f, b)
}

func TestPerformance(t *testing.T) {
	parsefun := [...]func (regexp string, input []string) []string{ ParseDFA, ParseNFA}
	for _, f := range parsefun {
		for i := 1; i <= 1; i++ {
			parse = f
			textSize = i * 1
			res := testing.Benchmark(BenchmarkParse)
			t.Logf("textSize=%d\n", textSize)
			t.Log(res)
		}
	}
}
