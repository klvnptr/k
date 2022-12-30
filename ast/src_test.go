package ast_test

import (
	t "testing"

	"github.com/klvnptr/k/testing"

	"github.com/stretchr/testify/suite"
)

type SrcTestSuite struct {
	testing.CompilerSuite
}

func (suite *SrcTestSuite) SetupTest() {
	suite.T().Parallel()
}

func (suite *SrcTestSuite) TestBasic() {
	src := `
i64 printf(i8 *fmt,... );

i64 main() {
	printf("10");
	return 0;
}
`
	suite.EqualProgramK(src, `10`)
}

func (suite *SrcTestSuite) TestBasicC() {
	src := `
#include <stdio.h>

int main() {
	printf("10");
	return 0;
}
`
	suite.EqualProgramC(src, `10`)
}

func (suite *SrcTestSuite) TestBasicInline() {
	suite.EqualExprK("i64 t = 10;", `"%d", t`, "10")
}

func (suite *SrcTestSuite) TestMathInline() {
	suite.EqualExprK("f64 t = round(10.6);", `"%.2f", t`, "11.00", testing.Declare("f64 round(f64 arg);"))
}

func (suite *SrcTestSuite) TestBasicInlineC() {
	suite.EqualExprC("int t = 10;", `"%d", t`, "10")
}

func (suite *SrcTestSuite) TestMathInlineC() {
	suite.EqualExprC("float t = roundf(10.6f);", `"%.2f", t`, "11.00", testing.Header("math.h"))
}

func (suite *SrcTestSuite) TestDoublePrefixC() {
	suite.ErrorClangExprC(`char *s = "hello"; ++(++s);`, "expression is not assignable")

	suite.EqualExprC(`char *s = "hello"; char *s0 = ++s; char *s1 = ++s;`, `"%s,%s", s0, s1`, "ello,llo")
	suite.EqualExprK(`i8 *s = "hello";`, `"%s,%s,%s", ++s, ++(++s), s`, "ello,lo,lo")

	suite.EqualExprC(`char *s = "hello"; char *p = s;`, `"%s,%s", p + 2, s`, "llo,hello")
	suite.EqualExprK(`i8 *s = "hello"; i8 *p = s;`, `"%s,%s", p + 2, s`, "llo,hello")

	// https://stackoverflow.com/questions/381542/with-arrays-why-is-it-the-case-that-a5-5a
	suite.EqualExprC(`char *s = "12345"; char p = s[2];`, `"%c", p`, "3")
	suite.EqualExprK(`i8 *s = "12345"; i8 p = s[2];`, `"%c", p`, "3")

	suite.EqualExprK(`i8 s = 'B';`, `"%c", s`, "B")
}

func (suite *SrcTestSuite) Test0() {
	suite.EqualExprC(`int a = 10; int *b = &a;`, `"%d", *b`, "10")
	suite.EqualExprK(`i64 a = 10; i64 *b = &a;`, `"%d", *b`, "10")
}

func (suite *SrcTestSuite) Test1() {
	suite.EqualExprC(`char *s = "12345";`, `"%c", *(s + 2)`, "3")
	suite.EqualExprK(`i8 *s = "12345";`, `"%c", *(s + 2)`, "3")

	suite.EqualExprC(`char *s = "12345";`, `"%s", &s[2]`, "345")
	suite.EqualExprK(`i8 *s = "12345";`, `"%s", &s[2]`, "345")

	suite.EqualExprC(`char *s = malloc(10); memset(s, 'A', 6); s[2] = 'X'; s[5] = 0;`, `"%s-", s`, "AAXAA-", testing.Header("stdlib.h"), testing.Header("string.h"))
	suite.EqualExprK(`i8 *s = malloc(10); memset(s, 'A', 6); s[2] = 'X'; s[5] = (i8)0;`, `"%s-", s`, "AAXAA-",
		testing.DeclareMalloc(),
	)

	suite.EqualExprK(`i8 *s = "12345"; *(s + 1) = 'C';`, `"%s", s`, "1C345")
	suite.EqualExprK(`i8 *s = "12345"; i8 *p = s; *(++p) = 'F'; p[2] = 'X';`, `"%s", s`, "1F3X5")
}

