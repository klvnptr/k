package ast

import (
	"fmt"

	"github.com/klvnptr/k/utils"
	"github.com/llir/llvm/ir"
)

func (m *Module) String() []string {
	lines := []string{}

	for _, td := range m.LocalTypes {
		lines = append(lines, td.String()...)
	}

	for _, fn := range m.Functions {
		lines = append(lines, fn.String()...)
	}

	m.Scope.PrefixLines(lines)

	lines = append([]string{"module \"" + m.Scope.File.Name + "\" {"}, lines...)
	lines = append(lines, "}")

	return lines
}

func (m *Module) Current() *Scope {
	return m.Scope
}

func (m *Module) AddModuleTypeDef(alias string, typ *Type) {
	m.ModuleTypeDefs = append(m.ModuleTypeDefs, &TypeDef{
		Alias: alias,
		Type:  typ,
	})
}

func (m *Module) AddLocalType(alias string, typ *Type) {
	m.LocalTypes = append(m.LocalTypes, &TypeDef{
		Alias: alias,
		Type:  typ,
	})
}

func (m *Module) AddGlobal(v *Variable) {
	m.Globals = append(m.Globals, &Global{
		Variable: v,
	})
}

func (m *Module) AddLocal(v *Variable) error {
	return fmt.Errorf("cannot add local variable to module")
}

func (m *Module) FindTypeDefByAlias(alias string) *TypeDef {
	for _, td := range m.LocalTypes {
		if td.Alias == alias {
			return td
		}
	}

	return nil
}

func (m *Module) FindTypeDefByType(typ *Type) *TypeDef {
	for _, td := range m.LocalTypes {
		if td.Type.Equals(typ) {
			return td
		}
	}

	return nil
}

func (m *Module) FindVariable(ident string) *Variable {
	for _, g := range m.Globals {
		if g.Variable.Ident == ident {
			return g.Variable
		}
	}

	return nil
}

func (m *Module) FindFunction(ident string) *Function {
	for _, fn := range m.Functions {
		if fn.Name == ident {
			return fn
		}
	}

	return nil
}

func (m *Module) CurrentModule() *Module {
	return m
}

func (m *Module) CurrentFunction() *Function {
	return nil
}

func (m *Module) BasicBlock() *ir.Block {
	return m.CurrentBlock
}

func (m *Module) SetBasicBlock(b *ir.Block) {
	m.CurrentBlock = b
}

func (m *Module) Generate() (*ir.Module, error) {
	m.Ptr = ir.NewModule()

	if len(m.Globals) > 0 {
		panic("not implemented")
	}

	for _, fn := range m.Functions {
		if err := fn.Generate(); err != nil {
			return nil, err
		}
	}

	// new TypeDefs might be added during function generation
	for _, td := range m.ModuleTypeDefs {
		for _, el := range m.Ptr.TypeDefs {
			if el.Name() == td.Alias {
				return nil, utils.WithPos(fmt.Errorf("type alias %s already exists", td.Alias), m.Current().File, td.Type.Pos)
			}
		}

		irType, err := td.Type.IRType()
		if err != nil {
			return nil, err
		}

		// global module type definition must have a name set, because ir.Module LLString() will use it.
		// normally this is done by ir.Module.NewTypeDef() but it has side-effects, so we do it manually.
		// see issue: https://github.com/llir/llvm/issues/226
		if irType.Name() != td.Alias {
			panic("IR type name does not match alias")
		}

		m.Ptr.TypeDefs = append(m.Ptr.TypeDefs, irType)
	}

	return m.Ptr, nil
}

func (m *Module) GenerateID(prefix string) string {
	id := fmt.Sprintf("%s.%d", prefix, m.LastID)

	m.LastID++

	return id
}
