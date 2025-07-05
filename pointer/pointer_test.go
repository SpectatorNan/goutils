package pointer

import "testing"

func TestP(t *testing.T) {
	t.Logf("test")
	a := 1
	b := "2"
	c := 3.0
	d := true
	e := struct {
		A int
	}{}
	pa := Deal(a)
	pb := Deal(b)
	pc := Deal(c)
	pd := Deal(d)
	pe := Deal(e)
	t.Logf("a: %v, b: %v, c: %v, d: %v, e: %v", a, b, c, d, e)
	t.Logf("pa: %v, pb: %v, pc: %v, pd: %v, pe: %v", pa, pb, pc, pd, pe)
}

type TestA struct {
	A string
}
type TestB struct {
	B int
}

func TestDealWithNil(t *testing.T) {
	var a *string
	var b *int
	bb := 1
	b = &bb
	c := "456789"
	d := 9954

	ra := DealWithConvert(a, func(target string) *TestA {
		return &TestA{A: target}
	})
	rb := DealWithConvert(b, func(target int) *TestB {
		return &TestB{B: target}
	})
	rc := DealWithConvert(&c, func(target string) *TestA {
		return &TestA{A: target}
	})
	rd := DealWithConvert(&d, func(target int) *TestB {
		return &TestB{B: target}
	})
	t.Logf("ra: %+v", ra)
	t.Logf("rb: %+v", rb)
	t.Logf("rc: %+v", rc)
	t.Logf("rd: %+v", rd)
}
