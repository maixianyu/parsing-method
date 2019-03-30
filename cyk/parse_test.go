package cyk

import (
	"testing"
	"strings"
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

func TestParse(t *testing.T) {
	gramfile := "../samples/number"
	gram := common.ReadGrammarFromFile(gramfile)

	input := "1.2"
	res, ok := parse(gram, input)
	expect := []string{
		"Number ->",
		"IntegerFraction ->",
		"1Fraction ->",
		"1T1Integer ->",
		"1.Integer ->",
		"1.2 ->",
	}
	if ok == false || common.StringSliceEqual(res, expect) == false {
		t.Errorf("got: %v", res)
	}

	input = "12.34e+56"
	res, ok = parse(gram, input)
	expect = []string{
		"Number ->",
		"N1Scale ->",
		"IntegerFractionScale ->",
		"IntegerDigitFractionScale ->",
		"1DigitFractionScale ->",
		"12FractionScale ->",
		"12T1IntegerScale ->",
		"12.IntegerScale ->",
		"12.IntegerDigitScale ->",
		"12.3DigitScale ->",
		"12.34Scale ->",
		"12.34N2Integer ->",
		"12.34T2SignInteger ->",
		"12.34eSignInteger ->",
		"12.34e+Integer ->",
		"12.34e+IntegerDigit ->",
		"12.34e+5Digit ->",
		"12.34e+56 ->",
	}
	t.Log(strings.Join(res, "\n"))
	if ok == false || common.StringSliceEqual(res, expect) == false {
		t.Errorf("got: %v", res)
	}
}