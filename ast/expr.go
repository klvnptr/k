package ast

import (
	"fmt"
	"strconv"

	"github.com/klvnptr/k/utils"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (b *BinaryOp) String() string {
	return fmt.Sprintf("%s %s %s", b.Left.String(), b.Op, b.Right.String())
}

func (b *BinaryOp) Value() (*Value, error) {
	left, err := b.Left.Value()
	if err != nil {
		return nil, err
	}

	right, err := b.Right.Value()
	if err != nil {
		return nil, err
	}

	bb := b.Scope.BasicBlock()
	var result value.Value

	// if eithe left or right is a pointer, we need to do some pointer arithmetic
	if left.Type.IsPointer() && !right.Type.IsPointer() && right.Type.IsInt() {
		ptrIRType, err := left.Type.Pointer().IRType()
		if err != nil {
			return nil, err
		}

		switch b.Op {
		case "+":
			result = bb.NewGetElementPtr(ptrIRType, left.Value, right.Value)
		case "-":
			result = bb.NewGetElementPtr(ptrIRType, left.Value, right.Value)
		case "==":
			ptrInt := bb.NewPtrToInt(left.Value, types.I64)
			result = bb.NewICmp(enum.IPredEQ, ptrInt, right.Value)
		case "!=":
			ptrInt := bb.NewPtrToInt(left.Value, types.I64)
			result = bb.NewICmp(enum.IPredNE, ptrInt, right.Value)
		}
	} else if !left.Type.IsPointer() && right.Type.IsPointer() && left.Type.IsInt() {
		ptrIRType, err := right.Type.Pointer().IRType()
		if err != nil {
			return nil, err
		}

		switch b.Op {
		case "+":
			result = bb.NewGetElementPtr(ptrIRType, right.Value, left.Value)
		case "-":
			result = bb.NewGetElementPtr(ptrIRType, right.Value, left.Value)
		case "==":
			ptrInt := bb.NewPtrToInt(right.Value, types.I64)
			result = bb.NewICmp(enum.IPredEQ, left.Value, ptrInt)
		case "!=":
			ptrInt := bb.NewPtrToInt(right.Value, types.I64)
			result = bb.NewICmp(enum.IPredNE, left.Value, ptrInt)
		}
	} else if !left.Type.Equals(right.Type) {
		return nil, utils.WithPos(fmt.Errorf("incompatible types %s and %s", left.Type.String(), right.Type.String()), b.Scope.Current().File, b.Pos)
	} else if left.Type.IsBool() {
		switch b.Op {
		case "&&":
			result = bb.NewAnd(left.Value, right.Value)
		case "||":
			result = bb.NewOr(left.Value, right.Value)
		case "==":
			result = bb.NewICmp(enum.IPredEQ, left.Value, right.Value)
		case "!=":
			result = bb.NewICmp(enum.IPredNE, left.Value, right.Value)
		}
	} else if left.Type.IsInt() {
		switch b.Op {
		case "|":
			result = bb.NewOr(left.Value, right.Value)
		case "&":
			result = bb.NewAnd(left.Value, right.Value)
		case "^":
			result = bb.NewXor(left.Value, right.Value)
		case "==":
			result = bb.NewICmp(enum.IPredEQ, left.Value, right.Value)
		case "!=":
			result = bb.NewICmp(enum.IPredNE, left.Value, right.Value)
		case "<":
			result = bb.NewICmp(enum.IPredSLT, left.Value, right.Value)
		case ">":
			result = bb.NewICmp(enum.IPredSGT, left.Value, right.Value)
		case "<=":
			result = bb.NewICmp(enum.IPredSLE, left.Value, right.Value)
		case ">=":
			result = bb.NewICmp(enum.IPredSGE, left.Value, right.Value)
		case "+":
			result = bb.NewAdd(left.Value, right.Value)
			result.(*ir.InstAdd).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
		case "-":
			result = bb.NewSub(left.Value, right.Value)
			result.(*ir.InstSub).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
		case "*":
			result = bb.NewMul(left.Value, right.Value)
			result.(*ir.InstMul).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
		case "/":
			result = bb.NewSDiv(left.Value, right.Value)
		case "%":
			result = bb.NewSRem(left.Value, right.Value)
		}
	} else if left.Type.IsUInt() {
		switch b.Op {
		case "|":
			result = bb.NewOr(left.Value, right.Value)
		case "&":
			result = bb.NewAnd(left.Value, right.Value)
		case "^":
			result = bb.NewXor(left.Value, right.Value)
		case "==":
			result = bb.NewICmp(enum.IPredEQ, left.Value, right.Value)
		case "!=":
			result = bb.NewICmp(enum.IPredNE, left.Value, right.Value)
		case "<":
			result = bb.NewICmp(enum.IPredULT, left.Value, right.Value)
		case ">":
			result = bb.NewICmp(enum.IPredUGT, left.Value, right.Value)
		case "<=":
			result = bb.NewICmp(enum.IPredULE, left.Value, right.Value)
		case ">=":
			result = bb.NewICmp(enum.IPredUGE, left.Value, right.Value)
		case "+":
			result = bb.NewAdd(left.Value, right.Value)
			result.(*ir.InstAdd).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
		case "-":
			result = bb.NewSub(left.Value, right.Value)
			result.(*ir.InstSub).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
		case "*":
			result = bb.NewMul(left.Value, right.Value)
			result.(*ir.InstMul).OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
		case "/":
			result = bb.NewUDiv(left.Value, right.Value)
		case "%":
			result = bb.NewURem(left.Value, right.Value)
		}
	} else if left.Type.IsFloat() {
		switch b.Op {
		case "==":
			result = bb.NewFCmp(enum.FPredOEQ, left.Value, right.Value)
		case "!=":
			result = bb.NewFCmp(enum.FPredONE, left.Value, right.Value)
		case "<":
			result = bb.NewFCmp(enum.FPredOLT, left.Value, right.Value)
		case ">":
			result = bb.NewFCmp(enum.FPredOGT, left.Value, right.Value)
		case "<=":
			result = bb.NewFCmp(enum.FPredOLE, left.Value, right.Value)
		case ">=":
			result = bb.NewFCmp(enum.FPredOGE, left.Value, right.Value)
		case "+":
			result = bb.NewFAdd(left.Value, right.Value)
		case "-":
			result = bb.NewFSub(left.Value, right.Value)
		case "*":
			result = bb.NewFMul(left.Value, right.Value)
		case "/":
			result = bb.NewFDiv(left.Value, right.Value)
		case "%":
			result = bb.NewFRem(left.Value, right.Value)
		}
	} else if left.Type.IsPointer() {
		switch b.Op {
		case "==":
			result = bb.NewICmp(enum.IPredEQ, left.Value, right.Value)
		case "!=":
			result = bb.NewICmp(enum.IPredNE, left.Value, right.Value)
		case "<":
			result = bb.NewICmp(enum.IPredULT, left.Value, right.Value)
		case ">":
			result = bb.NewICmp(enum.IPredUGT, left.Value, right.Value)
		case "<=":
			result = bb.NewICmp(enum.IPredULE, left.Value, right.Value)
		case ">=":
			result = bb.NewICmp(enum.IPredUGE, left.Value, right.Value)
		}
	}

	if result == nil {
		return nil, utils.WithPos(fmt.Errorf("operation %s is not implemented for %s", b.Op, left.Type), b.Scope.Current().File, b.Pos)
	}

	// if result is a boolean, we can't use left's type, we need to use a bool
	if result.Type().Equal(NewLLTypeBool()) {
		return &Value{
			Type:  NewTypeBasic(b.Scope, b.Pos, BasicTypeBool),
			Value: result,
		}, nil
	}

	return &Value{
		Type:  left.Type,
		Value: result,
	}, nil
}

func (u *UnaryOp) String() string {
	return fmt.Sprintf("%s %s", u.Op, u.Expr.String())
}

func (u *UnaryOp) Value() (*Value, error) {
	original, err := u.Expr.Value()

	if err != nil {
		return nil, err
	}

	var result value.Value

	irType, err := original.Type.IRType()
	if err != nil {
		return nil, err
	}

	switch u.Op {
	case "--", "++":
		if original.Ptr == nil {
			return nil, utils.WithPos(fmt.Errorf("cannot increment/decrement a value that's not stored in memory"), u.Scope.Current().File, u.Pos)
		}

		// load the value, just in case in has been modified
		original.Value = u.Scope.BasicBlock().NewLoad(irType, original.Ptr)

		if original.Type.IsPointer() {
			ptrIRType, err := original.Type.Pointer().IRType()
			if err != nil {
				return nil, err
			}

			if u.Op == "++" {
				result = u.Scope.BasicBlock().NewGetElementPtr(ptrIRType, original.Value, NewLLInt(32, 1))
			} else {
				result = u.Scope.BasicBlock().NewGetElementPtr(ptrIRType, original.Value, NewLLInt(32, -1))
			}
		} else if original.Type.IsInt() {
			rhs := constant.NewInt(original.Type.LLVMIntType(), 1)
			if u.Op == "++" {
				inst := u.Scope.BasicBlock().NewAdd(original.Value, rhs)
				inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
				result = inst
			} else {
				inst := u.Scope.BasicBlock().NewSub(original.Value, rhs)
				inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
				result = inst
			}
		} else if original.Type.IsUInt() {
			rhs := constant.NewInt(original.Type.LLVMIntType(), 1)
			if u.Op == "++" {
				inst := u.Scope.BasicBlock().NewAdd(original.Value, rhs)
				inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
				result = inst
			} else {
				inst := u.Scope.BasicBlock().NewSub(original.Value, rhs)
				inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
				result = inst
			}
		} else if original.Type.IsFloat() {
			rhs := constant.NewFloat(original.Type.LLVMFloatType(), 1)
			if u.Op == "++" {
				result = u.Scope.BasicBlock().NewFAdd(original.Value, rhs)
			} else {
				result = u.Scope.BasicBlock().NewFSub(original.Value, rhs)
			}
		} else {
			return nil, utils.WithPos(fmt.Errorf("cannot increment/decrement a %s", original.Type.String()), u.Scope.Current().File, u.Pos)
		}

		u.Scope.BasicBlock().NewStore(result, original.Ptr)

		if u.IsPostfix {
			return original, nil
		} else {
			return &Value{Type: original.Type, Ptr: original.Ptr, Value: result}, nil
		}
	case "-":
		if original.Type.IsPointer() {
			return nil, utils.WithPos(fmt.Errorf("cannot negate a pointer"), u.Scope.Current().File, u.Pos)
		} else if original.Type.IsInt() {
			inst := u.Scope.BasicBlock().NewSub(constant.NewInt(original.Type.LLVMIntType(), 0), original.Value)
			inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNSW}
			result = inst
		} else if original.Type.IsUInt() {
			inst := u.Scope.BasicBlock().NewSub(constant.NewInt(original.Type.LLVMIntType(), 0), original.Value)
			inst.OverflowFlags = []enum.OverflowFlag{enum.OverflowFlagNUW}
			result = inst
		} else if original.Type.IsFloat() {
			result = u.Scope.BasicBlock().NewFNeg(original.Value)
		} else {
			return nil, utils.WithPos(fmt.Errorf("cannot negate a %s", original.Type.String()), u.Scope.Current().File, u.Pos)
		}

		return &Value{Type: original.Type, Value: result}, nil
	case "!":
		if original.Type.IsInt() || original.Type.IsUInt() {
			result = u.Scope.BasicBlock().NewICmp(enum.IPredEQ, original.Value, constant.NewInt(original.Type.LLVMIntType(), 0))
		} else if original.Type.IsBool() {
			result = u.Scope.BasicBlock().NewICmp(enum.IPredEQ, original.Value, NewLLBool(false))
		} else {
			return nil, utils.WithPos(fmt.Errorf("cannot negate a %s", original.Type.String()), u.Scope.Current().File, u.Pos)
		}

		return &Value{Type: original.Type, Value: result}, nil
	case "&":
		if original.Ptr == nil {
			return nil, utils.WithPos(fmt.Errorf("cannot take the address of a non-variable"), u.Scope.Current().File, u.Pos)
		}

		return &Value{Type: original.Type.NewPointer(), Value: original.Ptr}, nil
	case "*":
		if !original.Type.IsPointer() {
			return nil, utils.WithPos(fmt.Errorf("cannot dereference a non-pointer type"), u.Scope.Current().File, u.Pos)
		}

		ptr := original.Type.Pointer()

		ptrIRType, err := ptr.IRType()
		if err != nil {
			return nil, err
		}

		return &Value{
			Type:  ptr,
			Ptr:   original.Value,
			Value: u.Scope.BasicBlock().NewLoad(ptrIRType, original.Value),
		}, nil
	default:
		return nil, utils.WithPos(fmt.Errorf("unary operator %s not implemented for type %s", u.Op, original.Type.String()), u.Scope.Current().File, u.Pos)
	}
}

