package interpreter

import (
	"fmt"
	"strconv"
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
	Value   interface{}
	process Process
}

func NewToken(t string, v interface{}) *Token {
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
	Text              string
	Pos               int
	CurrentToken      *Token
	IntegerTokenStack *Stack
	OpTokenStack      *Stack
}

func NewInterpreter(text string) *Interpreter {
	return &Interpreter{Text: text, Pos: 0, CurrentToken: nil, IntegerTokenStack: new(Stack), OpTokenStack: new(Stack)}
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
			num := text[pos:end]
			atoi, _ := strconv.Atoi(num)
			return NewToken(INTEGER, atoi)
		case ' ':
			pos = pos + 1
			continue
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

	for i.CurrentToken = i.GetNextToken(); i.CurrentToken.Type != EOF; i.CurrentToken = i.GetNextToken() {
		if i.CurrentToken.Type == INTEGER {
			i.IntegerTokenStack.Push(i.CurrentToken)
		} else if i.CurrentToken.Type == OP {
			i.OpTokenStack.Push(i.CurrentToken)
		}
		if i.IntegerTokenStack.Size() == 2 {
			s := i.IntegerTokenStack.Pop().(*Token)
			f := i.IntegerTokenStack.Pop().(*Token)
			op := i.OpTokenStack.Pop().(*Token)
			i.IntegerTokenStack.Push(NewToken(INTEGER, op.process(f.Value.(int), s.Value.(int))))
		}
		fmt.Println(i.CurrentToken.Value)
	}
	return i.IntegerTokenStack.Pop().(*Token).Value.(int)
	//i.CurrentToken = i.GetNextToken()
	//left := i.CurrentToken
	//i.Eat(INTEGER)
	//op := i.CurrentToken
	//i.Eat(OP)
	//right := i.CurrentToken
	//i.Eat(INTEGER)
	//
	//lv, err1 := strconv.Atoi(left.Value)
	//rv, err2 := strconv.Atoi(right.Value)
	//if err1 != nil || err2 != nil {
	//	panic("Invalid integer")
	//}
	//return op.process(lv, rv)
}
