package cyk

import (
	"testing"
	"github.com/maixianyu/parsing-method/common"
)

func TestBuildRecogTable(t *testing.T) {
	gramfile := "../samples/number"
	gram := common.ReadGrammarFromFile(gramfile)

	input := "1.2"
	rcgTable := buildRecogTable(gram, input)
	expect := recogTable([][][]string{
		/* col 0 */
		[][]string{
			[]string{"Number", "Integer", "Digit"},
			[]string{},
			[]string{"Number", "N1"},
		},
		/* col 1 */
		[][]string{
			[]string{"T1"},
			[]string{"Fraction"},
			nil,
		},
		/* col 2 */
		[][]string{
			[]string{"Number", "Integer", "Digit"},
			nil,
			nil,
		},
	})

	if common.StringSSSEqual([][][]string(rcgTable), expect) == false {
		t.Errorf("want: %v\n got: %v\n", expect, rcgTable)
	}
}