package ast

import (
	"strings"
)

type StatementLikeList []StatementLike

func (s StatementLikeList) String() []string {
	lines := []string{}
	for _, stmt := range s {
		lines = append(lines, stmt.String()...)
	}
	return lines
}

type ExpressionLikeList []ExpressionLike

func (e ExpressionLikeList) String() string {
	var ret []string
	for _, v := range e {
		ret = append(ret, v.String())
	}
	return strings.Join(ret, ", ")
}
