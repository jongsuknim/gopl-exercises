package eval

import (
	"fmt"
	"math"
	"strings"
)

type Var string
type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	s := fmt.Sprintf("%f", l)
	return s
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return "+" + u.x.String()
	case '-':
		return "-" + u.x.String()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	switch b.op {
	case '+':
		return b.x.String() + " + " + b.y.String()
	case '-':
		return b.x.String() + " - " + b.y.String()
	case '*':
		return b.x.String() + " * " + b.y.String()
	case '/':
		return b.x.String() + " / " + b.y.String()
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))

}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return "pow(" + c.args[0].String() + "," + c.args[1].String() + ")"
	case "sin":
		return "sin(" + c.args[0].String() + ")"
	case "sqrt":
		return "sqrt(" + c.args[0].String() + ")"
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
