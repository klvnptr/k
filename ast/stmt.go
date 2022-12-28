package ast

import (
	"fmt"

	"github.com/klvnptr/k/utils"
)

func (a *ExprStmt) String() []string {
	return []string{"expr " + a.Expr.String() + ";"}
}

func (a *ExprStmt) Generate() error {
	_, err := a.Expr.Value()

	if err != nil {
		return err
	}

	return nil
}

func (d *DeclStmt) String() []string {
	if d.Expr != nil {
		return []string{"decl " + d.Type.String() + " " + d.Ident + " = " + d.Expr.String() + ";"}
	}

	return []string{"decl " + d.Type.String() + " " + d.Ident + ";"}
}

func (d *DeclStmt) Generate() error {
	typ, err := d.Type.IRType()
	if err != nil {
		return err
	}

	ptr := d.Scope.BasicBlock().NewAlloca(typ)

	v := &Variable{
		Ident: d.Ident,
		Type:  d.Type,
		Ptr:   ptr,
		Pos:   d.Pos,
	}

	err = d.Scope.AddLocal(v)
	if err != nil {
		return utils.WithPos(err, d.Scope.Current().File, d.Pos)
	}

	if d.Expr != nil {
		expr, err := d.Expr.Value()
		if err != nil {
			return err
		}

		if !expr.Type.Equals(d.Type) {
			return utils.WithPos(fmt.Errorf("cannot assign %s to %s", expr.Type.String(), d.Type.String()), d.Scope.Current().File, d.Pos)
		}

		d.Scope.BasicBlock().NewStore(expr.Value, ptr)
	}

	return nil
}

func (a *AssignStmt) String() []string {
	return []string{"assign " + a.Left.String() + " = " + a.Right.String() + ";"}
}

func (a *AssignStmt) Generate() error {
	left, err := a.Left.Value()
	if err != nil {
		return err
	}

	right, err := a.Right.Value()
	if err != nil {
		return err
	}

	if !left.Type.Equals(right.Type) {
		return utils.WithPos(fmt.Errorf("cannot assign %s to %s", right.Type.String(), left.Type.String()), a.Scope.Current().File, a.Pos)
	}

	if left.Ptr == nil {
		return utils.WithPos(fmt.Errorf("cannot assign to non-variable"), a.Scope.Current().File, a.Pos)
	}

	a.Scope.BasicBlock().NewStore(right.Value, left.Ptr)

	return nil
}

func (r *ReturnStmt) String() []string {
	return []string{"return " + r.Expr.String() + ";"}
}

func (r *ReturnStmt) Generate() error {
	val, err := r.Expr.Value()

	if err != nil {
		return err
	}

	r.Scope.BasicBlock().NewRet(val.Value)

	return nil
}

func (i *IfStmt) String() []string {
	lines := []string{"if (" + i.Condition.String() + ") then"}

	for _, stmt := range i.Then {
		thenLines := stmt.String()

		// if the statement is a block, don't indent it
		if _, ok := stmt.(*Block); !ok {
			i.Scope.Current().PrefixLines(thenLines)
		}

		lines = append(lines, thenLines...)
	}

	if len(i.Else) > 0 {
		lines = append(lines, "else")

		for _, stmt := range i.Else {
			elseLines := stmt.String()

			// if the statement is a block, don't indent it
			if _, ok := stmt.(*Block); !ok {
				i.Scope.Current().PrefixLines(elseLines)
			}

			lines = append(lines, elseLines...)
		}
	}

	return lines
}