func (ao *AccessorOp) String() string {
	return fmt.Sprintf("%s.%s", ao.Expr.String(), ao.Field)
}

func (ao *AccessorOp) Value() (*Value, error) {
	expr, err := ao.Expr.Value()

	if err != nil {
		return nil, err
	}

	if ao.Dereference {
		if !expr.Type.IsPointer() {
			return nil, utils.WithPos(fmt.Errorf("cannot dereference a non-pointer type %s", expr.Type.String()), ao.Scope.Current().File, ao.Pos)
		}

		ptr := expr.Type.Pointer()

		ptrIRType, err := ptr.IRType()
		if err != nil {
			return nil, err
		}

		expr = &Value{
			Type:  ptr,
			Ptr:   expr.Value,
			Value: ao.Scope.BasicBlock().NewLoad(ptrIRType, expr.Value),
		}
	}

	if !expr.Type.IsStruct() {
		return nil, utils.WithPos(fmt.Errorf("cannot access field of non-struct type %s", expr.Type.String()), ao.Scope.Current().File, ao.Pos)
	}

	if expr.Ptr == nil {
		return nil, utils.WithPos(fmt.Errorf("cannot access field of non-variable"), ao.Scope.Current().File, ao.Pos)
	}

	index, field, err := expr.Type.Struct().FindField(ao.Field)
	if err != nil {
		return nil, utils.WithPos(err, ao.Scope.Current().File, ao.Pos)
	}

	exprIRType, err := expr.Type.IRType()
	if err != nil {
		return nil, err
	}

	ptr := ao.Scope.BasicBlock().NewGetElementPtr(exprIRType, expr.Ptr, NewLLInt(32, 0), NewLLInt(32, index))

	fieldIRType, err := field.Type.IRType()
	if err != nil {
		return nil, err
	}

	loaded := ao.Scope.BasicBlock().NewLoad(fieldIRType, ptr)

	return &Value{
		Type:  field.Type,
		Ptr:   ptr,
		Value: loaded,
	}, nil
}

