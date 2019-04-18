package earley

import(
	"testing"
	"reflect"

	"github.com/maixianyu/parsing-method/common"
)

func TestPredicorScannerCompleter(t *testing.T) {
	gramfile := "../../samples/grammar4earley"
	gram := common.ReadGrammarFromFile(gramfile)

	// initActiveSet
	actv0 := initActiveSet(gram)
	expect := subSet{
		item{
			dotPos: 0,
			at: 1,
			nonTerm: "S",
			rhSide: []string{ "E" },
		},
	}
	if reflect.DeepEqual(actv0, expect) == false {
		t.Fatalf("actv0 want:%v, got:%v\n", expect, actv0)
	}

	// predictor
	pred0, err := actv0.predictor(1, gram.Symb2NTerminal)
	expect = subSet{
		item{
			dotPos: 0,
			at: 1,
			nonTerm: "E",
			rhSide: []string{ "E", "Q", "F" },
		},
		item{
			dotPos: 0,
			at: 1,
			nonTerm: "E",
			rhSide: []string{ "F" },
		},
		item{
			dotPos: 0,
			at: 1,
			nonTerm: "F",
			rhSide: []string{ "a" },
		},
	}
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(pred0, expect) == false {
		t.Fatalf("pred0 want:%v, got:%v\n", expect, pred0)
	}

	// scanner
	iSet0 := itemSet{ actv0, pred0 }
	allSet := []itemSet{iSet0}
	compt1, actv1 := iSet0.scanner("a", false, gram.Symb2NTerminal)
	expect = subSet{
		item{
			dotPos: 1,
			at: 1,
			nonTerm: "F",
			rhSide: []string{ "a" },
		},
	}
	if reflect.DeepEqual(compt1, expect) == false {
		t.Fatalf("compt1 want:%v, got:%v\n", expect, compt1)
	}
	expect = subSet{
	}
	if reflect.DeepEqual(actv1, expect) == false {
		t.Fatalf("actv1 want:%v, got:%v\n", expect, actv1)
	}

	// completer
	moreComp, moreActv, err := compt1.completer(allSet, gram.Symb2NTerminal)
	compt1 = append(compt1, moreComp...)
	actv1 = append(actv1, moreActv...)
	expect = subSet{
		item{
			dotPos: 1,
			at: 1,
			nonTerm: "F",
			rhSide: []string{ "a" },
		},
		item{
			dotPos: 1,
			at: 1,
			nonTerm: "E",
			rhSide: []string{ "F" },
		},
		item{
			dotPos: 1,
			at: 1,
			nonTerm: "S",
			rhSide: []string{ "E" },
		},
	}
	if reflect.DeepEqual(compt1, expect) == false {
		t.Fatalf("compt1 want:%v, got:%v\n", expect, compt1)
	}
	expect = subSet{
		item{
			dotPos: 1,
			at: 1,
			nonTerm: "E",
			rhSide: []string{ "E", "Q", "F" },
		},
	}
	if reflect.DeepEqual(actv1, expect) == false {
		t.Fatalf("actv1 want:%v, got:%v\n", expect, actv1)
	}
}


func TestParse(t *testing.T) {
	gramfile := "../../samples/grammar4earley"
	gram := common.ReadGrammarFromFile(gramfile)

	res, err := parse(gram, "a - a + a")
	if err != nil {
		t.Error(err)
	}
	expect := []string{"S"}
	if reflect.DeepEqual(res, expect) == false {
		t.Fatalf("res want:%v, got:%v\n", expect, res)
	}
}