func (i *IfStmt) Generate() error {
	expr, err := i.Condition.Value()
	if err != nil {
		return err
	}

	thenBlock := i.Scope.CurrentFunction().Ptr.NewBlock(i.Scope.CurrentModule().GenerateID("if.then"))
	elseBlock := i.Scope.CurrentFunction().Ptr.NewBlock(i.Scope.CurrentModule().GenerateID("if.else"))
	mergeBlock := i.Scope.CurrentFunction().Ptr.NewBlock(i.Scope.CurrentModule().GenerateID("if.merge"))

	// c := i.Scope.Current()

	// v := expr.Value

	// if !expr.Type.IsBool() {
	// 	if expr.Type.IsInt() || expr.Type.IsUInt() {
	// 		irType := expr.Type.LLVMIntType()
	// 		v = c.BasicBlock.NewICmp(enum.IPredNE, expr.Value, constant.NewInt(irType, 0))
	// 	} else if expr.Type.IsFloat() {
	// 		irType := expr.Type.LLVMFloatType()
	// 		v = c.BasicBlock.NewFCmp(enum.FPredONE, expr.Value, constant.NewFloat(irType, 0))
	// 	} else {
	// 		return utils.WithPos(fmt.Errorf("cannot use %s as condition", expr.Type.String()), c.File, i.Pos)
	// 	}
	// }

	// entry block
	i.Scope.BasicBlock().NewCondBr(expr.Value, thenBlock, elseBlock)

	// then block
	i.Scope.SetBasicBlock(thenBlock)
	// c.BasicBlock = thenBlock
	for _, stmt := range i.Then {
		if err := stmt.Generate(); err != nil {
			return err
		}
	}
	// if the last statement in the then block doesn't terminate the block, add a branch to the merge block
	// this is necessary because all basic blocks must terminate
	if i.Scope.BasicBlock().Term == nil {
		i.Scope.BasicBlock().NewBr(mergeBlock)
	}

	// else block
	i.Scope.SetBasicBlock(elseBlock)
	// c.BasicBlock = elseBlock
	for _, stmt := range i.Else {
		if err := stmt.Generate(); err != nil {
			return err
		}
	}

	if i.Scope.BasicBlock().Term == nil {
		i.Scope.BasicBlock().NewBr(mergeBlock)
	}

	// merge block
	// c.BasicBlock = mergeBlock
	i.Scope.SetBasicBlock(mergeBlock)

	// fmt.Println("if, merge:", c.BasicBlock.Name())

	return nil
}

func (w *WhileStmt) String() []string {
	lines := []string{"while (" + w.Condition.String() + ")"}

	for _, stmt := range w.Body {
		bodyLines := stmt.String()

		// if the statement is a block, don't indent it
		if _, ok := stmt.(*Block); !ok {
			w.Scope.Current().PrefixLines(bodyLines)
		}

		lines = append(lines, bodyLines...)
	}

	return lines
}

func (w *WhileStmt) Generate() error {
	// c := w.Scope.Current()

	entryBlock := w.Scope.CurrentFunction().Ptr.NewBlock(w.Scope.CurrentModule().GenerateID("while.entry"))
	loopBlock := w.Scope.CurrentFunction().Ptr.NewBlock(w.Scope.CurrentModule().GenerateID("while.loop"))
	mergeBlock := w.Scope.CurrentFunction().Ptr.NewBlock(w.Scope.CurrentModule().GenerateID("while.merge"))

	// c.BasicBlock.NewBr(entryBlock)
	w.Scope.BasicBlock().NewBr(entryBlock)

	// entry block
	// c.BasicBlock = entryBlock
	w.Scope.SetBasicBlock(entryBlock)
	expr, err := w.Condition.Value()
	if err != nil {
		return err
	}

	w.Scope.BasicBlock().NewCondBr(expr.Value, loopBlock, mergeBlock)
	// c.BasicBlock.NewCondBr(expr.Value, loopBlock, mergeBlock)

	// loop block
	// c.BasicBlock = loopBlock
	w.Scope.SetBasicBlock(loopBlock)
	for _, stmt := range w.Body {
		if err := stmt.Generate(); err != nil {
			return err
		}
	}
	// fmt.Println("while, loop:", c.BasicBlock.Name())

	if w.Scope.BasicBlock().Term == nil {
		w.Scope.BasicBlock().NewBr(entryBlock)
	}

	// merge block
	// c.BasicBlock = mergeBlock
	w.Scope.SetBasicBlock(mergeBlock)

	return nil
}
