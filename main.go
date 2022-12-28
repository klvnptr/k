package main

import (
	"fmt"
	"strings"

	"github.com/alecthomas/repr"
	"github.com/klvnptr/k/ast"
	"github.com/klvnptr/k/parser"
	"github.com/klvnptr/k/utils"
)

func main() {

	src := `
i32 printf(i8 *fmt,... );

i64 fib(i64 **n) {
	printf(2);
	if (n < 2) return n;
	else { i32 t = 2; i32 f = t * -4.0;}
	{
		i32 k = 2;
		printf(k);
		i32 z = 2;
	}
	return fib(n - 1) + fib(n - 2);
}

i32 main() {
	return fib(10);
}
	`

	p := parser.NewParser()

	mainFile := utils.NewFile("main.c", src)
	mainScope := ast.NewScope(mainFile)

	module, err := p.ParseFile(mainFile)
	if err != nil {
		// repr.Print(err)
		panic(err)
	}
	// repr.Println("parser.module", module)

	tmodule := module.Transform(mainScope)

	repr.Println("transformed module", tmodule, repr.OmitEmpty(true))

	tcode := tmodule.String()
	tstr := strings.Join(tcode, "\n")

	fmt.Println(tstr)

	// fmt.Println(p)

}
