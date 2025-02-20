package lexer

import (
	"regexp"
	"strings"
)

func defaultHandler(kind TokenKind, value string) RegexHandler {
	return func(lex *Lexer, pattern *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(Token{Value: value, Kind: kind})
	}
}
func matchHandler(kind TokenKind) RegexHandler {
	return func(lex *Lexer, pattern *regexp.Regexp) {
		match := pattern.FindString(lex.remainder())
		lex.advanceN(len(match))
		match = strings.Trim(match, ":\"")
		lex.push(Token{
			Value: match,
			Kind:  kind,
		})
	}
}
func skipHandler(lex *Lexer, pattern *regexp.Regexp) {
	match := pattern.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}
