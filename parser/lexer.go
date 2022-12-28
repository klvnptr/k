package parser

import "github.com/alecthomas/participle/v2/lexer"

func BuildLexer() *lexer.StatefulDefinition {
	return lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "Comment", Pattern: `(?:#|//)[^\n]*\n?`, Action: nil},
			{Name: `StringStart`, Pattern: `"`, Action: lexer.Push("String")},
			{Name: `CharStart`, Pattern: `'`, Action: lexer.Push("Char")},
			{Name: "Number", Pattern: `(\d*\.)?\d+`, Action: nil},
			{Name: "BasicType", Pattern: `(bool|void|i8|i16|i32|i64|u8|u16|u32|u64|f32|f64)`, Action: nil},
			{Name: "Keyword", Pattern: `(if|else|while|type|return|sizeof|const|struct)`, Action: nil},
			{Name: "Ident", Pattern: `\w+`, Action: nil},
			// source: https://github.com/alecthomas/participle#stateful-lexer
			{Name: "Punct", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, Action: nil},
			{Name: "Whitespace", Pattern: `[\n\r\s]+`, Action: nil},
		},
		"String": {
			{Name: "Escaped", Pattern: `\\.`, Action: nil},
			{Name: "StringEnd", Pattern: `"`, Action: lexer.Pop()},
			{Name: "Chars", Pattern: `[^"\\]+`, Action: nil},
		},
		"Char": {
			{Name: "CharEnd", Pattern: `'`, Action: lexer.Pop()},
			{Name: "SingleChar", Pattern: `[^'\\]{1}`, Action: nil},
		},
	})
}
