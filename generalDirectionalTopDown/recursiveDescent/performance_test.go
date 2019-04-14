package recursiveDescent

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	input := "a a b c"
	for i := 0; i < b.N; i++ {
		_ = parse(input)
	}
}

func TestPerformance(t *testing.T) {
	res := testing.Benchmark(BenchmarkParse)
	t.Log(res)
}