package interpreter

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	INTEGER = "INTEGER"
	OP      = "OP"
	EOF     = "EOF"
)

type Process func(int, int) int

var plus = func(a int, b int) int {
	return a + b
}
var sub = func(a int, b int) int {
	return a - b
}

var mul = func(a int, b int) int {
	return a * b
}
var dvd = func(a int, b int) int {
	return a / b
}

type Token struct {
	Type    string
	Value   string
	process Process
}

func NewToken(t, v string) *Token {
	var fn Process
	if v == "+" {
		fn = plus
	} else if v == "-" {
		fn = sub
	} else if v == "*" {
		fn = mul
	} else if v == "/" {
		fn = dvd
	}
	return &Token{t, v, fn}
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
		case '+', '-', '*', '/':
			i.Pos = pos + 1
			return NewToken(OP, text[pos:pos+1])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			end := pos + 1
			for ; end < len(text); end++ {
				if !unicode.IsDigit(int32(text[end])) {
					break
				}
			}
			i.Pos = end
			return NewToken(INTEGER, text[pos:end])
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

func (i *Interpreter) preCompile() {
	i.Text = strings.ReplaceAll(i.Text, " ", "")
}

func (i *Interpreter) Expr() int {
	i.preCompile()
	i.CurrentToken = i.GetNextToken()
	left := i.CurrentToken
	i.Eat(INTEGER)
	op := i.CurrentToken
	i.Eat(OP)
	right := i.CurrentToken
	i.Eat(INTEGER)

	lv, err1 := strconv.Atoi(left.Value)
	rv, err2 := strconv.Atoi(right.Value)
	if err1 != nil || err2 != nil {
		panic("Invalid integer")
	}
	return op.process(lv, rv)
}
