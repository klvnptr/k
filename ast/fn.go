package ast

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func (f *Function) String() []string {
	lines := []string{}

	params := VariableList(f.Params).String()
	if f.Variadic {
		if len(params) > 0 {
			params += ", "
		}

		params += "..."
	}

	prefix := "fn"
	suffix := ""
	if f.OnlyDeclare {
		prefix = "decl"
		suffix = ";"
	}

	line := fmt.Sprintf("%s %s(%s) -> %s%s", prefix, f.Name, params, f.ReturnType.String(), suffix)

	lines = append(lines, line)

	for _, stmt := range f.Body {
		lines = append(lines, stmt.String()...)
	}

	f.Scope.PrefixLines(lines)

	return lines
}

func (f *Function) Current() *Scope {
	return f.Scope
}

func (f *Function) AddModuleTypeDef(alias string, typ *Type) {
	f.Scope.Parent.AddModuleTypeDef(alias, typ)
}

func (f *Function) AddLocalType(alias string, typ *Type) {
	f.Scope.Parent.AddLocalType(alias, typ)
}

func (f *Function) AddGlobal(v *Variable) {
	f.Scope.Parent.AddGlobal(v)
}

func (f *Function) AddLocal(v *Variable) error {
	if f.FindVariable(v.Ident) != nil {
		return fmt.Errorf("variable '%s' already exists in this scope", v.Ident)
	}

	f.Locals = append(f.Locals, v)

	return nil
}

func (f *Function) FindTypeDefByAlias(alias string) *TypeDef {
	return f.Scope.Parent.FindTypeDefByAlias(alias)
}

func (f *Function) FindTypeDefByType(typ *Type) *TypeDef {
	return f.Scope.Parent.FindTypeDefByType(typ)
}

func (f *Function) FindVariable(ident string) *Variable {
	for _, v := range f.Locals {
		if v.Ident == ident {
			return v
		}
	}

	for _, p := range f.Params {
		if p.Ident == ident {
			return p
		}
	}

	// if cannot find variable in current scope, search in parent scope
	return f.Scope.Parent.FindVariable(ident)
}

func (f *Function) FindFunction(ident string) *Function {
	return f.Scope.Parent.FindFunction(ident)
}

func (f *Function) CurrentModule() *Module {
	return f.Current().Parent.CurrentModule()
}

func (f *Function) CurrentFunction() *Function {
	return f
}

func (f *Function) BasicBlock() *ir.Block {
	return f.CurrentModule().BasicBlock()
}

func (f *Function) SetBasicBlock(block *ir.Block) {
	f.CurrentModule().SetBasicBlock(block)
}

func (f *Function) Generate() error {
	m := f.CurrentModule()

	params := []*ir.Param{}

	for _, p := range f.Params {
		typ, err := p.Type.IRType()
		if err != nil {
			return err
		}

		params = append(params, ir.NewParam(p.Ident, typ))
	}

	typ, err := f.ReturnType.IRType()
	if err != nil {
		return err
	}

	f.Ptr = m.Ptr.NewFunc(f.Name, typ, params...)

	if f.Variadic {
		f.Ptr.Sig.Variadic = true
	}

	if f.OnlyDeclare {
		return nil
	}

	// add entry block
	entry := f.Ptr.NewBlock("fn.entry")
	f.SetBasicBlock(entry)

	// initialize params
	for i, p := range f.Params {
		typ, err := p.Type.IRType()
		if err != nil {
			return err
		}

		ptr := f.BasicBlock().NewAlloca(typ)
		f.BasicBlock().NewStore(f.Ptr.Params[i], ptr)
		p.Ptr = ptr
	}

	for _, stmt := range f.Body {
		if err := stmt.Generate(); err != nil {
			return err
		}
	}

	return nil
}
