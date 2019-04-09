package regularGrammar

import (
	"strings"
	"testing"
	"github.com/maixianyu/parsing-method/common"
)

func repeatedString(rep string, N int) string {
	res := make([]string, N, N)
	for idx := range res {
		res[idx] = rep
	}
	return strings.Join(res, "")
}

func testcase1(textSize int, parse func (regexp string, input []string) []string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp := repeatedString("a?", textSize) + repeatedString("a", textSize)
		input := repeatedString("a", textSize)
		parse(regexp, []string{input})
	}
}

func testcase2(textSize int, parse func (regexp string, input []string) []string, b *testing.B) {
	alphabet := make([]string, 26, 26)
	for i := 0; i < 26; i++ {
		alphabet[i] = string(int('a') + i)
	}
	for i := 0; i < b.N; i++ {
		regexp := "0(" + strings.Join(alphabet, "|") + ")*1" 
		subinput := []string{}
		for j := 1; j <= textSize; j++ {
			/* be careful, here is 4, which makes DFA better than NFA.
			*  If it is 26, DFA performance will be equal with that of NFA.
			*/
			subinput = append(subinput, alphabet[j % 4])
		}
		input := "0" + strings.Join(subinput, "") + "1"
		parse(regexp, []string{input})
	}
}


var textSize int
var parse func (regexp string, input []string) []string

func BenchmarkParse(b *testing.B) {
	i := textSize
	f := parse
	/* choose a testcase to run benchmark */
	//testcase1(i, f, b)
	testcase2(i, f, b)
}

func TestPerformance(t *testing.T) {
	parsefun := [...]func (regexp string, input []string) []string{ ParseDFA, ParseNFA}
	for _, f := range parsefun {
		t.Log(common.GetFunctionName(f))
		for i := 1; i <= 5; i++ {
			parse = f
			textSize = i * 10
			res := testing.Benchmark(BenchmarkParse)
			t.Logf("textSize=%d\n", textSize)
			t.Log(res)
		}
	}
}
