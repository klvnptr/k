//nolint:govet,structtag
package parser

import "github.com/alecthomas/participle/v2/lexer"

// MODULE, FUNCTIONS

type Module struct {
	TypeDefs  []*TypeDef  `@@*`
	Functions []*Function `@@*`

	Pos lexer.Position
}

type TypeDef struct {
	Type  *Type  `"type" @@ `
	Ident string `@Ident ";"`

	Pos lexer.Position
}

type Function struct {
	Declarator  *Declarator   `@@ "("`
	Params      []*Declarator `( @@ ( "," @@ )* )?`
	Variadic    bool          `@( "," "." "." "." )? ")"`
	Body        *CompoundStmt `( @@`
	OnlyDeclare bool          `| @";" )`

	Pos lexer.Position
}

// STATEMENTS

type Stmt struct {
	DeclStmt     *DeclStmt     `@@`
	AssignStmt   *AssignStmt   `| @@`
	ExprStmt     *ExprStmt     `| @@`
	ReturnStmt   *ReturnStmt   `| @@`
	CompoundStmt *CompoundStmt `| @@`
	IfStmt       *IfStmt       `| @@`
	WhileStmt    *WhileStmt    `| @@`

	Pos lexer.Position
}

type ExprStmt struct {
	Expr *Expr `@@ ";"`

	Pos lexer.Position
}

type DeclStmt struct {
	Declarator *Declarator `@@`
	Expr       *Expr       `[ "=" @@ ] ";"`

	Pos lexer.Position
}

type AssignStmt struct {
	Left  *Expr `@@`
	Right *Expr `"=" @@ ";"`

	Pos lexer.Position
}

type ReturnStmt struct {
	Expr *Expr `"return" @@? ";"`

	Pos lexer.Position
}

type CompoundStmt struct {
	Stmts []*Stmt `"{" @@* "}"`

	Pos lexer.Position
}

type IfStmt struct {
	Condition *Expr `"if" "(" @@ ")"`
	Then      *Stmt `@@`
	Else      *Stmt `( "else" @@ )?`

	Pos lexer.Position
}

type WhileStmt struct {
	Condition *Expr `"while" "(" @@ ")"`
	Body      *Stmt `@@`

	Pos lexer.Position
}

// EXPRESSIONS

type Expr struct {
	LogicalExpr *LogicalExpr `@@`

	Pos lexer.Position
}

type LogicalExpr struct {
	Left  *InclusiveOrExpr `@@`
	Op    string           `[ @("&" "&" | "|" "|")`
	Right *LogicalExpr     `@@ ]`

	Pos lexer.Position
}

type InclusiveOrExpr struct {
	Left  *ExclusiveOrExpr `@@`
	Op    string           `[ @("|")`
	Right *InclusiveOrExpr `@@ ]`

	Pos lexer.Position
}

type ExclusiveOrExpr struct {
	Left  *AndExpr         `@@`
	Op    string           `[ @("^")`
	Right *ExclusiveOrExpr `@@ ]`

	Pos lexer.Position
}

type AndExpr struct {
	Left  *EqualityExpr `@@`
	Op    string        `[ @("&")`
	Right *AndExpr      `@@ ]`

	Pos lexer.Position
}

type EqualityExpr struct {
	Left  *ComparisonExpr `@@`
	Op    string          `[ @("=" "=" | "!" "=")`
	Right *EqualityExpr   `@@ ]`

	Pos lexer.Position
}

type ComparisonExpr struct {
	Left  *ShiftExpr      `@@`
	Op    string          `[ @("<" "=" | ">" "=" | "<" | ">")`
	Right *ComparisonExpr `@@ ]`

	Pos lexer.Position
}

type ShiftExpr struct {
	Head *AddExpr `@@`
	Tail []struct {
		Op   string   `@("<" "<" | ">" ">")`
		Expr *AddExpr `@@`

		Pos lexer.Position
	} `@@*`

	Pos lexer.Position
}

