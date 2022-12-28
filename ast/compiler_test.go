package ast_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/klvnptr/k/ast"
	"github.com/klvnptr/k/parser"
	"github.com/klvnptr/k/utils"
	"github.com/llir/llvm/ir"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type Option interface{ Option() }

type OptionBasename struct{ File string }

func Basename(file string) *OptionBasename { return &OptionBasename{File: file} }
func (o *OptionBasename) Option()          {}

func OverrideBasename(def string, opts []Option) string {
	for _, o := range opts {
		if ob, ok := o.(*OptionBasename); ok {
			return ob.File
		}
	}

	return def
}

type OptionHeader struct{ Header string }

func Header(header string) *OptionHeader { return &OptionHeader{Header: header} }
func (o *OptionHeader) Option()          {}

func JoinHeaders(opts []Option) string {
	headers := []string{}
	for _, o := range opts {
		if oh, ok := o.(*OptionHeader); ok {
			headers = append(headers, fmt.Sprintf("#include <%s>", oh.Header))
		}
	}

	return strings.Join(headers, "\n")
}

type OptionDeclare struct{ Declare []string }

func Declare(declare string) *OptionDeclare { return &OptionDeclare{Declare: []string{declare}} }
func DeclareMalloc() *OptionDeclare {
	return &OptionDeclare{Declare: []string{
		"i8* malloc(i64 size);",
		"i8 free(i8* ptr);",
		"i8* memset(i8* ptr, i8 val, i64 size);",
	}}
}
func (o *OptionDeclare) Option() {}

func JoinDeclares(opts []Option) string {
	Declares := []string{}
	for _, o := range opts {
		if oh, ok := o.(*OptionDeclare); ok {
			Declares = append(Declares, oh.Declare...)
		}
	}

	return strings.Join(Declares, "\n")
}

type CompilerSuite struct {
	suite.Suite

	parser    *parser.Parser
	clangPath string
	tmpFolder string
}

func NewCompilerTester() *CompilerSuite {
	return &CompilerSuite{}
}

func (suite *CompilerSuite) Init() error {
	suite.parser = parser.NewParser()

	path, err := exec.LookPath("clang")
	if err != nil {
		return errors.Wrap(err, "failed to find clang")
	}
	suite.clangPath = path

	dir, err := os.MkdirTemp("", "konstruktor-c")
	if err != nil {
		return errors.Wrap(err, "failed to create temp dir")
	}
	suite.tmpFolder = dir

	return nil
}

func (suite *CompilerSuite) Destroy() error {
	err := os.RemoveAll(suite.tmpFolder)
	if err != nil {
		return errors.Wrap(err, "failed to remove temp dir")
	}

	return nil
}

func (suite *CompilerSuite) SetupTest() {
	err := suite.Init()
	if err != nil {
		panic(err)
	}
}

func (suite *CompilerSuite) TearDownTest() {
	err := suite.Destroy()
	if err != nil {
		panic(err)
	}
}

