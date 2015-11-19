package main

import (
	"fmt"
	"container/list"
)

type Value fmt.Stringer

type ExprC interface {
	Interp(env Env) Value
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

func (n NumC) Interp(env Env) Value {
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

func (i ifC) Interp(env Env) Value {
    if i.B == true {
        return i.X.Interp(env)
    }

    return i.Y.Interp(env)
}

type idC struct {
    S string
}

func (i idC) Interp(env Env) Value {
    fmt.Println(i.Interp(env))
    panic("ahh")
}

func (b BinC) Interp(env Env) Value {
	vL := b.L.Interp(env)
	vR := b.R.Interp(env)
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
	params *list.List
	body ExprC
	env Env
}

type Binding struct {
	name string
	val Value
}

type Env struct {
	bindings *list.List
}

func (c CloV) String() string {
	return fmt.Sprint("#<procudure>")
}

type AppC struct {
	fun ExprC
	arg *list.List
}

func (a AppC) Interp(env Env) Value {
	f := a.fun.Interp(env)
	switch f := f.(type) {
	case CloV:
		newEnv := createNewEnv(f.params, interpAll(a.arg, env), f.env)
		return f.body.Interp(newEnv) 
	default:
		fmt.Println("Application of non-closure")
	}
	return nil
}

func interpAll(arg list.List, env Env) *list.List {
	l := list.New()
	for e := arg.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value.(ExprC).Interp(env))
	}
	return l
}

func createNewEnv(params list.List, arg *list.List, env Env) Env {
	if params.Len() == arg.Len() {
		p := params.Front()
		a := arg.Front()
		for {
			if a == nil {
				break
			}

			env.bindings.PushFront(Binding{p.Value.(string), a.Value.(Value)})

			p = p.Next()
			a = a.Next()
		}
	} else {
		panic("Unequal arguments and parameters")
	}
	return env
}

func main() {
	b := &BinC{"+", NumC{1}, NumC{2}}

    c := &ifC{true, NumC{1}, NumC{2}}
    d := &ifC{false, NumC{1}, NumC{2}}

    fmt.Println(b)
    //fmt.Println(b.Interp())

    fmt.Println(c)
    //fmt.Println(c.Interp())

    fmt.Println(d)
    //fmt.Println(d.Interp())
}