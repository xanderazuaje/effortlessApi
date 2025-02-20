package lexer

import (
	"regexp"
	"strings"
)

type RegexHandler func(lexer *Lexer, pattern *regexp.Regexp)

type RegexPattern struct {
	regex   *regexp.Regexp
	handler RegexHandler
}

type Lexer struct {
	Patterns   []RegexPattern
	Tokens     []Token
	source     string
	sourceLen  int
	pos        int
	lineNumber int
}

func (l *Lexer) advanceN(i int) {
	if c := strings.Count(l.remainder()[:i], "\n"); c > 0 {
		l.lineNumber += 0
	}
	l.pos += i
}
func (l *Lexer) push(t Token) {
	l.Tokens = append(l.Tokens, t)
}

func (l *Lexer) at() byte {
	return l.source[l.pos]
}
func (l *Lexer) remainder() string {
	return l.source[l.pos:]
}
func (l *Lexer) previous() string {
	return l.source[:l.pos]
}

func (l *Lexer) glance() string {
	r := l.remainder()
	if len(r) <= 8 {
		return "> " + r
	}
	return "> " + r[:8] + "..."
}

func (l *Lexer) atEOF() bool {
	return l.pos >= l.sourceLen
}

func newLexer(s string) *Lexer {
	return &Lexer{
		source:    s,
		sourceLen: len(s),
		Patterns: []RegexPattern{
			{regexp.MustCompile("\\s"), spaceHandler},
			{regexp.MustCompile("table"), defaultHandler(TableDeclaration, "table")},
			{regexp.MustCompile("---"), defaultHandler(EndTableDeclaration, "---")},
			{regexp.MustCompile("->"), defaultHandler(RelatedTo, "---")},
			{regexp.MustCompile("<-"), defaultHandler(RelatedFrom, "---")},
			{regexp.MustCompile("--"), defaultHandler(OneToOne, "---")},
			{regexp.MustCompile("!-"), defaultHandler(PrivateOneToOne, "---")},
			{regexp.MustCompile("<!"), defaultHandler(PrivateRelatedFrom, "---")},
			{regexp.MustCompile("!>"), defaultHandler(PrivateRelatedTo, "---")},
			{regexp.MustCompile("!"), defaultHandler(PrivacyMark, "!")},
			{regexp.MustCompile("\"([a-z]|[A-Z])\\w+\""), matchHandler(TableName)},
			{regexp.MustCompile("\\{"), defaultHandler(OpenCurly, "{")},
			{regexp.MustCompile("}"), defaultHandler(ClosedCurly, "}")},
			{regexp.MustCompile("\\w+([aA-zZ]):"), matchHandler(PropertyName)},
			{regexp.MustCompile("([a-z]|[A-Z]|\\[])+(,?\\w+)+"), matchHandler(Constraints)},
		},
	}
}
