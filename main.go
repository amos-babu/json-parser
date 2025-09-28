package main

type Type string

const (
	//Token/character we don't recognize
	Illegal Type = "ILLEGAL"

	//End of file
	EOF Type = "EOF"

	//Identifiers + literals
	// String Type = "STRING"
	// Number Type = "NUMBER"

	//Valuesee
	// True  Type = "TRUE"
	// False Type = "FALSE"
	// Null  Type = "NULL"

	//Six structural tokens
	LeftBrace  Type = "{"
	RightBrace Type = "}"
	// LeftBracket  Type = "["
	// RightBracket Type = "]"
	// Comma        Type = ","
	// Colon        Type = ":"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
	Start   int
	End     int
}

func main() {
}