func (suite *SrcTestSuite) Test2() {
	suite.EqualExprC(`char *s = "12345"; char **p = &s;`, `"%c", *(*p + 2)`, "3")
	suite.EqualExprK(`i8 *s = "12345"; i8 **p = &s;`, `"%c", *(*p + 2)`, "3")

	suite.EqualExprC(`char *s = "12345"; char **p = malloc(sizeof(char*) * 3); p[1] = s;`, `"%c,%c", *(p[1] + 2), p[1][1]`, "3,2", testing.Header("stdlib.h"))
	suite.EqualExprC(`char *s = "12345"; char **p = malloc(sizeof(char*) * 3); p[1] = s;`, `"%c,%c", *(p[1] + 2),(p[1])[1]`, "3,2", testing.Header("stdlib.h"))
	suite.EqualExprK(`i8 *s = "12345"; i8 **p = malloc(sizeof(i8*) * 3); p[1] = s;`, `"%c,%c", *(p[1] + 2), p[1][1]`, "3,2",
		testing.Declare("i8** malloc(i64 size);"),
	)
	suite.EqualExprK(`i8 *s = "12345"; i8 **p = malloc(sizeof(i8*) * 3); p[1] = s;`, `"%c,%c", *(p[1] + 2), (p[1])[1]`, "3,2",
		testing.Declare("i8** malloc(i64 size);"),
	)
}

func (suite *SrcTestSuite) TestMallocString() {
	c := `
#include <stdlib.h>
#include <stdio.h>

int main() {
	char *s = malloc(6);
	s[0] = 'h';
	s[1] = 'e';
	s[2] = 'l';
	s[3] = 'l';
	s[4] = 'o';
	s[5] = 0;
	printf("%s", &s[2]);
	free(s);
	return 0;
}
`
	suite.EqualProgramC(c, "llo")

	k := `
i8* malloc(i64 size);
i8 free(i8* ptr);
i8* memset(i8* ptr, i8 val, i64 size);
i64 printf(i8 *fmt,... );

i64 main() {
	i8 *s = malloc(6);
	memset(s, 'h', 1);
	memset(s + 1, 'e', 1);
	memset(s + 2, 'l', 1);
	memset(s + 3, 'l', 1);
	memset(s + 4, 'o', 1);
	memset(s + 5, (i8)0, 1);
	printf("%s-", &s[2]);
	free(s);
	return 0;
}
`
	suite.EqualProgramK(k, "llo-")
}

func (suite *SrcTestSuite) TestEmptyString() {
	suite.EqualExprK(`i8 *a = "";`, `"-%s-", a`, "--")
	suite.EqualExprK(`i8 *a = "a";`, `"-%s-", a`, "-a-")
}

func (suite *SrcTestSuite) TestAlias0a() {
	suite.ErrorGenerateExprK(`hello a = (hello)10; bool d = a == 10;`, "incompatible types hello and i64", testing.Declare("type i64 hello;"))
	suite.EqualExprK(`hello a = (hello)10; bool d = a == (hello)10;`, `"%d", d`, "1", testing.Declare("type i64 hello;"))
	suite.EqualExprK(`i64 f = 10; hello g = (hello)10; bool h = f == (i64)g;`, `"%d", h`, "1", testing.Declare("type i64 hello;"), testing.Basename("first"))
}

func (suite *SrcTestSuite) TestAlias0b() {
	suite.EqualExprK(`hello a = (hello)20; i8 b = (i8)a; i64 c = (i64)a;`, `"%d,%d,%d", a, b, c`, "20,20,20", testing.Declare("type i32 hello;"), testing.Basename("second"))

}

func (suite *SrcTestSuite) TestAlias1() {
	k := `
	type i64 hello;
	type hello world;
	type world* myptr;
	type bool bb;

	i64 printf(i8 *fmt,... );

	i64 main(hello b) {
		hello a = (hello)10;
		myptr c = (myptr)&a;
		bool d = ((i64)(*c) == 10);
		bb e = (bb)d;
		printf("%d-%c", c[0], e);
		return 0;
	}
	`
	suite.EqualProgramK(k, "10-\x01")
}

