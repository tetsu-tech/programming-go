package main

import (
	"fmt"
	"math"
	"testing"
)

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
}

// Varは変数を特定する
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literalは数値定数
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unaryは単項演算式を表す
type unary struct {
	op rune // '+' か '-'のどちらか
	x  Expr
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

type binary struct {
	op   rune // '+'、'-'、'*'、'/'のどれか
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.x.Eval(env)
	case '*':
		return b.x.Eval(env) * b.x.Eval(env)
	case '/':
		return b.x.Eval(env) / b.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn   string // "pow"、"sin"、"sqrt"のどれか
	args []Expr
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
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
	}

	var prevExpr string
	for _, test := range tests {
		// 変更されている時だけexprを表示する
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // パースエラー
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
