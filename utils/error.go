package utils

import (
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type withPos struct {
	cause error
	File  *File
	Pos   lexer.Position
}

func WithPos(cause error, file *File, pos lexer.Position) error {
	return &withPos{cause, file, pos}
}

func (w *withPos) Message() string {
	message := ""

	// start character position of the line where the error occurred
	start := 0
	for i := 0; i < len(w.File.Contents); i++ {
		c := w.File.Contents[i]

		if c == '\n' {
			// start with the next character
			start = i + 1
		}

		if i == w.Pos.Offset {
			// find end of line
			end := strings.IndexByte(w.File.Contents[start:], '\n')
			if end == -1 {
				end = len(w.File.Contents) - i
			}

			// print the previous line
			message += fmt.Sprintf("\n%s\n", w.File.Contents[start:start+end])

			// pad the next line before the caret
			for j := start; j < i; j++ {
				if w.File.Contents[j] == '\t' {
					message += "\t"
				} else {
					message += " "
				}
			}

			// print caret
			message += "^"
		}
	}

	return message
}

func (w *withPos) Error() string {
	return w.Message() + ": " + w.cause.Error()
}

func (w *withPos) Cause() error {
	return w.cause
}

func (w *withPos) Unwrap() error {
	return w.cause
}

func (w *withPos) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", w.Cause())
			_, err := io.WriteString(s, w.Message())
			if err != nil {
				panic(err)
			}
			return
		}
		fallthrough
	case 's', 'q':
		_, err := io.WriteString(s, w.Error())
		if err != nil {
			panic(err)
		}
	}
}
