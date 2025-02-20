package lexer

import "fmt"

func Tokenize(source string) []Token {
	lex := newLexer(source)

	for !lex.atEOF() {
		matched := false
		for _, r := range lex.Patterns {
			loc := r.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				r.handler(lex, r.regex)
				matched = true
				break
			}
		}
		if !matched {
			panic(fmt.Sprintf("Lexer Error!: Unrecognized token '%s'at line %d\n", lex.glance(), lex.lineNumber))
		}
	}
	lex.push(Token{Kind: EOF, Value: "EOF"})
	return lex.Tokens
}
