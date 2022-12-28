package ast

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func (b *Block) String() []string {
	lines := []string{}

	for _, stmt := range b.Stmts {
		lines = append(lines, stmt.String()...)
	}

	b.Scope.PrefixLines(lines)

	lines = append([]string{"{"}, lines...)
	lines = append(lines, "}")

	return lines
}

func (b *Block) Current() *Scope {
	return b.Scope
}

func (b *Block) AddModuleTypeDef(alias string, typ *Type) {
	b.Scope.Parent.AddModuleTypeDef(alias, typ)
}

func (b *Block) AddLocalType(alias string, typ *Type) {
	b.Scope.Parent.AddLocalType(alias, typ)
}

func (b *Block) AddGlobal(v *Variable) {
	b.Scope.Parent.AddGlobal(v)
}

func (b *Block) AddLocal(v *Variable) error {
	// check if the variable already exists in this scope
	if b.FindVariable(v.Ident) != nil {
		return fmt.Errorf("variable '%s' already exists in this scope", v.Ident)
	}

	b.Locals = append(b.Locals, v)

	return nil
}

func (b *Block) FindTypeDefByAlias(alias string) *TypeDef {
	return b.Scope.Parent.FindTypeDefByAlias(alias)
}

func (b *Block) FindTypeDefByType(typ *Type) *TypeDef {
	return b.Scope.Parent.FindTypeDefByType(typ)
}

func (b *Block) FindVariable(ident string) *Variable {
	for _, v := range b.Locals {
		if v.Ident == ident {
			return v
		}
	}

	// if we didn't find the variable in this scope, look in the parent
	return b.Scope.Parent.FindVariable(ident)
}

func (b *Block) FindFunction(ident string) *Function {
	return b.Scope.Parent.FindFunction(ident)
}

func (b *Block) CurrentModule() *Module {
	return b.Scope.Parent.CurrentModule()
}

func (b *Block) CurrentFunction() *Function {
	return b.Scope.Parent.CurrentFunction()
}

func (b *Block) BasicBlock() *ir.Block {
	return b.CurrentModule().BasicBlock()
}

func (b *Block) SetBasicBlock(bb *ir.Block) {
	b.CurrentModule().SetBasicBlock(bb)
}

func (b *Block) Generate() error {
	for _, stmt := range b.Stmts {
		if err := stmt.Generate(); err != nil {
			return err
		}
	}

	return nil
}