func (io *IndexOp) String() string {
	return fmt.Sprintf("%s[%s]", io.Expr.String(), io.IndexExpr.String())
}

func (io *IndexOp) Value() (*Value, error) {
	expr, err := io.Expr.Value()

	if err != nil {
		return nil, err
	}

	indexExpr, err := io.IndexExpr.Value()
	if err != nil {
		return nil, err
	}

	// TODO: can we index with uint?
	if !indexExpr.Type.IsInt() || indexExpr.Type.IsPointer() {
		return nil, utils.WithPos(fmt.Errorf("index must be an integer"), io.Scope.Current().File, io.Pos)
	}

	bb := io.Scope.BasicBlock()

	if expr.Type.IsPointer() {
		ptr := expr.Type.Pointer()

		ptrIRType, err := ptr.IRType()
		if err != nil {
			return nil, err
		}

		addr := bb.NewGetElementPtr(ptrIRType, expr.Value, indexExpr.Value)
		result := bb.NewLoad(ptrIRType, addr)

		return &Value{
			Type:  ptr,
			Ptr:   addr,
			Value: result,
		}, nil
	} else if expr.Type.IsArray() {
		arr := expr.Type.Array()

		if expr.Ptr == nil {
			return nil, utils.WithPos(fmt.Errorf("cannot index a non-variable"), io.Scope.Current().File, io.Pos)
		}

		arrayIRType, err := expr.Type.IRType()
		if err != nil {
			return nil, err
		}

		elemType := arr.Type

		elemIRType, err := elemType.IRType()
		if err != nil {
			return nil, err
		}

		addr := bb.NewGetElementPtr(arrayIRType, expr.Ptr, NewLLInt(32, 0), indexExpr.Value)
		result := bb.NewLoad(elemIRType, addr)

		return &Value{
			Type:  elemType,
			Ptr:   addr,
			Value: result,
		}, nil
	} else {
		return nil, utils.WithPos(fmt.Errorf("can only index pointers or arrays, but type is %s", expr.Type.String()), io.Scope.Current().File, io.Pos)
	}
}

