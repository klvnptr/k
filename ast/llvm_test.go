package ast_test

import (
	t "testing"

	"github.com/klvnptr/k/ast"
	"github.com/klvnptr/k/testing"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/suite"
)

type LlvmTestSuite struct {
	testing.CompilerSuite
}

func (suite *LlvmTestSuite) SetupTest() {
	suite.T().Parallel()
}

func (suite *LlvmTestSuite) TestGen() {
	i32 := ast.NewLLTypeInt(32)

	m := ir.NewModule()

	format := m.NewGlobalDef("fmt", constant.NewCharArrayFromString("fib(%d) = %d\x00"))

	i8ptr := types.NewPointer(ast.NewLLTypeInt(8))

	m.NewTypeDef("cpd", types.NewStruct(ast.NewLLTypeInt(32), i8ptr))

	printf := m.NewFunc("printf", ast.NewLLTypeInt(32), ir.NewParam("fmt", i8ptr))
	printf.Sig.Variadic = true

	fib := m.NewFunc("fib", i32, ir.NewParam("n", i32))

	// isolate scope
	{
		entry := fib.NewBlock("entry")
		then := fib.NewBlock("then")
		merge := fib.NewBlock("merge")

		// entry
		icmp := entry.NewICmp(enum.IPredSLT, fib.Params[0], constant.NewInt(i32, 2))
		entry.NewCondBr(icmp, then, merge)

		// then
		then.NewRet(fib.Params[0])

		// merge
		pin0 := merge.NewSub(fib.Params[0], constant.NewInt(i32, 1))
		pin1 := merge.NewSub(fib.Params[0], constant.NewInt(i32, 2))
		out0 := merge.NewCall(fib, pin0)
		out1 := merge.NewCall(fib, pin1)
		final := merge.NewAdd(out0, out1)
		merge.NewRet(final)
	}

	main := m.NewFunc("main", i32, ir.NewParam("n", i32))
	{
		entry := main.NewBlock("entry")

		nPtr := entry.NewAlloca(i32)
		entry.NewStore(constant.NewInt(i32, 9), nPtr)
		n := entry.NewLoad(i32, nPtr)

		res := entry.NewCall(fib, n)

		fmtPtr := entry.NewBitCast(format, i8ptr)
		entry.NewCall(printf, fmtPtr, n, res)

		entry.NewRet(constant.NewInt(i32, 0))
	}

	suite.EqualLL(m, "fib(9) = 34")
}

func (suite *LlvmTestSuite) TestAlias() {
	i32 := ast.NewLLTypeInt(32)

	m := ir.NewModule()

	format := m.NewGlobalDef("fmt", constant.NewCharArrayFromString("out: %d\x00"))

	i8ptr := types.NewPointer(ast.NewLLTypeInt(8))

	m.NewTypeDef("cpd", types.NewStruct(ast.NewLLTypeInt(32), i8ptr))

	printf := m.NewFunc("printf", ast.NewLLTypeInt(32), ir.NewParam("fmt", i8ptr))
	printf.Sig.Variadic = true

	main := m.NewFunc("main", i32)
	{
		entry := main.NewBlock("entry")

		res := constant.NewInt(i32, 42)

		fmtPtr := entry.NewBitCast(format, i8ptr)
		entry.NewCall(printf, fmtPtr, res)

		entry.NewRet(constant.NewInt(i32, 0))
	}

	suite.EqualLL(m, "out: 42")
}

func TestLlvmTestSuite(t *t.T) {
	suite.Run(t, new(LlvmTestSuite))
}
