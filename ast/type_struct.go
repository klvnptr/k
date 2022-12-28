package ast

import (
	"fmt"
	"strings"
)

func (st *StructField) String() string {
	return st.Ident + " " + st.Type.String()
}

func (st *StructType) FindField(field string) (int, *StructField, error) {
	for i, f := range st.Fields {
		if f.Ident == field {
			return i, f, nil
		}
	}

	return 0, nil, fmt.Errorf("field '%s' not found in struct", field)
}

func (st *StructType) String() string {
	var fields []string

	for _, f := range st.Fields {
		fields = append(fields, f.String())
	}

	return fmt.Sprintf("struct { %s, }", strings.Join(fields, ", "))
}

func (st *StructType) Equals(o *StructType) bool {
	if len(st.Fields) != len(o.Fields) {
		return false
	}

	for i, f := range st.Fields {
		if !f.Type.Equals(o.Fields[i].Type) {
			return false
		}
	}

	return true
}
