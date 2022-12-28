package ast

import "github.com/alecthomas/participle/v2/lexer"

func NewTypeBasic(scope ScopeLike, pos lexer.Position, typ BasicType) *Type {
	return &Type{
		_basic: typ,
		Scope:  scope,
		Pos:    pos,
	}
}

func NewTypeAlias(scope ScopeLike, pos lexer.Position, alias string) *Type {
	return &Type{
		_alias: alias,
		Scope:  scope,
		Pos:    pos,
	}
}

func NewTypeStruct(scope ScopeLike, pos lexer.Position, fields ...*StructField) *Type {
	return &Type{
		_struct: &StructType{
			Fields: fields,
		},
		Scope: scope,
		Pos:   pos,
	}
}

func (t *Type) NewPointer() *Type {
	return &Type{
		_pointer: t,
		Scope:    t.Scope,
		Pos:      t.Pos,
	}
}
