package main

import (
	"fmt"
	"jz/interpreter"
)

func main() {
	interpreter := interpreter.NewInterpreter("1+ 1")
	fmt.Println(interpreter.Expr())
}
