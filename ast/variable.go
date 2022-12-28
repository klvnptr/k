package ast

import (
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/llir/llvm/ir/value"
)

type Variable struct {
	Ident   string
	Type    *Type
	IsParam bool

	// LLVM IR pointer to the variable
	Ptr value.Value

	Pos lexer.Position
}

func (v *Variable) String() string {
	return v.Type.String() + " " + v.Ident
}

type VariableList []*Variable

func (vl VariableList) String() string {
	code := []string{}

	for _, v := range vl {
		code = append(code, v.String())
	}

	return strings.Join(code, ", ")
}
