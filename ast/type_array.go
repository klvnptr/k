package ast

import (
	"fmt"
)

type ArrayType struct {
	Type *Type
	Len  int
}

func (at *ArrayType) String() string {
	return fmt.Sprintf("%s[%d]", at.Type.String(), at.Len)
}

func (at *ArrayType) Equals(o *ArrayType) bool {
	if at.Len != o.Len {
		return false
	}

	return at.Type.Equals(o.Type)
}
