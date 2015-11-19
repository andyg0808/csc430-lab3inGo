package main

import (
	"fmt"
	"container/list"
)

type Value fmt.Stringer

type ExprC interface {
	Interp() Value
	//Interp(env Env) Value
}

type NumV struct {
	i float64
}

type BoolV struct {
    b bool
}

func (n NumV) String() string {
	return fmt.Sprint(n.i)
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

type CloV struct {
	params list.List
	body ExprC
	env list.List
}

type Binding struct {
	name string
	val Value
}

type Env struct {
	bindings list.List
}

func (c CloV) String() string {
	return fmt.Sprint("#<procudure>")
}

type AppC struct {
	fun ExprC
	arg list.List
}

//COMMENTED OUT FOR NOW SO IT RUNS

/*func (a AppC) Interp(env Env) Value {
	f := a.fun.Interp()
	switch f := f.(type) {
	case CloV:
		all-f := f.body.Interp(f.params, a.arg, env)

	default:
		fmt.Println("Application of non-closure")
	}
	return nil
}

func (a.AppC) Interp(params list.List, arg list.List, env Env) Value {

}*/

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