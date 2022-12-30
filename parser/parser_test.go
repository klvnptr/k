package parser_test

import (
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/klvnptr/k/ast"
	"github.com/klvnptr/k/parser"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	parser *parser.Parser
}

func (suite *ParserTestSuite) TestSimple() {
	_, err := suite.parser.ParseString("i32 main() { return 0; }")
	suite.NoError(err)
}

func (suite *ParserTestSuite) TestFnCallExpr() {
	p := parser.BuildParser[parser.FnCallExpr]()
	_, err := p.ParseString("main.c", "fn(2, 3, t-- + 1, 4)")
	suite.NoError(err)
}

func (suite *ParserTestSuite) TestPrefixExpr() {
	p := parser.BuildParser[parser.Expr]()
	_, err := p.ParseString("main.c", "fib(1) - 2")
	suite.NoError(err)
}

func (suite *ParserTestSuite) TestString() {
	p := parser.BuildParser[parser.PrimaryExpr]()

	expr, err := p.ParseString("main.c", `"he\"l\tlo"`)
	suite.NoError(err)
	suite.Equal(&parser.PrimaryExpr{
		String: &parser.InternalString{
			Parts: []string{
				"he",
				"\\\"",
				"l",
				"\\t",
				"lo",
			},
		},
		Pos: lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
	}, expr)

	cop := expr.Transform(&ast.Block{})
	suite.Equal(&ast.ConstantStringOp{
		Constant: "he\"l\tlo",
		Scope:    &ast.Block{},
		Pos:      lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
	}, cop)
}

func (suite *ParserTestSuite) TestDeclarator() {
	p := parser.BuildParser[parser.Declarator]()

	expr, err := p.ParseString("main.c", "i32 **n")
	suite.NoError(err)
	suite.Equal(&parser.Declarator{
		Type: &parser.Type{
			Basic:    "i32",
			Pointers: "**",
			Pos:      lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
		},
		Ident: "n",
		Pos:   lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
	}, expr)
}

func (suite *ParserTestSuite) TestIndexExpr() {
	p := parser.BuildParser[parser.Expr]()

	expr, err := p.ParseString("main.c", "c[12]")
	suite.NoError(err)

	transformed := expr.Transform(&ast.Block{})

	suite.Equal(&ast.IndexOp{
		Expr: &ast.LoadOp{
			Name:  "c",
			Scope: &ast.Block{},
			Pos:   lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
		},
		IndexExpr: &ast.ConstantNumberOp{
			Constant: "12",
			Scope:    &ast.Block{},
			Pos:      lexer.Position{Filename: "main.c", Offset: 2, Line: 1, Column: 3},
		},
		Scope: &ast.Block{},
		Pos:   lexer.Position{Filename: "main.c", Offset: 1, Line: 1, Column: 2},
	}, transformed)

	p2 := parser.BuildParser[parser.Expr]()

	expr2, err2 := p2.ParseString("main.c", "(c[42])[t]")
	suite.NoError(err2)

	transformed = expr2.Transform(&ast.Block{})

	suite.Equal(&ast.IndexOp{
		Expr: &ast.IndexOp{
			Expr: &ast.LoadOp{
				Name:  "c",
				Scope: &ast.Block{},
				Pos:   lexer.Position{Filename: "main.c", Offset: 1, Line: 1, Column: 2},
			},
			IndexExpr: &ast.ConstantNumberOp{
				Constant: "42",
				Scope:    &ast.Block{},
				Pos:      lexer.Position{Filename: "main.c", Offset: 3, Line: 1, Column: 4},
			},
			Scope: &ast.Block{},
			Pos:   lexer.Position{Filename: "main.c", Offset: 2, Line: 1, Column: 3},
		},
		IndexExpr: &ast.LoadOp{
			Name:  "t",
			Scope: &ast.Block{},
			Pos:   lexer.Position{Filename: "main.c", Offset: 8, Line: 1, Column: 9},
		},
		Scope: &ast.Block{},
		Pos:   lexer.Position{Filename: "main.c", Offset: 7, Line: 1, Column: 8},
	}, transformed)
}

