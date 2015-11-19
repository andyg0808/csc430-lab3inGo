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

type BoolV struct {
    b bool
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

type ifC struct {
    B bool
    X ExprC
    Y ExprC
}

func (i ifC) Interp() Value {
    if i.B == true {
        return i.X.Interp()
    }

    return i.Y.Interp()
}

func (b BinC) Interp() Value {
    vL := b.L.Interp()
    vR := b.R.Interp()
    return NumV{vL.(NumV).i + vR.(NumV).i}
}

func main() {
	b := &BinC{"+", NumC{1}, NumC{2}}

    c := &ifC{true, NumC{1}, NumC{2}}
    d := &ifC{false, NumC{1}, NumC{2}}

    fmt.Println(b)
    fmt.Println(b.Interp())

    fmt.Println(c)
    fmt.Println(c.Interp())

    fmt.Println(d)
    fmt.Println(d.Interp())

}




