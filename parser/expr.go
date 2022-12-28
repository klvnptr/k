package parser

import (
	"strings"

	"github.com/klvnptr/k/ast"
)

func (e *Expr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	return e.LogicalExpr.Transform(scope)
}

func (le *LogicalExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if le.Op == "" {
		return le.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  le.Left.Transform(scope),
		Op:    le.Op,
		Right: le.Right.Transform(scope),
		Scope: scope,
		Pos:   le.Pos,
	}
}

func (io *InclusiveOrExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if io.Op == "" {
		return io.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  io.Left.Transform(scope),
		Op:    io.Op,
		Right: io.Right.Transform(scope),
		Scope: scope,
		Pos:   io.Pos,
	}
}

func (ea *ExclusiveOrExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if ea.Op == "" {
		return ea.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  ea.Left.Transform(scope),
		Op:    ea.Op,
		Right: ea.Right.Transform(scope),
		Scope: scope,
		Pos:   ea.Pos,
	}
}

func (ae *AndExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if ae.Op == "" {
		return ae.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  ae.Left.Transform(scope),
		Op:    ae.Op,
		Right: ae.Right.Transform(scope),
		Scope: scope,
		Pos:   ae.Pos,
	}
}

func (ee *EqualityExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if ee.Op == "" {
		return ee.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  ee.Left.Transform(scope),
		Op:    ee.Op,
		Right: ee.Right.Transform(scope),
		Scope: scope,
		Pos:   ee.Pos,
	}
}

func (ce *ComparisonExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if ce.Op == "" {
		return ce.Left.Transform(scope)
	}

	return &ast.BinaryOp{
		Left:  ce.Left.Transform(scope),
		Op:    ce.Op,
		Right: ce.Right.Transform(scope),
		Scope: scope,
		Pos:   ce.Pos,
	}
}

func (ae *AddExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	head := ae.Head.Transform(scope)

	for _, tail := range ae.Tail {
		head = &ast.BinaryOp{
			Left:  head,
			Op:    tail.Op,
			Right: tail.Expr.Transform(scope),
			Scope: scope,
			Pos:   tail.Pos,
		}
	}

	return head
}

func (me *MulExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	head := me.Head.Transform(scope)

	for _, tail := range me.Tail {
		head = &ast.BinaryOp{
			Left:  head,
			Op:    tail.Op,
			Right: tail.Expr.Transform(scope),
			Scope: scope,
			Pos:   tail.Pos,
		}
	}

	return head
}

func (ce *CastingExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if ce.Type == nil {
		return ce.Expr.Transform(scope)
	}

	return &ast.CastingOp{
		Type:  ce.Type.Transform(scope),
		Expr:  ce.Expr.Transform(scope),
		Scope: scope,
		Pos:   ce.Pos,
	}
}

func (pe *PrefixExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	if pe.Next != nil {
		return pe.Next.Transform(scope)
	}

	return &ast.UnaryOp{
		Op:    pe.Op,
		Expr:  pe.Expr.Transform(scope),
		Scope: scope,
		Pos:   pe.Pos,
	}
}

func (pe *PostfixExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	next := pe.Next.Transform(scope)
	return pe.Postfix.Transform(scope, next)
}

func (pet *PostfixExprTick) Transform(scope ast.ScopeLike, ae ast.ExpressionLike) ast.ExpressionLike {
	if pet.Op != "" {
		return &ast.UnaryOp{
			Op:        pet.Op,
			Expr:      pet.Expr.Transform(scope, ae),
			IsPostfix: true,
			Scope:     scope,
			Pos:       pet.Pos,
		}
	}

	return ae
}

func (ae *AccessorExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	head := ae.Head.Transform(scope)

	for _, tail := range ae.Tail {
		deref := false
		if tail.Op == "->" {
			deref = true
		}

		head = &ast.AccessorOp{
			Expr:        head,
			Field:       tail.Field,
			Dereference: deref,
			Scope:       scope,
			Pos:         tail.Pos,
		}
	}

	return head
}

func (ie *IndexExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	expr := ie.Expr.Transform(scope)

	if ie.Index != nil {
		return &ast.IndexOp{
			Expr:      expr,
			IndexExpr: ie.Index.Transform(scope),
			Scope:     scope,
			Pos:       ie.Pos,
		}
	}

	return expr
}

func (ue *UnaryExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	switch {
	case ue.SizeOfExpr != nil:
		return ue.SizeOfExpr.Transform(scope)
	case ue.FnCallExpr != nil:
		return ue.FnCallExpr.Transform(scope)
	case ue.PrimaryExpr != nil:
		return ue.PrimaryExpr.Transform(scope)
	default:
		panic("unknown unary expression")
	}
}

func (soe *SizeOfExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	return &ast.SizeOfOp{
		Type:  soe.Type.Transform(scope),
		Scope: scope,
		Pos:   soe.Pos,
	}
}

func (fce *FnCallExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	args := make([]ast.ExpressionLike, len(fce.Args))

	for i, arg := range fce.Args {
		args[i] = arg.Transform(scope)
	}

	return &ast.FnCallOp{
		Ident: fce.Ident,
		Args:  args,
		Scope: scope,
		Pos:   fce.Pos,
	}
}

func (pe *PrimaryExpr) Transform(scope ast.ScopeLike) ast.ExpressionLike {
	switch {
	case pe.Variable != "":
		if pe.Variable == "true" || pe.Variable == "false" {
			return &ast.ConstantBoolOp{
				Constant: pe.Variable,
				Scope:    scope,
				Pos:      pe.Pos,
			}
		}

		return &ast.LoadOp{
			Name:  pe.Variable,
			Scope: scope,
			Pos:   pe.Pos,
		}
	case pe.Char != "":
		return &ast.ConstantCharOp{
			Constant: pe.Char,
			Scope:    scope,
			Pos:      pe.Pos,
		}
	case pe.Number != "":
		return &ast.ConstantNumberOp{
			Sign:     pe.Sign,
			Constant: pe.Number,
			Scope:    scope,
			Pos:      pe.Pos,
		}
	case pe.String != nil:
		str := ""
		for _, el := range pe.String.Parts {
			// TODO: should handle more escape sequences
			if strings.HasPrefix(el, "\\") {
				switch el[1:] {
				case "n":
					str += "\n"
				case "r":
					str += "\r"
				case "t":
					str += "\t"
				case `"`:
					str += `"`
				default:
					str += el
				}
			} else {
				str += el
			}
		}

		return &ast.ConstantStringOp{
			Constant: str,
			Scope:    scope,
			Pos:      pe.Pos,
		}
	case pe.Expr != nil:
		return pe.Expr.Transform(scope)
	default:
		panic("unknown primary expression")
	}
}
