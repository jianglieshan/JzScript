package interpreter

import (
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {

	tests := []struct {
		text string
		want int
	}{
		// TODO: Add test cases.
		{"2+1", 3},
		{"2+3", 5},
		{"3+3", 6},
		{"4+5", 9},
	}
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			i := NewInterpreter(tt.text)
			if got := i.Expr(); got != tt.want {
				t.Errorf("Expr() = %v, want %v", got, tt.want)
			}
		})
	}
}
