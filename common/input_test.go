package common

import (
	//"fmt"
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
	expect = [][]string{[]string{"12345678"}}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{[]string{"12345678"}}
	checkRes(t, ok, res, expectOk, expect)

	input = "12345678"
	num = 2
	/*NoERule*/
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		[]string{"1", "2345678"},
		[]string{"12", "345678"},
		[]string{"123", "45678"},
		[]string{"1234", "5678"},
		[]string{"12345", "678"},
		[]string{"123456", "78"},
		[]string{"1234567", "8"},
	}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		[]string{"", "12345678"},
		[]string{"1", "2345678"},
		[]string{"12", "345678"},
		[]string{"123", "45678"},
		[]string{"1234", "5678"},
		[]string{"12345", "678"},
		[]string{"123456", "78"},
		[]string{"1234567", "8"},
		[]string{"12345678", ""},
	}
	checkRes(t, ok, res, expectOk, expect)

	/*NoERule*/
	input = "12345678"
	num = 8
	rule = NoERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		[]string{"1", "2", "3", "4", "5", "6", "7", "8"},
	}
	checkRes(t, ok, res, expectOk, expect)
	/* HasERule */
	input = "123"
	num = 3
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		[]string{"", "", "123"},
		[]string{"", "1", "23"},
		[]string{"", "12", "3"},
		[]string{"", "123", ""},
		[]string{"1", "", "23"},
		[]string{"1", "2", "3"},
		[]string{"1", "23", ""},
		[]string{"12", "", "3"},
		[]string{"12", "3", ""},
		[]string{"123", "", ""},
	}
	checkRes(t, ok, res, expectOk, expect)

	/* HasERule */
	input = "d"
	num = 3
	rule = HasERule
	res, ok = GeneratePartitions(input, num, rule)
	expect = [][]string{
		[]string{"", "", "d"},
		[]string{"", "d", ""},
		[]string{"d", "", ""},
	}
	checkRes(t, ok, res, expectOk, expect)
}