// removing left recursion, source: https://github.com/alecthomas/participle/blob/master/_examples/expr3/main.go#L33
type AddExpr struct {
	Head *MulExpr `@@`
	Tail []struct {
		Op   string   `@("+" | "-")`
		Expr *MulExpr `@@`

		Pos lexer.Position
	} `@@*`

	Pos lexer.Position
}

type MulExpr struct {
	Head *CastingExpr `@@`
	Tail []struct {
		Op   string       `@("*" | "/" | "%")`
		Expr *CastingExpr `@@`

		Pos lexer.Position
	} `@@*`

	Pos lexer.Position
}

type CastingExpr struct {
	Type *Type       `[ "(" @@ ")" ]`
	Expr *SizeOfExpr `@@`

	Pos lexer.Position
}

type SizeOfExpr struct {
	Head *PrefixExpr `@@`

	Type *Type       `| "sizeof" "(" @@ ")"`
	Expr *PrefixExpr `| "sizeof" @@`

	Pos lexer.Position
}

type PrefixExpr struct {
	Op   string       `( @( "+" "+" | "-" "-" | "!" | "-" | "*" | "&" )`
	Expr *PrefixExpr  `@@`
	Next *PostfixExpr `| @@ )`

	Pos lexer.Position
}

// PO -> PO o | AE
// removing left recursion because participle doesn't support it: https://cyberzhg.github.io/toolbox/left_rec
// PO -> AE PO'
// PO' -> o PO' | ε
type PostfixExpr struct {
	Next    *AccessorExpr    `@@`
	Postfix *PostfixExprTick `@@`

	Pos lexer.Position
}

type PostfixExprTick struct {
	Op   string           `[ @("+" "+" | "-" "-")`
	Expr *PostfixExprTick `@@ ]`

	Pos lexer.Position
}

type AccessorExpr struct {
	Head *IndexExpr `@@`
	Tail []struct {
		Op    string `@("-" ">" | ".")`
		Field string `@Ident`

		Pos lexer.Position
	} `@@*`

	Pos lexer.Position
}

type IndexExpr struct {
	Head *UnaryExpr `@@`
	Tail []struct {
		Index *Expr `'[' @@ ']'`

		Pos lexer.Position
	} `@@*`

	Pos lexer.Position
}

type UnaryExpr struct {
	// SizeOfExpr  *SizeOfExpr  `@@`
	FnCallExpr  *FnCallExpr  `@@`
	PrimaryExpr *PrimaryExpr `| @@`

	Pos lexer.Position
}

type FnCallExpr struct {
	Ident string  `@Ident`
	Args  []*Expr `"(" (@@ ("," @@)*)? ")"`

	Pos lexer.Position
}

type StructField struct {
	Field string `"." @Ident`
	Expr  *Expr  `"=" @@`
}

type StructExpr struct {
	Alias        string         `@Ident`
	StructFields []*StructField `"{" @@ ( "," @@ )* "," "}"`
}

type PrimaryExpr struct {
	// StructExpr must be before Ident, because StructExpr starts with Ident
	Struct   *StructExpr `@@`
	Variable string      `| @Ident`
	Sign     string      `| @("+" | "-")?`
	Number   string      `@Number`
	Char     string      `| ( CharStart @SingleChar CharEnd )`
	// we need to use a struct here, because we want to allow empty strings
	String *InternalString `| ( StringStart @@ StringEnd )`
	Expr   *Expr           `| "(" @@ ")"`

	Pos lexer.Position
}

type InternalString struct {
	Parts []string `( @Escaped | @Chars )*`
}

// TYPES

type Declarator struct {
	Type  *Type  `@@`
	Ident string `@Ident`

	Pos lexer.Position
}

type Struct struct {
	Fields []*Declarator `"struct" "{" @@ ( "," @@ )* "," "}"`
}

type Type struct {
	Lengths []int `( "[" @Number "]" )*`

	Struct *Struct `( @@`
	// must match lexer.go BasicType AND ast.Type
	Basic string `| @("bool" | "void" | "i8" | "i16" | "i32" | "i64" | "u8" | "u16" | "u32" | "u64" | "f32" | "f64")`
	Alias string `| @Ident )`

	Pointers string `@"*"*`

	Pos lexer.Position
}
