package parser

import (
	"github.com/klvnptr/k/ast"
)

func (m *Module) Transform(scope *ast.Scope) *ast.Module {
	module := &ast.Module{
		ModuleTypeDefs: []*ast.TypeDef{},
		LocalTypes:     []*ast.TypeDef{},
		Functions:      []*ast.Function{},
		Scope:          scope,
		Pos:            m.Pos,
	}

	for _, td := range m.TypeDefs {
		module.LocalTypes = append(module.LocalTypes, td.Transform(module))
	}

	for _, f := range m.Functions {
		module.Functions = append(module.Functions, f.Transform(module))
	}

	return module
}

func (t *TypeDef) Transform(scope ast.ScopeLike) *ast.TypeDef {
	typ := t.Type.Transform(scope)

	return &ast.TypeDef{
		Alias: t.Ident,
		Type:  typ,
	}
}

func (f *Function) Transform(scope ast.ScopeLike) *ast.Function {
	childScope := ast.NewScopeFromParent(scope)

	fn := &ast.Function{
		Name:        f.Declarator.Ident,
		Params:      []*ast.Variable{},
		Variadic:    f.Variadic,
		Body:        []ast.StatementLike{},
		OnlyDeclare: f.OnlyDeclare,

		Scope: childScope,
		Pos:   f.Pos,
	}

	fn.ReturnType = f.Declarator.Type.Transform(fn)

	if f.Params != nil {
		for _, param := range f.Params {
			fn.Params = append(fn.Params, &ast.Variable{
				Ident:   param.Ident,
				Type:    param.Type.Transform(fn),
				IsParam: true,
			})
		}
	}

	// in case of function declaration, we don't have body
	if f.Body != nil {
		fn.Body = append(fn.Body, f.Body.Transform(fn))
	}

	return fn
}
