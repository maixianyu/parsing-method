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
		createItem(0, 0, "S", []string{ "E" }, 0),
	}
	if reflect.DeepEqual(actv0, expect) == false {
		t.Fatalf("actv0 want:%v, got:%v\n", expect, actv0)
	}

	// predictor
	pred0, err := actv0.predictor(0, gram.Symb2NTerminal)
	expect = subSet{
		createItem(0, 0, "E", []string{ "E", "Q", "F" }, 0),
		createItem(0, 0, "E", []string{ "F" }, 1),
		createItem(0, 0, "F", []string{ "a" }, 0),
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
		createItem(1, 0, "F", []string{ "a" }, 0),
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
		createItem(1, 0, "F", []string{ "a" }, 0),
		createItem(1, 0, "E", []string{ "F" }, 1),
		createItem(1, 0, "S", []string{ "E" }, 0),
	}
	if reflect.DeepEqual(compt1, expect) == false {
		t.Fatalf("compt1 want:%v, got:%v\n", expect, compt1)
	}
	expect = subSet{
		item{
			dotPos: 1,
			at: 0,
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
	expect := []string{
		"E",
		"EQF",
		"EQFQF",
		"FQFQF",
		"aQFQF",
		"a-FQF",
		"a-aQF",
		"a-a+F",
		"a-a+a",
	}
	if reflect.DeepEqual(expect, res) == false {
		t.Fatalf("parseLog want:%v, got:%v\n", expect, res)
	}
}


func TestAmbiguous(t *testing.T) {
	gramfile := "../../samples/ambiguous"
	gram := common.ReadGrammarFromFile(gramfile)

	res, err := parse(gram, "x x x")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}