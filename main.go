package main

import (
	"fmt"
	"unicode"
)

const (
	//Token/character we don't recognize
	Illegal = "ILLEGAL"

	//End of file
	EOF = "EOF"

	//Identifiers + literals
	// String = "STRING"
	// Number = "NUMBER"

	//Valuesee
	// True  = "TRUE"
	// False = "FALSE"
	// Null  = "NULL"

	//Six structural tokens
	LeftBrace  = "{"
	RightBrace = "}"
	// LeftBracket  = "["
	// RightBracket = "]"
	// Comma        = ","
	// Colon        = ":"
)

type Token struct {
	tokenType string
	value     string
}

type Lexer struct {
	input    []rune
	position int
}

// func NewLexer(s string) *Lexer {
// 	return &Lexer{input: s}
// }

func (l *Lexer) lexer() []Token {
	tokens := []Token{}

	//skip whitespace
	for l.position < len(l.input) && unicode.IsSpace(l.input[l.position]) {
		l.position++
	}

	//End of line
	if l.position >= len(l.input) {
		return []Token{}
	}

	character := l.input[l.position]
	l.position++
	switch character {
	case '{':
		tokens = append(tokens, Token{tokenType: LeftBrace, value: "{"})
	case '}':
		tokens = append(tokens, Token{tokenType: RightBrace, value: "}"})
	default:
		tokens = append(tokens, Token{tokenType: Illegal, value: string(character)})
	}
	return tokens
}

func main() {
	lexer := &Lexer{input: []rune("{}")}

	fmt.Println(lexer)
}
