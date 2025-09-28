package main

import "fmt"

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

func lexer(s string) []Token {
	fmt.Println(s)
	currentPosition := 0
	tokens := []Token{}
	for currentPosition < len(s) {
		charater := string(s[currentPosition])

		if charater == "{" {
			tokens = append(tokens, Token{tokenType: LeftBrace, value: charater})
			currentPosition++
		} else if charater == "}" {
			tokens = append(tokens, Token{tokenType: RightBrace, value: charater})
			currentPosition++
		} else if charater == "EOF" {
			break
		} else {
			fmt.Println("Illegal character unable to parse")
		}
	}
	return tokens
}
func main() {
}
