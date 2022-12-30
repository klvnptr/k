package ast

import (
	"fmt"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/klvnptr/k/utils"
	"github.com/llir/llvm/ir/types"
)

type BasicType string

const (
	BasicTypeBool BasicType = "bool"
	BasicTypeVoid BasicType = "void"

	BasicTypeI8  BasicType = "i8"
	BasicTypeI16 BasicType = "i16"
	BasicTypeI32 BasicType = "i32"
	BasicTypeI64 BasicType = "i64"

	BasicTypeUI8  BasicType = "u8"
	BasicTypeUI16 BasicType = "u16"
	BasicTypeUI32 BasicType = "u32"
	BasicTypeUI64 BasicType = "u64"

	BasicTypeF32 BasicType = "f32"
	BasicTypeF64 BasicType = "f64"
)

type Type struct {
	_basic   BasicType
	_array   *ArrayType
	_struct  *StructType
	_pointer *Type
	_alias   string

	_cached types.Type

	Scope ScopeLike
	Pos   lexer.Position
}

// AliasedType returns the type at the end of the alias chain.
func (t *Type) AliasedType() *Type {
	if t.IsAlias() {
		td := t.Scope.FindTypeDefByAlias(t.Alias())

		// should be impossible, because all aliased types should be defined
		if td == nil {
			panic(fmt.Sprintf("unknown type alias '%s'", t.Alias()))
		}

		return td.Type.AliasedType()
	} else {
		return t
	}
}

func (t *Type) Basic() BasicType {
	return t.AliasedType()._basic
}

func (t *Type) Array() *ArrayType {
	return t.AliasedType()._array
}

func (t *Type) Struct() *StructType {
	return t.AliasedType()._struct
}

func (t *Type) Pointer() *Type {
	return t.AliasedType()._pointer
}

func (t *Type) Alias() string {
	return t._alias
}

func (t *Type) String() string {
	// we need to check for alias first, because IsBasic() and etc. will return true for aliases
	if t.IsAlias() {
		return t.Alias()
	}

	// not an alias, so we can just check the type
	if t.IsBasic() {
		return string(t.Basic())
	} else if t.IsArray() {
		return t.Array().String()
	} else if t.IsStruct() {
		return t.Struct().String()
	} else if t.IsPointer() {
		return t.Pointer().String() + "*"
	} else {
		// TODO: better error handling
		panic("unknown type")
	}
}

func (t *Type) IRType() (types.Type, error) {
	if t._cached != nil {
		return t._cached, nil
	}

	var final types.Type

	// we must start with the aliased type, because any other type can be an alias type as well
	// TODO: do we need this at all?
	if t.IsAlias() {
		td := t.Scope.FindTypeDefByAlias(t.Alias())

		if td == nil {
			return nil, utils.WithPos(fmt.Errorf("unknown type alias '%s'", t.Alias()), t.Scope.Current().File, t.Pos)
		}

		typ, err := td.Type.IRType()
		if err != nil {
			return nil, err
		}

		final = typ
	} else if t.IsBasic() {
		if t.IsBool() {
			final = NewLLTypeBool()
		}

		if t.IsVoid() {
			final = NewLLTypeVoid()
		}

		if t.IsInt() || t.IsUInt() {
			final = t.LLVMIntType()
		}

		if t.IsFloat() {
			final = t.LLVMFloatType()
		}

		if final == nil {
			return nil, utils.WithPos(fmt.Errorf("unknown basic type '%s'", t.Basic()), t.Scope.Current().File, t.Pos)
		}
	} else if t.IsArray() {
		at := t.Array()

		if at.Len <= 0 {
			return nil, utils.WithPos(fmt.Errorf("array length must be greater than 0"), t.Scope.Current().File, t.Pos)
		}

		et, err := at.Type.IRType()
		if err != nil {
			return nil, err
		}

		final = types.NewArray(uint64(at.Len), et)
	} else if t.IsStruct() {
		// check if we already have a type definition for this struct
		found := t.Scope.FindTypeDefByType(t)

		// if we do (and it's not the current type, in case we found ourselves), then we need to use it
		if found != nil && found.Type != t {
			typ, err := found.Type.IRType()
			if err != nil {
				return nil, err
			}

			final = typ
		} else {
			st := t.Struct()

			typ := types.NewStruct()

			id := t.Scope.CurrentModule().GenerateID("st")
			t.Scope.AddLocalType(id, t)

			typ.SetName(id)
			t.Scope.AddModuleTypeDef(id, t)

			t._cached = typ

			for _, f := range st.Fields {
				ft, err := f.Type.IRType()
				if err != nil {
					return nil, err
				}

				typ.Fields = append(typ.Fields, ft)
			}

			final = typ
		}
	} else if t.IsPointer() {
		typ := t.Pointer()
		if typ.IsVoid() {
			return nil, utils.WithPos(fmt.Errorf("cannot have pointer to void"), t.Scope.Current().File, t.Pos)
		}

		irType, err := typ.IRType()
		if err != nil {
			return nil, err
		}

		final = types.NewPointer(irType)
	} else {
		return nil, utils.WithPos(fmt.Errorf("unknown type"), t.Scope.Current().File, t.Pos)
	}

	t._cached = final

	return t._cached, nil
}

