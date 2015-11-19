package main

import "testing"

func TestInterpNumC(t *testing.T) {
	n := NumC{5}
	value := n.Interp().(NumV)
	if value.i != 5 {
		t.Error("Invalid number")
	}
}

func TestInterpBinC(t *testing.T) {
	b := BinC{"+", NumC{2}, NumC{3}}
	value := b.Interp().(NumV)
	if value.i != 5 {
		t.Error("Invalid number")
	}
}
