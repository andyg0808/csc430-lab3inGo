package main

import (
	"fmt"
)

type Value fmt.Stringer

type ExprC interface {
	Interp() Value
}

type NumV struct {
	i float64
}

func (n NumV) String() string {
	return fmt.Sprint(n.i)
}

type BoolV struct {
	b bool
}

func (b BoolV) String() string {
	return fmt.Sprint(b.b)
}

type NumC struct {
	X float64
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
	switch b.op {
	case "+":
		return NumV{vL.(NumV).i + vR.(NumV).i}
	case "-":
		return NumV{vL.(NumV).i - vR.(NumV).i}
	case "*":
		return NumV{vL.(NumV).i * vR.(NumV).i}
	case "/":
		return NumV{vL.(NumV).i / vR.(NumV).i}
	case "eq?":
		return BoolV{vL == vR}
	}
	panic("Unknown binop")
}

func main() {
	b := &BinC{"+", NumC{1}, NumC{2}}
	fmt.Println(b)
	fmt.Println(NumC{1}.Interp())
}
