package main

import (
	"strconv"
	"syscall/js"
)

type TokenType int

const (
	EQUAL TokenType = iota
	ADD
	SUB
	MUL
	DIV
	EXP
	NOP
	LPAREN
	RPAREN
	DECIMAL
	NUMBER
)

type Token struct {
	Symbol string
	T      TokenType
}

func matchToken(s string) (*Token, error) {
	switch s {
	case "=":
		return &Token{Symbol: s, T: EQUAL}, nil
	case "+":
		return &Token{Symbol: s, T: ADD}, nil
	case "-":
		return &Token{Symbol: s, T: SUB}, nil
	case "*":
		return &Token{Symbol: s, T: MUL}, nil
	case "/":
		return &Token{Symbol: s, T: DIV}, nil
	case "^":
		return &Token{Symbol: s, T: EXP}, nil
	case "(":
		return &Token{Symbol: s, T: LPAREN}, nil
	case ")":
		return &Token{Symbol: s, T: RPAREN}, nil
	case ".":
		return &Token{Symbol: s, T: DECIMAL}, nil
	default:
		if _, err := strconv.Atoi(s); err == nil {
			return &Token{Symbol: s, T: NUMBER}, nil
		}
		return &Token{Symbol: s, T: NOP}, nil
	}
}

type Pair struct {
	first  int
	second int
}

func (p Pair) First() int {
	return p.first
}

func (p Pair) Second() int {
	return p.second
}

func retrievePair(i []js.Value) Pair {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	return Pair{first: int1, second: int2}
}

func add(i []js.Value) {
	p := retrievePair(i)
	o := p.First() + p.Second()
	js.Global().Set("output", o)
	println(o)
}

func pushToQueue(i []js.Value) {
	v := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	js.Global().Get("console").Call("log", v)
}

func registerCallbacks() {
	js.Global().Set("pushToQueue", js.NewCallback(pushToQueue))
}

func main() {
	println("Go wasm initialized")
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
