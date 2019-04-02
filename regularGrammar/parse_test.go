package regularGrammar

import (
	"testing"
)

func TestRe2post(t *testing.T) {
	input := []rune("a(b|c)*d")
	res, err := re2post(input)
	expect := "abc|*.d."
	if err != nil || string(res) != expect {
		t.Fatalf("err=%v, want: %v, got: %v", err, expect, string(res))
	}

	input = []rune("a((b|c)+|d)*e")
	res, err = re2post(input)
	expect = "abc|+d|*.e."
	if err != nil || string(res) != expect {
		t.Fatalf("err=%v, want: %v, got: %v", err, expect, string(res))
	}

	input = []rune("a(*e")
	res, err = re2post(input)
	if err == nil {
		t.Fatalf("err=%v, want: %v, got: %v", err, expect, string(res))
	}
}