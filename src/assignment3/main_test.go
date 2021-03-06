package main

import "testing"
import "math"
import "container/list"

func TestInterpNumC(t *testing.T) {
	n := NumC{5}
	value := n.Interp(Env{list.New()}).(NumV)
	if value.i != 5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCPlus(t *testing.T) {
	b := BinC{"+", NumC{2}, NumC{3}}
	value := b.Interp(Env{list.New()}).(NumV)
	if value.i != 5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCMinus(t *testing.T) {
	b := BinC{"-", NumC{2}, NumC{3}}
	value := b.Interp(Env{list.New()}).(NumV)
	if value.i != -1 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCTimes(t *testing.T) {
	b := BinC{"*", NumC{2}, NumC{3}}
	value := b.Interp(Env{list.New()}).(NumV)
	if value.i != 6 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCDivide(t *testing.T) {
	b := BinC{"/", NumC{2}, NumC{4}}
	value := b.Interp(Env{list.New()}).(NumV)
	if value.i != 0.5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCDivideByZero(t *testing.T) {
	b := BinC{"/", NumC{1}, NumC{0}}
	value := b.Interp(Env{list.New()}).(NumV)
	if ! math.IsInf(value.i, 1) {
		t.Error("Div by zero isn't nan")
	}
}

func TestInterpBinCEq(t *testing.T) {
	b := BinC{"eq?", NumC{2}, NumC{4}}
	value := b.Interp(Env{list.New()}).(BoolV)
	if value.b {
		t.Error("Unequal are equal!")
	}

	b = BinC{"eq?", NumC{3}, NumC{3}}
	value = b.Interp(Env{list.New()}).(BoolV)
	if ! value.b {
		t.Error("Equal are unequal!")
	}
}

func TestStringNumV(t *testing.T) {
	n := NumV{1}
	if n.String() != "1" {
		t.Error("Incorrect string output from NumV")
	}
}

func TestStringBoolV(t *testing.T) {
	b := BoolV{true}
	if b.String() != "true" {
		t.Error("Incorrect string output from BoolV")
	}
}

func TestInterpIfCTrue(t *testing.T) {
    n := &ifC{true, NumC{1}, NumC{2}}
    value := n.Interp().(NumV)
    if value.i != 1 {
        t.Error("Incorrect result!")
    }
}

func TestInterpIfCFalse(t *testing.T) {
    n := &ifC{false, NumC{1}, NumC{2}}
    value := n.Interp().(NumV)
    if value.i != 2 {
        t.Error("Incorrect result!")
    }

func TestAppC(t *testing.T) {
	c := CloV{list.New(), BinC{"+", NumC{3}, NumC{5}}, Env{list.New()}}
	a := AppC{c, list.New()}
	if a.Interp(Env{list.New()}).(NumV).i != 8 {
		t.Error("Incorrect number output from NumV")
	}
}