func (suite *CompilerSuite) runClang(srcFilePath, binaryFilePath string) (string, error) {
	cmd := exec.Command(suite.clangPath, "-o", binaryFilePath, srcFilePath) //nolint:gosec
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func (suite *CompilerSuite) runBinary(binaryFilePath string) (string, error) {
	cmd := exec.Command(binaryFilePath) //nolint:gosec
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func (suite *CompilerSuite) EqualProgramK(src, expected string, opts ...Option) {
	basename := OverrideBasename("main", opts)

	input := utils.NewFile(OverrideBasename(basename, opts), src)
	scope := ast.NewScope(input)

	parsedAst, err := suite.parser.ParseFile(input)
	suite.NoError(err, "parse error:\nsource:\n%s", src)
	if err != nil {
		return
	}

	transformedAst := parsedAst.Transform(scope)

	bitCode, err := transformedAst.Generate()
	suite.NoError(err, "generate error:\nsource:\n%s", src)
	if err != nil {
		return
	}

	bitCodeStr := bitCode.String()

	// fmt.Println(bitCodeStr)

	srcFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.ll", basename))
	err = os.WriteFile(srcFilePath, []byte(bitCodeStr), 0600)
	suite.NoError(err)
	if err != nil {
		return
	}

	binaryFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.bin", basename))
	result, err := suite.runClang(srcFilePath, binaryFilePath)
	suite.NoError(err, "clang error:\n%s\nsource:\n%s\n", result, bitCodeStr)
	if err != nil {
		return
	}

	result, err = suite.runBinary(binaryFilePath)
	suite.NoError(err, "binary error:\n%s\nsource:\n%s\n", result, bitCodeStr)
	if err != nil {
		return
	}

	suite.Equal(expected, result, "output mismatch, source:\n%s\n", bitCodeStr)
}

func (suite *CompilerSuite) ErrorParseProgramK(src, expected string, opts ...Option) {
	basename := OverrideBasename("main", opts)

	input := utils.NewFile(OverrideBasename(basename, opts), src)
	scope := ast.NewScope(input)

	parsedAst, err := suite.parser.ParseFile(input)
	if err != nil {
		suite.Contains(err.Error(), expected, "parse error:\n%s\nsource:\n%s", err.Error(), src)
		return
	}

	transformedAst := parsedAst.Transform(scope)

	_, err = transformedAst.Generate()
	if err != nil {
		suite.Contains(err.Error(), expected, "generate error:\n%s\nsource:\n%s", err.Error(), src)
		return
	}

	suite.Fail("expected error, but got none", "source:\n%s", src)
}

func (suite *CompilerSuite) ErrorClangProgramK(src, contains string, opts ...Option) {
	basename := OverrideBasename("main", opts)

	input := utils.NewFile(OverrideBasename(basename, opts), src)
	scope := ast.NewScope(input)

	parsedAst, err := suite.parser.ParseFile(input)
	suite.NoError(err, "parse error:\n%s\nsource:\n%s", err.Error(), src)
	if err != nil {
		return
	}

	transformedAst := parsedAst.Transform(scope)

	bitCode, err := transformedAst.Generate()
	suite.NoError(err, "generate error:\n%s\nsource:\n%s", err.Error(), src)
	if err != nil {
		return
	}

	bitCodeStr := bitCode.String()

	srcFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.ll", basename))
	err = os.WriteFile(srcFilePath, []byte(bitCodeStr), 0600)
	suite.NoError(err)
	if err != nil {
		return
	}

	binaryFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.bin", basename))

	result, _ := suite.runClang(srcFilePath, binaryFilePath)

	suite.Contains(result, contains, "output mismatch, source:\n%s\n", bitCodeStr)
}

func (suite *CompilerSuite) exprToK(expr, format string, opts []Option) string {
	declares := JoinDeclares(opts)

	src := fmt.Sprintf(`
%s
i64 printf(i8 *fmt,... );

i64 main() {
	%s
	printf(%s);
	return 0;
}
`, declares, expr, format)

	return src
}

func (suite *CompilerSuite) EqualExprK(expr, format, expected string, opts ...Option) {
	src := suite.exprToK(expr, format, opts)
	suite.EqualProgramK(src, expected, opts...)
}

func (suite *CompilerSuite) ErrorParseExprK(expr, contains string, opts ...Option) {
	src := suite.exprToK(expr, `"a"`, opts)
	suite.ErrorParseProgramK(src, contains, opts...)
}

func (suite *CompilerSuite) ErrorClangExprK(expr, contains string, opts ...Option) {
	src := suite.exprToK(expr, `""`, opts)
	suite.ErrorClangProgramK(src, contains, opts...)
}

func (suite *CompilerSuite) EqualProgramC(src, expected string, opts ...Option) {
	file := OverrideBasename("main", opts)

	srcFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.c", file))

	err := os.WriteFile(srcFilePath, []byte(src), 0600)
	suite.NoError(err)
	if err != nil {
		return
	}

	binaryFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.bin", file))

	result, err := suite.runClang(srcFilePath, binaryFilePath)
	suite.NoError(err, "clang error:\n%s\nsource:\n%s\n", result, src)
	if err != nil {
		return
	}

	result, err = suite.runBinary(binaryFilePath)
	suite.NoError(err, "binary error:\n%s\nsource:\n%s\n", result, src)
	if err != nil {
		return
	}

	suite.Equal(expected, result, "output mismatch, source:\n%s\n", src)
}

func (suite *CompilerSuite) ErrorClangProgramC(src, contains string, opts ...Option) {
	file := OverrideBasename("main", opts)

	srcFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.c", file))

	err := os.WriteFile(srcFilePath, []byte(src), 0600)
	suite.NoError(err)
	if err != nil {
		return
	}

	binaryFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.bin", file))

	result, _ := suite.runClang(srcFilePath, binaryFilePath)

	suite.Contains(result, contains, "output mismatch, result:\n%s\nsource:\n%s\n", result, src)
}

func (suite *CompilerSuite) exprToC(expr, format string, opts []Option) string {
	headers := JoinHeaders(opts)

	src := fmt.Sprintf(`
#include <stdio.h>
%s

int main() {
	%s
	printf(%s);
	return 0;
}
`, headers, expr, format)

	return src
}

func (suite *CompilerSuite) EqualExprC(expr, format, expected string, opts ...Option) {
	src := suite.exprToC(expr, format, opts)
	suite.EqualProgramC(src, expected, opts...)
}

func (suite *CompilerSuite) ErrorClangExprC(expr, contains string, opts ...Option) {
	src := suite.exprToC(expr, `""`, opts)
	suite.ErrorClangProgramC(src, contains, opts...)
}

func (suite *CompilerSuite) EqualLL(m *ir.Module, expected string, opts ...Option) {
	file := OverrideBasename("main", opts)

	buf := &strings.Builder{}
	_, err := m.WriteTo(buf)
	suite.NoError(err)
	if err != nil {
		return
	}

	src := buf.String()

	srcFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.ll", file))
	err = os.WriteFile(srcFilePath, []byte(src), 0600)
	suite.NoError(err)
	if err != nil {
		return
	}

	binaryFilePath := filepath.Join(suite.tmpFolder, fmt.Sprintf("%s.bin", file))
	result, err := suite.runClang(srcFilePath, binaryFilePath)
	suite.NoError(err, "clang error:\n%s\nsource:\n%s\n", result, src)
	if err != nil {
		return
	}

	result, err = suite.runBinary(binaryFilePath)
	suite.NoError(err, "binary error:\n%s\nsource:\n%s\n", result, src)
	if err != nil {
		return
	}

	suite.Equal(expected, result, "output mismatch, source:\n%s\n", src)
}
