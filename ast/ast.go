package ast

import (
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// MODULE, FUNCTIONS

type Module struct {
	Functions []*Function

	ModuleTypeDefs []*TypeDef
	LocalTypes     []*TypeDef
	Globals        []*Variable

	Ptr          *ir.Module
	LastID       int
	CurrentBlock *ir.Block

	Scope *Scope
	Pos   lexer.Position
}

type TypeDef struct {
	Alias string
	Type  *Type
}

type Function struct {
	Name        string
	Params      []*Variable
	Variadic    bool
	ReturnType  *Type
	Body        []StatementLike
	OnlyDeclare bool

	Locals []*Variable

	Ptr *ir.Func

	Scope *Scope
	Pos   lexer.Position
}

// STATEMENTS

type StatementLike interface {
	String() []string
	Generate() error
}

type ExprStmt struct {
	Expr ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type DeclStmt struct {
	Ident string
	Type  *Type
	Expr  ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type AssignStmt struct {
	Type  *Type
	Left  ExpressionLike
	Right ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type ReturnStmt struct {
	Expr ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type Block struct {
	Stmts []StatementLike

	Locals []*Variable

	Scope *Scope
	Pos   lexer.Position
}

type IfStmt struct {
	Condition ExpressionLike
	Then      []StatementLike
	Else      []StatementLike

	Scope ScopeLike
	Pos   lexer.Position
}

type WhileStmt struct {
	Condition ExpressionLike
	Body      []StatementLike

	Scope ScopeLike
	Pos   lexer.Position
}

// EXPRESSIONS

type Value struct {
	Type  *Type
	Ptr   value.Value
	Value value.Value
}

type ExpressionLike interface {
	String() string
	Value() (*Value, error)
}

type BinaryOp struct {
	Left  ExpressionLike
	Op    string
	Right ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type UnaryOp struct {
	Op        string
	Expr      ExpressionLike
	IsPostfix bool

	Scope ScopeLike
	Pos   lexer.Position
}

type AccessorOp struct {
	Expr        ExpressionLike
	Field       string
	Dereference bool

	Scope ScopeLike
	Pos   lexer.Position
}

type IndexOp struct {
	Op        string
	Expr      ExpressionLike
	IndexExpr ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type CastingOp struct {
	Type *Type
	Expr ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type FnCallOp struct {
	Ident string
	Args  []ExpressionLike

	Scope ScopeLike
	Pos   lexer.Position
}

type SizeOfOp struct {
	Type *Type

	Scope ScopeLike
	Pos   lexer.Position
}

type LoadOp struct {
	Name string

	Scope ScopeLike
	Pos   lexer.Position
}

type ConstantBoolOp struct {
	Constant string

	Scope ScopeLike
	Pos   lexer.Position
}

type ConstantNumberOp struct {
	Sign     string
	Constant string

	Scope ScopeLike
	Pos   lexer.Position
}

type ConstantCharOp struct {
	Constant string

	Scope ScopeLike
	Pos   lexer.Position
}

type ConstantStringOp struct {
	Constant string

	Scope ScopeLike
	Pos   lexer.Position
}
