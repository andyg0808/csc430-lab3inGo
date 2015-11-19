package main

import "testing"
import "math"

func TestInterpNumC(t *testing.T) {
	n := NumC{5}
	value := n.Interp().(NumV)
	if value.i != 5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCPlus(t *testing.T) {
	b := BinC{"+", NumC{2}, NumC{3}}
	value := b.Interp().(NumV)
	if value.i != 5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCMinus(t *testing.T) {
	b := BinC{"-", NumC{2}, NumC{3}}
	value := b.Interp().(NumV)
	if value.i != -1 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCTimes(t *testing.T) {
	b := BinC{"*", NumC{2}, NumC{3}}
	value := b.Interp().(NumV)
	if value.i != 6 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCDivide(t *testing.T) {
	b := BinC{"/", NumC{2}, NumC{4}}
	value := b.Interp().(NumV)
	if value.i != 0.5 {
		t.Error("Incorrect result!")
	}
}

func TestInterpBinCDivideByZero(t *testing.T) {
	b := BinC{"/", NumC{1}, NumC{0}}
	value := b.Interp().(NumV)
	if ! math.IsInf(value.i, 1) {
		t.Error("Div by zero isn't nan")
	}
}


func TestStringNumV(t *testing.T) {
	n := NumV{1}
	if n.String() != "1" {
		t.Error("Incorrect string output from NumV")
	}
}