func (suite *ParserTestSuite) TestFunc() {
	p := parser.BuildParser[parser.Module]()

	src := `
i64 fib(i64 *n) {
	if (n < 2) {
		return n;
	}
	
	return fib(n - 1) + fib(n - 2);
}

i32 main() {
	return fib(10);
}
`
	_, err := p.ParseString("main.c", src)
	suite.NoError(err)
}

func (suite *ParserTestSuite) TestChar() {
	p := parser.BuildParser[parser.PrimaryExpr]()

	result, err := p.ParseString("main.c", `'a'`)
	suite.NoError(err)

	suite.Equal(&parser.PrimaryExpr{
		Char: "a",
		Pos:  lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
	}, result)
}

func (suite *ParserTestSuite) TestReturn() {
	p := parser.BuildParser[parser.Module]()

	_, err := p.ParseString("main.c", `
	i64
	main()
	{
		i64 x;
		return x - 3;
	}
	`)
	suite.NoError(err)
}

func (suite *ParserTestSuite) TestPostfixRecursion() {
	p := parser.BuildParser[parser.Expr]()

	result, err := p.ParseString("main.c", `x----`)
	suite.NoError(err)

	transformed := result.Transform(&ast.Block{})

	suite.Equal(&ast.UnaryOp{
		Op: "--",
		Expr: &ast.UnaryOp{
			Op: "--",
			Expr: &ast.LoadOp{
				Name:  "x",
				Scope: &ast.Block{},
				Pos:   lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
			},
			IsPostfix: true,
			Scope:     &ast.Block{},
			Pos:       lexer.Position{Filename: "main.c", Offset: 3, Line: 1, Column: 4},
		},
		IsPostfix: true,
		Scope:     &ast.Block{},
		Pos:       lexer.Position{Filename: "main.c", Offset: 1, Line: 1, Column: 2},
	}, transformed)
}

func (suite *ParserTestSuite) TestDecl() {
	p := parser.BuildParser[parser.DeclStmt]()

	result, err := p.ParseString("main.c", `i32 x = 0;`)
	suite.NoError(err)

	transformed := result.Transform(&ast.Block{})

	suite.Equal(&ast.DeclStmt{
		Ident: "x",
		Type:  ast.NewTypeBasic(&ast.Block{}, lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1}, ast.BasicTypeI32),
		Expr: &ast.ConstantNumberOp{
			Constant: "0",
			Scope:    &ast.Block{},
			Pos:      lexer.Position{Filename: "main.c", Offset: 8, Line: 1, Column: 9},
		},
		Scope: &ast.Block{},
		Pos:   lexer.Position{Filename: "main.c", Offset: 0, Line: 1, Column: 1},
	}, transformed)
}

func (suite *ParserTestSuite) TestStruct() {
	p0 := parser.BuildParser[parser.Declarator]()

	_, err := p0.ParseString("main.c", `
	struct {
		i32 a,
		i32 b,
	} s
	`)
	suite.NoError(err)

	_, err = p0.ParseString("main.c", `
	hello s
	`)
	suite.NoError(err)

	p1 := parser.BuildParser[parser.TypeDef]()

	_, err = p1.ParseString("main.c", `
	type struct {
		i32 a,
		i32 b,
	} hello;
	`)
	suite.NoError(err)

	p2 := parser.BuildParser[parser.PrefixExpr]()

	_, err = p2.ParseString("main.c", `
	hello {
		.a = 1 + 2,
		.b = 2,
	}
	`)
	suite.NoError(err)

	p3 := parser.BuildParser[parser.Module]()

	_, err = p3.ParseString("main.c", `
	type struct {
		i32 a,
		struct {
			i32 f,
		} b,
	} st;

	type struct {
		f64 c,
		i32 d,
	} hello;

	int main(hello *n) {
		st s = 2;
		st s2 = st {
			.a = 1,
			.b = hello {
				.c = 2 + sin(2),
				.d = 3,
			},
		};
	}
	`)
	suite.NoError(err)
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
