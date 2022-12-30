package testing

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/klvnptr/k/ast"
	"github.com/llir/llvm/ir"

	"github.com/klvnptr/k/parser"
	"github.com/klvnptr/k/utils"
)

type Context struct {
	parser    *parser.Parser
	clangPath string
	tmpFolder string
}

func NewContext() *Context {
	c := &Context{}

	c.parser = parser.NewParser()

	path, err := exec.LookPath("clang")
	if err != nil {
		panic(err)
	}
	c.clangPath = path

	dir, err := os.MkdirTemp("", "konstruktor-c")
	if err != nil {
		panic(err)
	}
	c.tmpFolder = dir

	return c
}

func (c *Context) Destroy() {
	err := os.RemoveAll(c.tmpFolder)
	if err != nil {
		panic(err)
	}
}

func (c *Context) runClang(srcFilePath, binaryFilePath string) (string, error) {
	cmd := exec.Command(c.clangPath, "-Werror", "-Wno-override-module", "-o", binaryFilePath, srcFilePath) //nolint:gosec
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func (c *Context) runBinary(binaryFilePath string) (string, error) {
	cmd := exec.Command(binaryFilePath) //nolint:gosec
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func (c *Context) RunProgramK(src string, opts []Option) (string, error) {
	basename := OverrideBasename("main", opts)

	input := utils.NewFile(OverrideBasename(basename, opts), src)
	scope := ast.NewScope(input)

	parsedAst, err := c.parser.ParseFile(input)
	if err != nil {
		return "", NewParseError(err, src)
	}

	transformedAst := parsedAst.Transform(scope)

	bitCode, err := transformedAst.Generate()
	if err != nil {
		return "", NewGenerateError(err, src)
	}

	bitCodeStr := bitCode.String()

	// fmt.Println(bitCodeStr)

	srcFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.ll", basename))

	err = os.WriteFile(srcFilePath, []byte(bitCodeStr), 0600)
	if err != nil {
		return "", err
	}

	binaryFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.bin", basename))

	result, err := c.runClang(srcFilePath, binaryFilePath)
	if err != nil {
		return "", NewClangRunError(err, src, result)
	}

	result, err = c.runBinary(binaryFilePath)
	if err != nil {
		return "", NewBinaryRunError(err, src, result)
	}

	return result, nil
}

func (c *Context) RunProgramC(src string, opts []Option) (string, error) {
	file := OverrideBasename("main", opts)

	srcFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.c", file))

	err := os.WriteFile(srcFilePath, []byte(src), 0600)
	if err != nil {
		return "", err
	}

	binaryFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.bin", file))

	result, err := c.runClang(srcFilePath, binaryFilePath)
	if err != nil {
		return "", NewClangRunError(err, src, result)
	}

	result, err = c.runBinary(binaryFilePath)
	if err != nil {
		return "", NewBinaryRunError(err, src, result)
	}

	return result, nil
}

func (c *Context) RunProgramLL(m *ir.Module, opts []Option) (string, error) {
	file := OverrideBasename("main", opts)

	buf := &strings.Builder{}
	_, err := m.WriteTo(buf)
	if err != nil {
		return "", err
	}

	src := buf.String()

	srcFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.ll", file))

	err = os.WriteFile(srcFilePath, []byte(src), 0600)
	if err != nil {
		return "", err
	}

	binaryFilePath := filepath.Join(c.tmpFolder, fmt.Sprintf("%s.bin", file))

	result, err := c.runClang(srcFilePath, binaryFilePath)
	if err != nil {
		return "", NewClangRunError(err, src, result)
	}

	result, err = c.runBinary(binaryFilePath)

	if err != nil {
		return "", NewBinaryRunError(err, src, result)
	}

	return result, nil
}
