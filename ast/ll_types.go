package ast

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func NewLLTypeVoid() *types.VoidType {
	return &types.VoidType{}
}

func NewLLTypeBool() *types.IntType {
	return &types.IntType{
		BitSize: 1,
	}
}

func NewLLTypeInt(bitSize uint64) *types.IntType {
	return &types.IntType{
		BitSize: bitSize,
	}
}

func NewLLTypeFloat(bitSize uint64) *types.FloatType {
	if bitSize == 32 {
		return &types.FloatType{
			Kind: types.FloatKindFloat,
		}
	} else if bitSize == 64 {
		return &types.FloatType{
			Kind: types.FloatKindDouble,
		}
	} else {
		panic("invalid bit size")
	}
}

func NewLLBool(value bool) *constant.Int {
	if value {
		return constant.NewInt(NewLLTypeBool(), 1)
	} else {
		return constant.NewInt(NewLLTypeBool(), 0)
	}
}

func NewLLInt(bitSize uint64, value int) *constant.Int {
	// TODO: check bitSize
	return constant.NewInt(NewLLTypeInt(bitSize), int64(value))
}

func NewLLFloat(bitSize uint64, value float64) *constant.Float {
	return constant.NewFloat(NewLLTypeFloat(bitSize), value)
}
