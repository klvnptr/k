package testing

import "fmt"

type ParseError struct {
	Err    error
	Source string
}

func NewParseError(err error, source string) *ParseError {
	return &ParseError{
		Err:    err,
		Source: source,
	}
}

func (pe *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s\nsource:\n%s\n", pe.Err.Error(), pe.Source)
}

type GenerateError struct {
	Err    error
	Source string
}

func NewGenerateError(err error, source string) *GenerateError {
	return &GenerateError{
		Err:    err,
		Source: source,
	}
}

func (te *GenerateError) Error() string {
	return fmt.Sprintf("code generate error: %s\nsource:\n%s\n", te.Err.Error(), te.Source)
}

type ClangRunError struct {
	Err    error
	Source string
	Result string
}

func NewClangRunError(err error, source, result string) *ClangRunError {
	return &ClangRunError{
		Err:    err,
		Source: source,
		Result: result,
	}
}

func (cre *ClangRunError) Error() string {
	return fmt.Sprintf("clang run error: %s\nresult:\n%s\nsource:\n%s\n", cre.Err.Error(), cre.Result, cre.Source)
}

type BinaryRunError struct {
	Err    error
	Source string
	Result string
}

func NewBinaryRunError(err error, source, result string) *BinaryRunError {
	return &BinaryRunError{
		Err:    err,
		Source: source,
	}
}

func (bre *BinaryRunError) Error() string {
	return fmt.Sprintf("binary error: %s\nresult:\n%s\nsource:\n%s\n", bre.Err.Error(), bre.Result, bre.Source)
}
