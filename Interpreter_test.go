package interpreter

import (
	"testing"
)

func TestInterpreter_Expr(t *testing.T) {

	tests := []struct {
		text string
		want int
	}{
		{"2+1", 3},
		{"12+3", 15},
		{"  33 + 3 ", 36},
		{"44-5", 39},
		{"45/5", 9},
		{"44*5", 220},
		{"44*5/5+5", 49},
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
