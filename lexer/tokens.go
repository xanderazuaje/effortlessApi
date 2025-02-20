package lexer

import "fmt"

type TokenKind int

const (
	EOF = iota

	//Keywords
	TableDeclaration    // "table"
	PrivacyMark         // "!"
	EndTableDeclaration // "---"
	RelatedTo           // ->
	RelatedFrom         // <-
	OneToOne            // --
	PrivateOneToOne     // !-
	PrivateRelatedFrom  // <!
	PrivateRelatedTo    // !>

	//table declaration
	TableName
	OpenCurly
	ClosedCurly
	PropertyName //any word followed by colons
	Constraints  //sequence of words separated by commas, without space
)

func (t TokenKind) String() string {
	switch t {
	case EOF:
		return "EOF"
	case TableDeclaration:
		return "table_declaration"
	case PrivacyMark:
		return "privacy_mark"
	case EndTableDeclaration:
		return "end_table_declaration"
	case TableName:
		return "table_name"
	case OpenCurly:
		return "open_curly_brace"
	case ClosedCurly:
		return "closed_curly_brace"
	case PropertyName:
		return "property_name"
	case Constraints:
		return "constraints"
	case RelatedTo:
		return "related_to"
	case RelatedFrom:
		return "related_from"
	case OneToOne:
		return "one_to_one"
	case PrivateOneToOne:
		return "private_one_to_one"
	case PrivateRelatedFrom:
		return "private_related_from"
	case PrivateRelatedTo:
		return "private_related_to"
	default:
		return "UNKNOWN!"
	}
}

type Token struct {
	Value string
	Kind  TokenKind
}

func (t Token) Debug() {
	switch t.Kind {
	case TableName, PropertyName, Constraints:
		fmt.Printf("%s (%s)\n", t.Kind, t.Value)
	default:
		fmt.Printf("%s ()\n", t.Kind)
	}
}

func (t Token) IsOneOf(expectedTokens ...TokenKind) bool {
	for _, v := range expectedTokens {
		if t.Kind == v {
			return true
		}
	}
	return false
}
