package ast

import (
	"fmt"
)

func (td *TypeDef) String() []string {
	return []string{fmt.Sprintf("type %s %s", td.Type.String(), td.Alias)}
}
