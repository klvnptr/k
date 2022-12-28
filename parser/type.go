package parser

import (
	"github.com/klvnptr/k/ast"
)

func (t *Type) Transform(scope ast.ScopeLike) *ast.Type {
	var typ *ast.Type

	ptrLevel := len(t.Pointers)

	if t.Struct != nil {
		fields := []*ast.StructField{}

		for _, f := range t.Struct.Fields {
			fields = append(fields, &ast.StructField{
				Ident: f.Ident,
				Type:  f.Type.Transform(scope),
			})
		}

		typ = ast.NewTypeStruct(scope, t.Pos, fields...)
	} else if t.Basic != "" {
		typ = ast.NewTypeBasic(scope, t.Pos, ast.BasicType(t.Basic))
	} else if t.Alias != "" {
		typ = ast.NewTypeAlias(scope, t.Pos, t.Alias)
	}

	for i := 0; i < ptrLevel; i++ {
		typ = typ.NewPointer()
	}

	return typ
}