func (suite *SrcTestSuite) TestLoop() {
	src := `
	i64 printf(i8 *fmt,... );

	i64 main() {
		i64 a = 10;

		while (a-- > 0) {
			printf("%d\n", a);
		}

		return 0;
	}
	`
	suite.EqualProgramK(src, "9\n8\n7\n6\n5\n4\n3\n2\n1\n0\n")
}

func (suite *SrcTestSuite) TestPointer() {
	src := `
	i8* malloc(i64 size);
	i8 free(i8* ptr);
	i8* memset(i8* ptr, i8 val, i64 size);
	i64 printf(i8 *fmt,... );
	i32 strlen(i8* str);

	i64 fib(i64 n) {
		return n + 42;
	}

	i64 main() {
		i8* p = malloc(sizeof(i32) * 42);
		memset(p, (i8)0, sizeof(i32) * 42);
		memset(p, (i8)66, sizeof(i32) * 8);
		
		i8* c = "12345";
		i64 n = 42;
		printf("%s %d,%d", &c[2], sizeof(i32*), strlen(&p[2]));
		free(p);

		return 0;
	}
	`
	suite.EqualProgramK(src, "345 8,30")
}

func (suite *SrcTestSuite) TestCharCopyPointer() {
	src := `
	i64 printf(i8 *fmt,... );

	i64 main() {
		i8* p = "hello";
		i8* c = p;
		printf("%s", c);
		return 0;
	}
	`
	suite.EqualProgramK(src, "hello")
}

func (suite *SrcTestSuite) TestStruct() {
	src := `
	type struct { i64 a, i64 b, } my0;

	i64 printf(i8 *fmt,... );

	i64 main() {
		struct { i64 c, i64 d, } m1;
		struct { i64 f, i64 g, } m2;
		
		m1.c = 10;
		m1.d = 20;

		m2.f = 30;
		m2.g = 40;

		m1 = m2;
		
		printf("%d,%d,%d", m1.c, m1.d, sizeof(my0));

		return 0;
	}
	`
	suite.EqualProgramK(src, "30,40,16")
}

func (suite *SrcTestSuite) TestAvoidLeftRecursion() {
	suite.EqualExprK(`i64 x = 2 * 3 * 4;`, `"%d", x`, "24")
	suite.EqualExprK(`i64 x = 2;`, `"%d", ++++x`, "4")
	suite.EqualExprK(`i64 x = 10;`, `"%d,%d", x--, x`, "10,9")
	suite.EqualExprC(`int x = 10; int y = x--;`, `"%d,%d", y, x`, "10,9")
	suite.EqualExprK(`i64 x = 10;`, `"%d,%d", x----, x`, "9,8")
	suite.ErrorClangExprC(`int x = 10; int b = x----;`, "expression is not assignable")

	// check if expressions are evaluated from left to right
	suite.EqualExprK(`i64 x = 3 - 1 - 2;`, `"%d", x`, "0")
	suite.EqualExprC(`int x = 3 - 1 - 2;`, `"%d", x`, "0")
	suite.EqualExprK(`i64 x = 20 / 5 / 2;`, `"%d", x`, "2")
	suite.EqualExprC(`int x = 20 / 5 / 2;`, `"%d", x`, "2")

	suite.EqualExprK(`i64 x = 3 -3;`, `"%d", x`, "0")
	suite.EqualExprK(`i64 x = 3 + -3;`, `"%d", x`, "0")
	suite.EqualExprK(`i64 x = 3 - (-3);`, `"%d", x`, "6")
	// suite.EqualExprK(`i64 x = 3 - -3;`, `"%d", x`, "6")

	suite.EqualExprC(`int x = 3 -3;`, `"%d", x`, "0")
	suite.EqualExprC(`int x = 3 - -3;`, `"%d", x`, "6")
}

func TestSrcTestSuite(t *t.T) {
	suite.Run(t, new(SrcTestSuite))
}