func (t *Type) LLVMIntType() *types.IntType {
	if !t.IsBasic() {
		panic("not a basictype")
	}

	switch t.Basic() {
	case BasicTypeBool:
		return nil
	case BasicTypeVoid:
		return nil
	case BasicTypeI8:
		return NewLLTypeInt(8)
	case BasicTypeI16:
		return NewLLTypeInt(16)
	case BasicTypeI32:
		return NewLLTypeInt(32)
	case BasicTypeI64:
		return NewLLTypeInt(64)
	case BasicTypeUI8:
		return NewLLTypeInt(8)
	case BasicTypeUI16:
		return NewLLTypeInt(16)
	case BasicTypeUI32:
		return NewLLTypeInt(32)
	case BasicTypeUI64:
		return NewLLTypeInt(64)
	case BasicTypeF32:
		fallthrough
	case BasicTypeF64:
		return nil
	default:
		return nil
	}
}

func (t *Type) LLVMFloatType() *types.FloatType {
	if !t.IsBasic() {
		panic("not a basic type")
	}

	switch t.Basic() {
	case BasicTypeBool:
		return nil
	case BasicTypeVoid:
		return nil
	case BasicTypeI8:
		fallthrough
	case BasicTypeI16:
		fallthrough
	case BasicTypeI32:
		fallthrough
	case BasicTypeI64:
		fallthrough
	case BasicTypeUI8:
		fallthrough
	case BasicTypeUI16:
		fallthrough
	case BasicTypeUI32:
		fallthrough
	case BasicTypeUI64:
		return nil
	case BasicTypeF32:
		return NewLLTypeFloat(32)
	case BasicTypeF64:
		return NewLLTypeFloat(64)
	default:
		return nil
	}
}

func (t *Type) BasicSize() int {
	if !t.IsBasic() {
		panic("not a basic type")
	}

	switch t.Basic() {
	case BasicTypeBool:
		return 1
	case BasicTypeVoid:
		return 0
	case BasicTypeI8:
		return 8
	case BasicTypeI16:
		return 16
	case BasicTypeI32:
		return 32
	case BasicTypeI64:
		return 64
	case BasicTypeUI8:
		return 8
	case BasicTypeUI16:
		return 16
	case BasicTypeUI32:
		return 32
	case BasicTypeUI64:
		return 64
	case BasicTypeF32:
		return 32
	case BasicTypeF64:
		return 64
	}

	panic("unknown basic type")
}

func (t *Type) IsBasic() bool {
	return t.Basic() != ""
}

func (t *Type) IsArray() bool {
	return t.Array() != nil
}

func (t *Type) IsStruct() bool {
	return t.Struct() != nil
}

func (t *Type) IsPointer() bool {
	return t.Pointer() != nil
}

func (t *Type) IsAlias() bool {
	return t.Alias() != ""
}

func (t *Type) IsBool() bool {
	return t.IsBasic() && t.Basic() == BasicTypeBool
}

func (t *Type) IsVoid() bool {
	return t.IsBasic() && t.Basic() == BasicTypeVoid
}

func (t *Type) IsInt() bool {
	return t.IsBasic() && (t.Basic() == BasicTypeI8 ||
		t.Basic() == BasicTypeI16 ||
		t.Basic() == BasicTypeI32 ||
		t.Basic() == BasicTypeI64)
}

func (t *Type) IsUInt() bool {
	return t.IsBasic() && (t.Basic() == BasicTypeUI8 ||
		t.Basic() == BasicTypeUI16 ||
		t.Basic() == BasicTypeUI32 ||
		t.Basic() == BasicTypeUI64)
}

func (t *Type) IsFloat() bool {
	return t.IsBasic() && (t.Basic() == BasicTypeF32 ||
		t.Basic() == BasicTypeF64)
}

func (t *Type) PointerLevel() int {
	if t.Pointer() == nil {
		return 0
	}

	return 1 + t.Pointer().PointerLevel()
}

func (t *Type) Equals(o *Type) bool {
	// an alias points to a TypeDef in the module
	if t.IsAlias() {
		typ := t.AliasedType()
		if typ == o {
			return true
		}
	}

	if o.IsAlias() {
		typ := o.AliasedType()
		if typ == t {
			return true
		}
	}

	// if either is alias, then both must be aliases
	if (t.IsAlias() && !o.IsAlias()) || (!t.IsAlias() && o.IsAlias()) {
		return false
	}

	// if both are aliases, then they must be the same alias
	if t.IsAlias() && o.IsAlias() {
		return t.Alias() == o.Alias()
	}

	// here, both are not aliases
	if t.IsBasic() && o.IsBasic() {
		return t.Basic() == o.Basic()
	} else if t.IsArray() && o.IsArray() {
		return t.Array().Equals(o.Array())
	} else if t.IsStruct() && o.IsStruct() {
		return t.Struct().Equals(o.Struct())
	} else if t.IsPointer() && o.IsPointer() {
		return t.Pointer().Equals(o.Pointer())
	} else {
		return false
	}
}

// AliasOf returns true if types are compatible to each other to be used for casting
func (t *Type) AliasOf(o *Type) bool {
	// contrary to Equal, in this case we only care about the final type
	// and the end of the alias chain
	if t.IsBasic() && o.IsBasic() {
		return t.Basic() == o.Basic()
	} else if t.IsStruct() && o.IsStruct() {
		return t.Struct().Equals(o.Struct())
	} else if t.IsPointer() && o.IsPointer() {
		return t.Pointer().Equals(o.Pointer())
	} else {
		return false
	}
}
