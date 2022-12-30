package parser

import (
	"github.com/klvnptr/k/ast"
)

func (s *Stmt) Transform(scope ast.ScopeLike) []ast.StatementLike {
	if s.DeclStmt != nil {
		return []ast.StatementLike{s.DeclStmt.Transform(scope)}
	} else if s.AssignStmt != nil {
		return []ast.StatementLike{s.AssignStmt.Transform(scope)}
	} else if s.ExprStmt != nil {
		return []ast.StatementLike{s.ExprStmt.Transform(scope)}
	} else if s.ReturnStmt != nil {
		return []ast.StatementLike{s.ReturnStmt.Transform(scope)}
	} else if s.CompoundStmt != nil {
		return []ast.StatementLike{s.CompoundStmt.Transform(scope)}
	} else if s.IfStmt != nil {
		return []ast.StatementLike{s.IfStmt.Transform(scope)}
	} else if s.WhileStmt != nil {
		return []ast.StatementLike{s.WhileStmt.Transform(scope)}
	}

	panic("unknown statement")
}

func (a *ExprStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	return &ast.ExprStmt{
		Expr:  a.Expr.Transform(scope),
		Scope: scope,
		Pos:   a.Pos,
	}
}

func (a *DeclStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	ds := &ast.DeclStmt{
		Ident: a.Declarator.Ident,
		Type:  a.Declarator.Type.Transform(scope),
		Scope: scope,
		Pos:   a.Pos,
	}

	if a.Expr != nil {
		ds.Expr = a.Expr.Transform(scope)
	}

	return ds
}

func (a *AssignStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	return &ast.AssignStmt{
		Left:  a.Left.Transform(scope),
		Right: a.Right.Transform(scope),
		Scope: scope,
		Pos:   a.Pos,
	}
}

func (r *ReturnStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	var expr ast.ExpressionLike

	if r.Expr != nil {
		expr = r.Expr.Transform(scope)
	}

	return &ast.ReturnStmt{
		Expr:  expr,
		Scope: scope,
		Pos:   r.Pos,
	}
}

func (c *CompoundStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	childScope := ast.NewScopeFromParent(scope)

	block := &ast.Block{
		Stmts: []ast.StatementLike{},
		Scope: childScope,
		Pos:   c.Pos,
	}

	for _, stmt := range c.Stmts {
		stmt := stmt.Transform(block)
		block.Stmts = append(block.Stmts, stmt...)
	}

	return block
}

func (i *IfStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	var el []ast.StatementLike
	if i.Else != nil {
		el = i.Else.Transform(scope)
	}

	return &ast.IfStmt{
		Condition: i.Condition.Transform(scope),
		Then:      i.Then.Transform(scope),
		Else:      el,
		Scope:     scope,
		Pos:       i.Pos,
	}
}

func (w *WhileStmt) Transform(scope ast.ScopeLike) ast.StatementLike {
	return &ast.WhileStmt{
		Condition: w.Condition.Transform(scope),
		Body:      w.Body.Transform(scope),
		Scope:     scope,
		Pos:       w.Pos,
	}
}
