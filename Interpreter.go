package interpreter

import (
	"fmt"
	"strconv"
)

const (
	INTEGER = "INTEGER"
	PLUS    = "PLUS"
	EOF     = "EOF"
)

type Token struct {
	Type  string
	Value string
}

func NewToken(t, v string) *Token {
	return &Token{t, v}
}

type Interpreter struct {
	Text         string
	Pos          int
	CurrentToken *Token
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{Text: text, Pos: 0, CurrentToken: nil}
}

func (i *Interpreter) GetNextToken() *Token {
	text := i.Text
	pos := i.Pos
	for pos < len(text) {
		switch text[pos] {
		case '+':
			i.Pos = pos + 1
			return NewToken(PLUS, "+")
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			i.Pos = pos + 1
			return NewToken(INTEGER, string(text[pos]))
		default:
			panic(fmt.Sprintf("Invalid character, '%c' at position %d", text[pos], pos))
		}
	}
	return NewToken(EOF, "")
}

func (i *Interpreter) Eat(tokenType string) bool {
	if i.CurrentToken.Type == tokenType {
		i.CurrentToken = i.GetNextToken()
		return true
	}
	return false
}

func (i *Interpreter) Expr() int {
	i.CurrentToken = i.GetNextToken()
	left := i.CurrentToken
	i.Eat(INTEGER)
	i.Eat(PLUS)
	right := i.CurrentToken
	i.Eat(INTEGER)

	lv, err1 := strconv.Atoi(left.Value)
	rv, err2 := strconv.Atoi(right.Value)
	if err1 != nil || err2 != nil {
		panic("Invalid integer")
	}
	return lv + rv
}
