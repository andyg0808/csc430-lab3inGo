package main

import (
	"fmt"
)

type Value fmt.Stringer

type ExprC interface {
	Interp() Value
}

type NumV struct {
	i int
}

func (n NumV) String() string {
	return fmt.Sprint(n.i)
}

type NumC struct {
	X int
}

func (n NumC) Interp() Value {
	return NumV{n.X}
}

type BinC struct {
	op string
	L ExprC
	R ExprC
}

func (b BinC) Interp() Value {
	vL := b.L.Interp()
	vR := b.R.Interp()
	return NumV{vL.(NumV).i + vR.(NumV).i}
}

func main() {
	b := &BinC{"+", NumC{1}, NumC{2}}
	fmt.Println(b)
	fmt.Println(NumC{1}.Interp())
}

