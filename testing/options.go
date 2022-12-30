package testing

import (
	"fmt"
	"strings"
)

type Option interface{ Option() }

type OptionBasename struct{ File string }

func Basename(file string) *OptionBasename { return &OptionBasename{File: file} }
func (o *OptionBasename) Option()          {}

func OverrideBasename(def string, opts []Option) string {
	for _, o := range opts {
		if ob, ok := o.(*OptionBasename); ok {
			return ob.File
		}
	}

	return def
}

type OptionHeader struct{ Header string }

func Header(header string) *OptionHeader { return &OptionHeader{Header: header} }
func (o *OptionHeader) Option()          {}

func JoinHeaders(opts []Option) string {
	headers := []string{}
	for _, o := range opts {
		if oh, ok := o.(*OptionHeader); ok {
			headers = append(headers, fmt.Sprintf("#include <%s>", oh.Header))
		}
	}

	return strings.Join(headers, "\n")
}

type OptionDeclare struct{ Declare []string }

func Declare(declare string) *OptionDeclare { return &OptionDeclare{Declare: []string{declare}} }
func DeclareMalloc() *OptionDeclare {
	return &OptionDeclare{Declare: []string{
		"i8* malloc(i64 size);",
		"i8 free(i8* ptr);",
		"i8* memset(i8* ptr, i8 val, i64 size);",
	}}
}
func (o *OptionDeclare) Option() {}

func JoinDeclares(opts []Option) string {
	Declares := []string{}
	for _, o := range opts {
		if oh, ok := o.(*OptionDeclare); ok {
			Declares = append(Declares, oh.Declare...)
		}
	}

	return strings.Join(Declares, "\n")
}
