package parser

import (
	"errors"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/klvnptr/k/utils"
)

func BuildParser[T any]() *participle.Parser[T] {
	return participle.MustBuild[T](
		participle.Lexer(BuildLexer()),
		participle.UseLookahead(1024),
		participle.Elide("Whitespace"),
		participle.Elide("Comment"),
	)
}

type Parser struct {
	// contains filtered or unexported fields
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parser(lexer *lexer.StatefulDefinition) *participle.Parser[Module] {
	return BuildParser[Module]()
}

func (p *Parser) ParseString(s string) (*Module, error) {
	return p.ParseFile(&utils.File{
		// TODO: fix hardcoded name
		Name:     "main.c",
		Contents: s,
	})
}

func (p *Parser) ParseFile(file *utils.File) (*Module, error) {
	l := BuildLexer()
	parser := p.Parser(l)

	program, err := parser.ParseString(file.Name, file.Contents)

	if err != nil {
		var parseError participle.Error

		if errors.As(err, &parseError) {
			return nil, utils.WithPos(err, file, parseError.Position())
		}

		return nil, err
	}

	return program, nil
}

func (p *Parser) String() string {
	parser := p.Parser(BuildLexer())

	return parser.String()
}