func (c *CastingOp) String() string {
	return "(" + c.Type.String() + ")" + c.Expr.String()
}

func (c *CastingOp) Value() (*Value, error) {
	expr, err := c.Expr.Value()
	if err != nil {
		return nil, err
	}

	if expr.Type.Equals(c.Type) {
		return &Value{Type: c.Type, Ptr: expr.Ptr, Value: expr.Value}, nil
	}

	if expr.Type.AliasOf(c.Type) {
		return &Value{Type: c.Type, Value: expr.Value}, nil
	}

	var result value.Value

	targetIRType, err := c.Type.IRType()
	if err != nil {
		return nil, err
	}

	if c.Type.IsInt() && expr.Type.IsInt() {
		if c.Type.BasicSize() > expr.Type.BasicSize() {
			result = c.Scope.BasicBlock().NewSExt(expr.Value, targetIRType)
		} else {
			result = c.Scope.BasicBlock().NewTrunc(expr.Value, targetIRType)
		}
	} else if c.Type.IsUInt() && expr.Type.IsUInt() {
		if c.Type.BasicSize() > expr.Type.BasicSize() {
			result = c.Scope.BasicBlock().NewZExt(expr.Value, targetIRType)
		} else {
			result = c.Scope.BasicBlock().NewTrunc(expr.Value, targetIRType)
		}
	} else if c.Type.IsFloat() && expr.Type.IsFloat() {
		if c.Type.BasicSize() > expr.Type.BasicSize() {
			result = c.Scope.BasicBlock().NewFPExt(expr.Value, targetIRType)
		} else {
			result = c.Scope.BasicBlock().NewFPTrunc(expr.Value, targetIRType)
		}
	} else if c.Type.IsPointer() {
		if expr.Type.IsPointer() {
			result = c.Scope.BasicBlock().NewBitCast(expr.Value, targetIRType)
		} else if expr.Type.IsArray() {
			elemIRType, err := expr.Type.Array().Type.IRType()
			if err != nil {
				return nil, err
			}
			result = c.Scope.BasicBlock().NewBitCast(expr.Value, types.NewPointer(elemIRType))
		} else {
			return nil, utils.WithPos(fmt.Errorf("casting from %s to %s not implemented", expr.Type.String(), c.Type.String()), c.Scope.Current().File, c.Pos)
		}
	} else {
		return nil, utils.WithPos(fmt.Errorf("casting from %s to %s not implemented", expr.Type.String(), c.Type.String()), c.Scope.Current().File, c.Pos)
	}

	return &Value{
		Type:  c.Type,
		Value: result,
	}, nil

}

