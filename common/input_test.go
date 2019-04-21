package common

import (
	//"fmt"
	"reflect"
	"testing"
)

func checkRes(t *testing.T, ok bool, res [][]string, expectOk bool, expect [][]string) {
	if ok != expectOk || StringSSEqual(res, expect) == false {
		t.Errorf("ok: %v, want %v, got %v", ok, expect, res)
	}
}

func TestGeneratePartitionsNoERule(t *testing.T) {
	var input string
	var num int
	var res [][]string
	var expect [][]string
	var rule ERule
	var ok bool
	var expectOk bool

	input = "12345678"
	num = 0
	/*NoERule*/
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{}
	expectOk = false
	checkRes(t, ok, res, expectOk, expect)
	/*HasERule*/
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{}
	checkRes(t, ok, res, expectOk, expect)

	expectOk = true
	input = "12345678"
	num = 1
	/* NoERule */
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{{"12345678"}}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{{"12345678"}}
	checkRes(t, ok, res, expectOk, expect)

	input = "12345678"
	num = 2
	/*NoERule*/
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		{"1", "2345678"},
		{"12", "345678"},
		{"123", "45678"},
		{"1234", "5678"},
		{"12345", "678"},
		{"123456", "78"},
		{"1234567", "8"},
	}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		{"", "12345678"},
		{"1", "2345678"},
		{"12", "345678"},
		{"123", "45678"},
		{"1234", "5678"},
		{"12345", "678"},
		{"123456", "78"},
		{"1234567", "8"},
		{"12345678", ""},
	}
	checkRes(t, ok, res, expectOk, expect)

	/*NoERule*/
	input = "12345678"
	num = 8
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		{"1", "2", "3", "4", "5", "6", "7", "8"},
	}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	input = "123"
	num = 3
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		{"", "", "123"},
		{"", "1", "23"},
		{"", "12", "3"},
		{"", "123", ""},
		{"1", "", "23"},
		{"1", "2", "3"},
		{"1", "23", ""},
		{"12", "", "3"},
		{"12", "3", ""},
		{"123", "", ""},
	}
	checkRes(t, ok, res, expectOk, expect)

	/* HasERule */
	input = "d"
	num = 3
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		{"", "", "d"},
		{"", "d", ""},
		{"d", "", ""},
	}
	checkRes(t, ok, res, expectOk, expect)
}

func TestAssemblyPartitions(t *testing.T) {
	// M < N
	M := 2
	input := []string{"apple", "bar", "c", "dog"}
	res, v := AssemblyPartitions(input, M, NoERule)
	expectV := true
	if v != expectV {
		t.Errorf("v want:%v, got:%v", expectV, v)
	}
	expect := [][]string{
		[]string{"a", "pplebarcdog"},
		[]string{"ap", "plebarcdog"},
		[]string{"app", "lebarcdog"},
		[]string{"appl", "ebarcdog"},
		[]string{"apple", "barcdog"},
		[]string{"appleb", "arcdog"},
		[]string{"appleba", "rcdog"},
		[]string{"applebar", "cdog"},
		[]string{"applebarc", "dog"},
		[]string{"applebarcd", "og"},
		[]string{"applebarcdo", "g"},
	}
	if reflect.DeepEqual(res, expect) == false {
		t.Errorf("res want:%v, got:%v", expect, res)
	}

	// M == N
	M = 4
	res, v = AssemblyPartitions(input, M, NoERule)
	expectV = true
	expect = [][]string{input}
	if v != expectV {
		t.Errorf("v want:%v, got:%v", expectV, v)
	}
	if reflect.DeepEqual(res, expect) == false {
		t.Errorf("res want:%v, got:%v", expect, res)
	}

	// N = 4, M = 5, M > N
	M = 5
	res, v = AssemblyPartitions(input, M, NoERule)
	expectV = true
	if v != expectV {
		t.Errorf("v want:%v, got:%v", expectV, v)
	}
	expect = [][]string{
		{"apple", "bar", "c", "d", "og"},
		{"apple", "bar", "c", "do", "g"},
		{"apple", "b", "ar", "c", "dog"},
		{"apple", "ba", "r", "c", "dog"},
		{"a", "pple", "bar", "c", "dog"},
		{"ap", "ple", "bar", "c", "dog"},
		{"app", "le", "bar", "c", "dog"},
		{"appl", "e", "bar", "c", "dog"},
	}
	if reflect.DeepEqual(expect, res) == false {
		t.Errorf("res want:%v, got:%v", expect, res)
	}

}
