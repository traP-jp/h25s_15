package expressions

import (
	"errors"
	"math/big"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// ========== AST定義 ==========

type Expr struct {
	Left  *Term `parser:"@@"`
	Right []*Op `parser:"{ @@ }"`
}

type Op struct {
	Operator string `parser:"@(\"+\" | \"-\")"`
	Term     *Term  `parser:"@@"`
}

type Term struct {
	Left  *Factor  `parser:"@@"`
	Right []*MulOp `parser:"{ @@ }"`
}

type MulOp struct {
	Operator string  `parser:"@(\"*\" | \"/\")"`
	Factor   *Factor `parser:"@@"`
}

type Factor struct {
	Number *string `parser:"@Number"`
	Nested *Expr   `parser:"| \"(\" @@ \")\""`
}

// ========== 評価関数 ==========

func (e *Expr) Eval() (*big.Rat, error) {
	left, err := e.Left.Eval()
	if err != nil {
		return nil, err
	}
	result := new(big.Rat).Set(left)
	for _, op := range e.Right {
		val, err := op.Term.Eval()
		if err != nil {
			return nil, err
		}
		switch op.Operator {
		case "+":
			result.Add(result, val)
		case "-":
			result.Sub(result, val)
		}
	}
	return result, nil
}

func (t *Term) Eval() (*big.Rat, error) {
	left, err := t.Left.Eval()
	if err != nil {
		return nil, err
	}
	result := new(big.Rat).Set(left)
	for _, op := range t.Right {
		val, err := op.Factor.Eval()
		if err != nil {
			return nil, err
		}
		switch op.Operator {
		case "*":
			result.Mul(result, val)
		case "/":
			if val.Sign() == 0 {
				return nil, errZeroDivision
			}
			result.Quo(result, val)
		}
	}
	return result, nil
}

var errZeroDivision = errors.New("division by zero")

func (f *Factor) Eval() (*big.Rat, error) {
	if f.Number != nil {
		r := new(big.Rat)
		r.SetString(*f.Number)
		return r, nil
	}
	return f.Nested.Eval()
}

func Parser() (*participle.Parser[Expr], error) {
	return participle.Build[Expr](
		participle.Lexer(lexer.MustSimple([]lexer.SimpleRule{
			{Name: "Number", Pattern: `[0-9]`}, // 1桁の数字のみ
			{Name: "Operator", Pattern: `[-+*/()]`},
			{Name: "Whitespace", Pattern: `\s+`},
		})),
		participle.Elide("Whitespace"), // 空白文字を無視
	)
}