func (f *FnCallOp) String() string {
	return fmt.Sprintf("%s(%s)", f.Ident, ExpressionLikeList(f.Args).String())
}

func (f *FnCallOp) Value() (*Value, error) {
	fn := f.Scope.FindFunction(f.Ident)

	if fn == nil {
		return nil, utils.WithPos(fmt.Errorf("function %s not found", f.Ident), f.Scope.Current().File, f.Pos)
	}

	values := []value.Value{}
	// TODO: check if the number of arguments is correct
	// TODO: check if the types of the arguments are correct
	for _, arg := range f.Args {
		v, err := arg.Value()

		if err != nil {
			return nil, err
		}

		values = append(values, v.Value)
	}

	ret := f.Scope.BasicBlock().NewCall(fn.Ptr, values...)

	return &Value{
		Type:  fn.ReturnType,
		Value: ret,
	}, nil
}

func (s *SizeOfOp) String() string {
	return "sizeof(" + s.Type.String() + ")"
}

func (s *SizeOfOp) Value() (*Value, error) {
	bb := s.Scope.BasicBlock()

	irType, err := s.Type.IRType()
	if err != nil {
		return nil, err
	}

	// source: https://stackoverflow.com/questions/14608250/how-can-i-find-the-size-of-a-type
	size := bb.NewGetElementPtr(
		irType,
		constant.NewNull(types.NewPointer(irType)),
		NewLLInt(32, 1),
	)
	sizeInt := bb.NewPtrToInt(size, NewLLTypeInt(64))

	return &Value{
		Type:  NewTypeBasic(s.Scope, s.Pos, BasicTypeI64),
		Value: sizeInt,
	}, nil
}

