package expressions

import (
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

func (e *Expr) Eval() *big.Rat {
	result := new(big.Rat).Set(e.Left.Eval())
	for _, op := range e.Right {
		val := op.Term.Eval()
		switch op.Operator {
		case "+":
			result.Add(result, val)
		case "-":
			result.Sub(result, val)
		}
	}
	return result
}

func (t *Term) Eval() *big.Rat {
	result := new(big.Rat).Set(t.Left.Eval())
	for _, op := range t.Right {
		val := op.Factor.Eval()
		switch op.Operator {
		case "*":
			result.Mul(result, val)
		case "/":
			result.Quo(result, val)
		}
	}
	return result
}

func (f *Factor) Eval() *big.Rat {
	if f.Number != nil {
		r := new(big.Rat)
		r.SetString(*f.Number)
		return r
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
