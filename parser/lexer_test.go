package parser_test

import (
	"fmt"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/klvnptr/k/parser"
	"github.com/stretchr/testify/suite"
)

type LexerTestSuite struct {
	suite.Suite
	lexer *lexer.StatefulDefinition

	symbols map[string]lexer.TokenType
	reverse map[lexer.TokenType]string
}

func (suite *LexerTestSuite) SetupTest() {
	suite.lexer = parser.BuildLexer()
	suite.symbols = suite.lexer.Symbols()
	suite.reverse = make(map[lexer.TokenType]string)

	for k, v := range suite.symbols {
		suite.reverse[v] = k
	}
}

func (suite *LexerTestSuite) EqualToken(l lexer.Lexer, token, value string) {
	tok, err := l.Next()
	suite.NoError(err)
	suite.Equal(token, suite.reverse[tok.Type])
	suite.Equal(value, tok.Value)
}

func (suite *LexerTestSuite) DumpAllTokens(l lexer.Lexer) {
	for {
		tok, err := l.Next()

		if err != nil {
			panic(err)
		}

		if tok.EOF() {
			break
		}

		fmt.Println(fmt.Sprintf("suite.EqualToken(tokens, \"%s\", `%s`)", suite.reverse[tok.Type], tok.Value))
	}
}

func (suite *LexerTestSuite) TestSimple() {
	tokens, err := suite.lexer.LexString("main.c", "int main() { return 0; }")
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Ident", `int`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `main`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `{`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Keyword", `return`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Number", `0`)
	suite.EqualToken(tokens, "Punct", `;`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `}`)
}

func (suite *LexerTestSuite) TestString() {
	tokens, err := suite.lexer.LexString("main.c", `int main() { return "h\"ali" + "hello"; }`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Ident", `int`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `main`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `{`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Keyword", `return`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "StringStart", `"`)
	suite.EqualToken(tokens, "Chars", `h`)
	suite.EqualToken(tokens, "Escaped", `\"`)
	suite.EqualToken(tokens, "Chars", `ali`)
	suite.EqualToken(tokens, "StringEnd", `"`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `+`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "StringStart", `"`)
	suite.EqualToken(tokens, "Chars", `hello`)
	suite.EqualToken(tokens, "StringEnd", `"`)
	suite.EqualToken(tokens, "Punct", `;`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `}`)
}

func (suite *LexerTestSuite) TestComment() {
	tokens, err := suite.lexer.LexString("main.c", `
int main() { // huha
	// hello
	return 0; // hallo
}
`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Whitespace", `
`)
	suite.EqualToken(tokens, "Ident", `int`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `main`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `{`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Comment", `// huha
`)
	suite.EqualToken(tokens, "Whitespace", `	`)
	suite.EqualToken(tokens, "Comment", `// hello
`)
	suite.EqualToken(tokens, "Whitespace", `	`)
	suite.EqualToken(tokens, "Keyword", `return`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Number", `0`)
	suite.EqualToken(tokens, "Punct", `;`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Comment", `// hallo
`)
	suite.EqualToken(tokens, "Punct", `}`)
	suite.EqualToken(tokens, "Whitespace", `
`)
}

func (suite *LexerTestSuite) TestFnCall() {
	tokens, err := suite.lexer.LexString("main.c", `int main() { abc(); }`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Ident", `int`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `main`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `{`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `abc`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Punct", `;`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `}`)
}

func (suite *LexerTestSuite) TestIndexExpr() {
	tokens, err := suite.lexer.LexString("main.c", `int main() { c[12]; }`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Ident", `int`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `main`)
	suite.EqualToken(tokens, "Punct", `(`)
	suite.EqualToken(tokens, "Punct", `)`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `{`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Ident", `c`)
	suite.EqualToken(tokens, "Punct", `[`)
	suite.EqualToken(tokens, "Number", `12`)
	suite.EqualToken(tokens, "Punct", `]`)
	suite.EqualToken(tokens, "Punct", `;`)
	suite.EqualToken(tokens, "Whitespace", ` `)
	suite.EqualToken(tokens, "Punct", `}`)
}

func (suite *LexerTestSuite) TestPrefixExpr() {
	tokens, err := suite.lexer.LexString("main.c", `++ha`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Punct", `+`)
	suite.EqualToken(tokens, "Punct", `+`)
	suite.EqualToken(tokens, "Ident", `ha`)

	tokens, err = suite.lexer.LexString("main.c", `[ha]`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Punct", `[`)
	suite.EqualToken(tokens, "Ident", `ha`)
	suite.EqualToken(tokens, "Punct", `]`)

	tokens, err = suite.lexer.LexString("main.c", `---ha---`)
	suite.NoError(err)
	// suite.DumpAllTokens(tokens)

	suite.EqualToken(tokens, "Punct", `-`)
	suite.EqualToken(tokens, "Punct", `-`)
	suite.EqualToken(tokens, "Punct", `-`)
	suite.EqualToken(tokens, "Ident", `ha`)
	suite.EqualToken(tokens, "Punct", `-`)
	suite.EqualToken(tokens, "Punct", `-`)
	suite.EqualToken(tokens, "Punct", `-`)
}

func TestLexerTestSuite(t *testing.T) {
	suite.Run(t, new(LexerTestSuite))
}