func (l *LoadOp) String() string {
	return "load(" + l.Name + ")"
}

func (l *LoadOp) Value() (*Value, error) {
	v := l.Scope.FindVariable(l.Name)
	if v == nil {
		return nil, utils.WithPos(fmt.Errorf("variable %s not found", l.Name), l.Scope.Current().File, l.Pos)
	}

	irType, err := v.Type.IRType()
	if err != nil {
		return nil, err
	}

	inst := l.Scope.BasicBlock().NewLoad(irType, v.Ptr)

	return &Value{
		Type:  v.Type,
		Ptr:   v.Ptr,
		Value: inst,
	}, nil
}

func (c *ConstantBoolOp) String() string {
	return fmt.Sprintf("%s", c.Constant)
}

func (c *ConstantBoolOp) Value() (*Value, error) {
	value := false

	if c.Constant == "true" {
		value = true
	}

	return &Value{
		Type:  NewTypeBasic(c.Scope, c.Pos, BasicTypeBool),
		Value: constant.NewBool(value),
	}, nil
}

func (c *ConstantNumberOp) String() string {
	return fmt.Sprintf("%s%s", c.Sign, c.Constant)
}

func (c *ConstantNumberOp) Value() (*Value, error) {
	negative := false
	if c.Sign == "-" {
		negative = true
	}

	// try to parse it as an integer first
	i, err := strconv.ParseInt(c.Constant, 10, 64)

	if err != nil {
		// if failed, try to parse it as a float
		f, err := strconv.ParseFloat(c.Constant, 64)

		if err != nil {
			return nil, utils.WithPos(fmt.Errorf("can't parse number"), c.Scope.Current().File, c.Pos)
		}

		if negative {
			f = -f
		}

		return &Value{
			Type:  NewTypeBasic(c.Scope, c.Pos, BasicTypeF64),
			Value: NewLLFloat(64, f),
		}, nil
	}

	if negative {
		i = -i
	}

	return &Value{
		Type:  NewTypeBasic(c.Scope, c.Pos, BasicTypeI64),
		Value: NewLLInt(64, int(i)),
	}, nil
}

func (c *ConstantCharOp) String() string {
	return `'` + c.Constant + `'`
}

func (c *ConstantCharOp) Value() (*Value, error) {
	return &Value{
		Type:  NewTypeBasic(c.Scope, c.Pos, BasicTypeI8),
		Value: NewLLInt(8, int(c.Constant[0])),
	}, nil
}

func (c *ConstantStringOp) String() string {
	return `"` + c.Constant + `"`
}

func (c *ConstantStringOp) Value() (*Value, error) {
	m := c.Scope.CurrentModule()
	id := m.GenerateID("id")
	ptr := m.Ptr.NewGlobalDef(id, constant.NewCharArrayFromString(c.Constant+"\x00"))

	return &Value{
		Type:  NewTypeBasic(c.Scope, c.Pos, BasicTypeI8).NewPointer(),
		Value: c.Scope.BasicBlock().NewBitCast(ptr, types.NewPointer(NewLLTypeInt(8))),
	}, nil
}
