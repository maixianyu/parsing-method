package common

import (
	"fmt"
	"testing"
)


func TestGeneratePartitions(t *testing.T) {
	var input string
	var num int
	var res [][]string
	var expect [][]string

	input = "12345678"
	num = 0
	res = GeneratePartitions(input, num)
	expect = [][]string{}
	if StringSSEqual(res, expect) == false {
		t.Errorf("want %v, got %v", expect, res)
	}

	input = "12345678"
	num = 1
	res = GeneratePartitions(input, num)
	expect = [][]string{[]string{"12345678"}}
	if StringSSEqual(res, expect) == false {
		t.Errorf("want %v, got %v", expect, res)
	}

	input = "12345678"
	num = 2
	res = GeneratePartitions(input, num)
	expect = [][]string{
		[]string{"1", "2345678"},
		[]string{"12", "345678"},
		[]string{"123", "45678"},
		[]string{"1234", "5678"},
		[]string{"12345", "678"},
		[]string{"123456", "78"},
		[]string{"1234567", "8"},
	}
	if StringSSEqual(res, expect) == false {
		t.Errorf("want %v, got %v", expect, res)
	}

	input = "12345678"
	num = 8
	fmt.Printf("\ntesting input:%v num:%v\n", input, num)
	res = GeneratePartitions(input, num)
	expect = [][]string{
		[]string{"1", "2", "3", "4", "5", "6", "7", "8"},
	}
	if StringSSEqual(res, expect) == false {
		t.Errorf("want %v, got %v", expect, res)
	}

	fmt.Printf("\ntesting\n")
	expect = [][]string{
		[]string{"1", "2", "3", "4", "5", "6", "7"},
	}
	res, _ = genPartitionsHelper([]string{"1", "2", "3", "4"}, "567", 3)
	if StringSSEqual(res, expect) == false {
		t.Errorf("want %v, got %v", expect, res)
	}

}