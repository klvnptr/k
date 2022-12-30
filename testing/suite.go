package testing

import (
	"github.com/llir/llvm/ir"
	"github.com/stretchr/testify/suite"
)

type CompilerSuite struct {
	suite.Suite
}

func NewCompilerTester() *CompilerSuite {
	return &CompilerSuite{}
}

func (suite *CompilerSuite) EqualProgramK(src, expected string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	result, err := c.RunProgramK(src, opts)
	suite.NoError(err)

	suite.Equal(expected, result)
}

func (suite *CompilerSuite) ErrorParseProgramK(src, contains string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	_, err := c.RunProgramK(src, opts)
	cre := &ParseError{}
	suite.ErrorAs(err, &cre)
	suite.Contains(err.Error(), contains)
}

func (suite *CompilerSuite) ErrorGenerateProgramK(src, contains string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	_, err := c.RunProgramK(src, opts)
	cre := &GenerateError{}
	suite.ErrorAs(err, &cre)
	suite.Contains(err.Error(), contains)
}

func (suite *CompilerSuite) ErrorClangProgramK(src, contains string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	_, err := c.RunProgramK(src, opts)
	cre := &ClangRunError{}
	suite.ErrorAs(err, &cre)
	suite.Contains(err.Error(), contains)
}

func (suite *CompilerSuite) EqualExprK(expr, format, expected string, opts ...Option) {
	src := ExprToProgramK(expr, format, opts)
	suite.EqualProgramK(src, expected, opts...)
}

func (suite *CompilerSuite) ErrorParseExprK(expr, contains string, opts ...Option) {
	src := ExprToProgramK(expr, `""`, opts)
	suite.ErrorParseProgramK(src, contains, opts...)
}

func (suite *CompilerSuite) ErrorGenerateExprK(expr, contains string, opts ...Option) {
	src := ExprToProgramK(expr, `""`, opts)
	suite.ErrorGenerateProgramK(src, contains, opts...)
}

func (suite *CompilerSuite) ErrorClangExprK(expr, contains string, opts ...Option) {
	src := ExprToProgramK(expr, `""`, opts)
	suite.ErrorClangProgramK(src, contains, opts...)
}

func (suite *CompilerSuite) EqualProgramC(src, expected string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	result, err := c.RunProgramC(src, opts)
	suite.NoError(err)

	suite.Equal(expected, result)
}

func (suite *CompilerSuite) ErrorClangProgramC(src, contains string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	_, err := c.RunProgramC(src, opts)
	cre := &ClangRunError{}
	suite.ErrorAs(err, &cre)
	suite.Contains(err.Error(), contains)
}

func (suite *CompilerSuite) EqualExprC(expr, format, expected string, opts ...Option) {
	src := ExprToProgramC(expr, format, opts)
	suite.EqualProgramC(src, expected, opts...)
}

func (suite *CompilerSuite) ErrorClangExprC(expr, contains string, opts ...Option) {
	src := ExprToProgramC(expr, `""`, opts)
	suite.ErrorClangProgramC(src, contains, opts...)
}

func (suite *CompilerSuite) EqualLL(m *ir.Module, expected string, opts ...Option) {
	c := NewContext()
	defer c.Destroy()

	result, err := c.RunProgramLL(m, opts)
	suite.NoError(err)

	suite.Equal(expected, result)
}
