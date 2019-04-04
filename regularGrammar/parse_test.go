package regularGrammar

import (
	"testing"
	"reflect"
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


func TestList(t *testing.T) {
	/* test state creation */
	s1 := state(1, nil, nil)
	expect := State{
		c: 1,
		out: nil,
		out1: nil,
		lastlist: 0,
	}
	if reflect.DeepEqual(*s1, expect) == false {
		t.Fatalf("want:%v, got:%v", *s1, expect)
	}

	/* test Ptrlist creation */
	l1 := listl(s1)
	expectList := Ptrlist{
		next: nil,
		s: s1,
	}
	if reflect.DeepEqual(*l1, expectList) == false {
		t.Fatalf("want:%v, got:%v", *l1, expectList)
	}

	/* test append */
	l2 := listl(nil)
	l1 = appendList(l1, l2)
	expectList = *l1
	expectList.next = &Ptrlist{
		next: nil,
		s: nil,
	}
	if reflect.DeepEqual(*l1, expectList) == false {
		t.Fatalf("want:%v, got:%v", *l1, expectList)
	}

	/* test patch */
	s2 := state(2, nil, nil)
	patch(l1, s2)
	expectList.s = s2
	expectList.next.s = s2
	if reflect.DeepEqual(*l1, expectList) == false {
		t.Fatalf("want:%v, got:%v", *l1, expectList)
	}

}