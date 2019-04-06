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

func benchParse(textSize int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp := repeatedString("a?", textSize) + repeatedString("a", textSize)
		input := repeatedString("a", textSize)
		parse(regexp, []string{input})
	}
}

var textSize int

func BenchmarkParse(b *testing.B) {
	i := textSize
	benchParse(i, b)
}

func TestPerformance(t *testing.T) {
	for i := 1; i <= 10; i++ {
		textSize = i * 10
		res := testing.Benchmark(BenchmarkParse)
		t.Logf("textSize=%d\n", textSize)
		t.Log(res)
	}
}
