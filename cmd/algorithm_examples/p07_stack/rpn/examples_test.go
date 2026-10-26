package main

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"классический пример", []string{"2", "1", "+", "3", "*"}, 9},
		{"деление с остатком", []string{"4", "13", "5", "/", "+"}, 6},
		{"сложное выражение", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{"одно число", []string{"42"}, 42},
		{"отрицательное число", []string{"-1", "2", "*"}, -2},
		{"вычитание не коммутирует", []string{"3", "2", "-"}, 1},
		{"деление к нулю", []string{"-7", "2", "/"}, -3},
		{"минус впереди", []string{"0", "3", "-"}, -3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EvalRPN(tc.tokens); got != tc.want {
				t.Errorf("EvalRPN(%v) = %d, want %d", tc.tokens, got, tc.want)
			}
		})
	}
}

func TestCalcStringII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"приоритет умножения", "3+2*2", 7},
		{"деление без остатка", " 3/2 ", 1},
		{"деление с остатком", " 3+5 / 2 ", 5},
		{"одно число", "42", 42},
		{"вычитание", " 13- 5 ", 8},
		{"только умножение", "2*3*4", 24},
		{"минус потом деление", "14-3/2", 13},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CalcStringII(tc.s); got != tc.want {
				t.Errorf("CalcStringII(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}

func TestCalcStringIIMatchesEvalRPN(t *testing.T) {
	// Оба вычислителя обязаны дать одинаковый результат:
	// инфиксная запись и её ОПН-эквивалент.
	pairs := []struct {
		infix string
		rpn   []string
	}{
		{"3+2*2", []string{"3", "2", "2", "*", "+"}},
		{"1+2*3-4", []string{"1", "2", "3", "*", "4", "-", "+"}},
		{"100-25*2+5", []string{"100", "25", "2", "*", "-", "5", "+"}},
	}
	for _, p := range pairs {
		gotInfix := CalcStringII(p.infix)
		gotRPN := EvalRPN(p.rpn)
		if gotInfix != gotRPN {
			t.Errorf("%q: infix = %d, rpn = %d", p.infix, gotInfix, gotRPN)
		}
	}
}
