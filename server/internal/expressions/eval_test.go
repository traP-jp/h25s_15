package expressions_test

import (
	"testing"

	"github.com/traP-jp/h25s_15/internal/expressions"
)

func TestEval(t *testing.T) {
	parser, err := expressions.Parser()
	if err != nil {
		t.Fatalf("failed to create parser: %v", err)
	}

	testCases := []struct {
		expr     string
		expected string
	}{
		{"1 + 2 * 3 - 4 / 5", "31/5"},
		{"(1 + 2 * 3 - 4 / 5)", "31/5"},
		{"(1 + 2) * 3 - 4 / 5", "41/5"},
		{"1 + (2 * 3) - (4 / 5)", "31/5"},
		{"1 + 2 * (3 - 4/5)", "27/5"},
		{"1 + 2 * 3 - 4 / 5 + (6 / 7)", "247/35"},
		{"2+8", "10"},
	}

	for _, tc := range testCases {
		t.Run(tc.expr, func(t *testing.T) {
			t.Parallel() // Run tests in parallel for efficiency
			expr, err := parser.ParseString("", tc.expr)
			if err != nil {
				t.Fatalf("failed to parse expression: %v", err)
			}

			result, _ := expr.Eval()
			if result.RatString() != tc.expected {
				t.Errorf("unexpected result for '%s': got %s, want %s", tc.expr, result.RatString(), tc.expected)
			}
		})
	}
}
