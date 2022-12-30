package testing

import "fmt"

func ExprToProgramK(expr, format string, opts []Option) string {
	declares := JoinDeclares(opts)

	src := fmt.Sprintf(`
%s
i64 printf(i8 *fmt,... );

i64 main() {
	%s
	printf(%s);
	return 0;
}
`, declares, expr, format)

	return src
}

func ExprToProgramC(expr, format string, opts []Option) string {
	headers := JoinHeaders(opts)

	src := fmt.Sprintf(`
#include <stdio.h>
%s

int main() {
	%s
	printf(%s);
	return 0;
}
`, headers, expr, format)

	return src
}
