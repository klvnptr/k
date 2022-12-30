package ast_test

import (
	t "testing"

	"github.com/klvnptr/k/testing"

	"github.com/stretchr/testify/suite"
)

type TinyProgramsTestSuite struct {
	testing.CompilerSuite
}

func (suite *TinyProgramsTestSuite) TestFib() {
	src := `
i64 printf(i8 *fmt, ...);

i64 fib(i64 n) {
	if (n < 2) {
		return n;
	}

	return fib(n - 1) + fib(n - 2);
}

i64 main() {
	i64 i = 0;
	
	while (i++ < 6)
		printf("fib(%d) = %d\n", i, fib(i));
	
	return 0;
}
`

	expected := `fib(1) = 1
fib(2) = 1
fib(3) = 2
fib(4) = 3
fib(5) = 5
fib(6) = 8
`
	suite.EqualProgramK(src, expected)
}

func (suite *TinyProgramsTestSuite) TestStrRev() {
	src := `
	i64 printf(i8 *fmt, ...);

	i8* strrev(i8 *s) {
		i8 *p = s;
		i8 *q = s;
		i8 tmp;

		while (*q != (i8)0) {
			q++;
		}

		q--;

		while (p < q) {
			tmp = *p;
			*p = *q;
			*q = tmp;
			p++;
			q--;
		}

		return s;
	}

	i64 main() {
		i8 *s = "hello";
		printf("%s", strrev(s));
		
		return 0;
	}
	`
	suite.EqualProgramK(src, `olleh`)
}

func (suite *TinyProgramsTestSuite) TestBinSearch() {
	src := `
	i8* malloc(i64 size);
	i8 free(i8* ptr);
	i64 printf(i8 *fmt, ...);

	i64 binsearch(i64 x, i64* v, i64 n) {
		i64 low = 0;
		i64 high = n - 1;
		i64 mid;

		while (low <= high) {
			mid = (low + high) / 2;
			if (x < v[mid]) {
				high = mid - 1;
			} else if (x > v[mid]) {
				low = mid + 1;
			} else {
				return mid;
			}
		}

		return -1;
	}

	i64 main() {
		i64* v = (i64*)malloc(sizeof(i64) * 10);
		v[0] = 1;
		v[1] = 2;
		v[2] = 3;
		v[3] = 4;
		v[4] = 5;
		v[5] = 6;
		v[6] = 7;
		v[7] = 8;
		v[8] = 9;
		v[9] = 10;

		i64 x = 5;
		i64 n = 10;

		printf("%d", binsearch(x, v, n));

		return 0;
	}
	`
	suite.EqualProgramK(src, "4")
}

func TestTinyProgramsTestSuite(t *t.T) {
	suite.Run(t, new(TinyProgramsTestSuite))
}
