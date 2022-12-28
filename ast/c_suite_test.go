package ast_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CSuiteSuite struct {
	CompilerSuite
}

func (suite *CSuiteSuite) TestCSuite1() {
	suite.EqualProgramK(`
	i64
	main()
	{
		return 0;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite2() {
	suite.EqualProgramK(`
	i64
	main()
	{
		return 3-3;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite3() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		
		x = 4;
		return x - 4;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite4() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		i64 *p;
		
		x = 4;
		p = &x;
		*p = 0;
	
		return *p;
	}	
	`, "")
}

func (suite *CSuiteSuite) TestCSuite5_mod() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		i64 *p;
		i64 **pp;
	
		x = 0;
		p = &x;
		pp = &p;
	
		// we support strict bools
		if(*p != 0)
			return 1;

		// we support strict bools
		if(**pp != 0)
			return 1;
		else
			**pp = 1;
	
		// we support strict bools
		if(x != 0)
			return 0;
		else
			return 1;

		// we always need a return
		return 99;
	}	
	`, "")
}

func (suite *CSuiteSuite) TestCSuite6_mod() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
	
		x = 50;
		// we support strict bools
		while (x != 0)
			x = x - 1;
		return x;
	}	
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite7() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64 x;

// 		x = 1;
// 		for(x = 10; x; x = x - 1)
// 			;
// 		if(x)
// 			return 1;
// 		x = 10;
// 		for (;x;)
// 			x = x - 1;
// 		return x;
// 	}
// 	`, "")
// }

// func (suite *CSuiteSuite) TestCSuite8() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64 x;

// 		x = 50;
// 		do
// 			x = x - 1;
// 		while(x);
// 		return x;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite9() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		
		x = 1;
		x = x * 10;
		x = x / 2;
		x = x % 3;
		return x - 2;
	}	
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite10() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		start:
// 			goto next;
// 			return 1;
// 		success:
// 			return 0;
// 		next:
// 		foo:
// 			goto success;
// 			return 1;
// 	}
// 	`, "")
// }

// func (suite *CSuiteSuite) TestCSuite11() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64 x;
// 		i64 y;
// 		x = y = 0;
// 		return x;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite12() {
	suite.EqualProgramK(`
	i64
	main()
	{
		return (2 + 2) * 2 - 8;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite13() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		i64 *p;
		
		x = 0;
		p = &x;
		return p[0];
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite14() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		i64 *p;
		
		x = 1;
		p = &x;
		p[0] = 0;
		return x;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite15() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64[2] arr;

		arr[0] = 1;
		arr[1] = 2;

		return arr[0] + arr[1] - 3;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite16() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64[2] arr;
		i64 *p;

		p = &arr[1];
		*p = 0;
		return arr[1];
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite17() {
	suite.EqualProgramK(`
	i64
	main()
	{
		struct { i64 x, i64 y, } s;
		
		s.x = 3;
		s.y = 5;
		return s.y - s.x - 2; 
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite18_mod() {
	suite.EqualProgramK(`
	type struct { i64 x, i64 y, } S;

	i64
	main()
	{
		S s;
		S *p;

		p = &s;	
		s.x = 1;
		p->y = 2;
		return p->y + p->x - 3; 
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite19_mod1() {
	suite.EqualProgramK(`
	type struct { S *p, i64 x, } S;

	i64
	main()
	{
		S s;

		s.x = 42;
		s.p = &s;
		return 42 - s.p->p->p->p->p->x;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite20() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		i64 *p;
		i64 **pp;
		
		x = 0;
		p = &x;
		pp = &p;
		return **pp;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite21() {
	suite.EqualProgramK(`
	i64
	foo(i64 a, i64 b)
	{
		return 2 + a - b;
	}

	i64
	main()
	{
		return foo(1, 3);
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite22() {
	suite.EqualProgramK(`
	type i64 x;

	i64
	main()
	{
		x v;
		v = (x)0;
		return v;
	}
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite23() {
// 	suite.EqualProgramK(`
// 	i64 x;

// 	i64
// 	main()
// 	{
// 		x = 0;
// 		return x;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite24() {
	suite.EqualProgramK(`
	type struct { i64 x, i64 y, } s;

	i64
	main()
	{
		s v;

		v.x = 1;
		v.y = 2;
		return 3 - v.x - v.y;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite24_mod1() {
	suite.EqualProgramK(`
	i64
	main()
	{
		struct { i64 x, i64 y, } v;

		v.x = 1;
		v.y = 2;
		return 3 - v.x - v.y;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite25() {
	suite.EqualProgramK(`
	i64 strlen(i8* s);

	i64
	main()
	{
		i8 *p;
		
		p = "hello";
		return strlen(p) - 5;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite26() {
	suite.EqualProgramK(`
	i8
	main()
	{
		i8 *p;
		
		p = "hello";
		return p[0] - (i8)104;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite27() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		
		x = 1;
		x = x | 4;
		return x - 5;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite28() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		
		x = 1;
		x = x & 3;
		return x - 1;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite29() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;
		
		x = 1;
		x = x ^ 3;
		return x - 2;
	}
	`, "")
}

// https://github.dev/c-testsuite/c-testsuite/blob/master/tests/single-exec/00030.c
func (suite *CSuiteSuite) TestCSuite30() {
	suite.EqualProgramK(`
	i64
	f()
	{
		return 100;
	}

	i64
	main()
	{
		if (f() > 1000)
			return 1;
		if (f() >= 1000)
			return 1;
		if (1000 < f())
			return 1;
		if (1000 <= f())
			return 1;
		if (1000 == f())
			return 1;
		if (100 != f())
			return 1;
		return 0;
	}
	`, "")
}

// https://github.dev/c-testsuite/c-testsuite/blob/master/tests/single-exec/00031.c
func (suite *CSuiteSuite) TestCSuite31() {
	suite.EqualProgramK(`
	i64
	zero()
	{
		return 0;
	}

	i64
	one()
	{
		return 1;
	}

	i64
	main()
	{
		i64 x;
		i64 y;
		
		x = zero();
		y = ++x;
		if (x != 1)
			return 1;
		if (y != 1)
			return 1;
		
		x = one();	
		y = --x;
		if (x != 0)
			return 1;
		if (y != 0)
			return 1;
		
		x = zero();
		y = x++;
		if (x != 1)
			return 1;
		if (y != 0)
			return 1;
		
		x = one();
		y = x--;
		if (x != 0)
			return 1;
		if (y != 1)
			return 1;
		
		return 0;
	}
	`, "")
}

// https://github.dev/c-testsuite/c-testsuite/blob/master/tests/single-exec/00032.c
func (suite *CSuiteSuite) TestCSuite32() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64[2] arr;
		i64 *p;
		
		arr[0] = 2;
		arr[1] = 3;
		p = &arr[0];
		if(*(p++) != 2)
			return 1;
		if(*(p++) != 3)
			return 2;
		
		p = &arr[1];
		if(*(p--) != 3)
			return 1;
		if(*(p--) != 2)
			return 2;
			
		p = &arr[0];
		if(*(++p) != 3)
			return 1;
		
		p = &arr[1];
		if(*(--p) != 2)
			return 1;

		return 0;
	}
	`, "")
}

func (suite *CSuiteSuite) TestCSuite35() {
	suite.EqualProgramK(`
	i64
	main()
	{
		i64 x;

		x = 4;
		if(!x != 0)
			return 1;
		if(!!x != 1)
			return 1;
		if(-x != 0 - 4)
			return 1;
		return 0;
	}
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite36() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64 x;

// 		x = 0;
// 		x += 2;
// 		x += 2;
// 		if (x != 4)
// 			return 1;
// 		x -= 1;
// 		if (x != 3)
// 			return 2;
// 		x *= 2;
// 		if (x != 6)
// 			return 3;

// 		return 0;
// 	}
// 	`, "")
// }

// func (suite *CSuiteSuite) TestCSuite37() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64[2] x;
// 		i64 *p;

// 		x[1] = 7;
// 		p = &x[0];
// 		p = p + 1;

// 		if(*p != 7)
// 			return 1;
// 		if(&x[1] - &x[0] != 1)
// 			return 1;

// 		return 0;
// 	}
// 	`, "")
// }

// func (suite *CSuiteSuite) TestCSuite38_mod() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		i64 x;
// 		i64 *p;

// 		if (sizeof(0) < 2)
// 			return 1;
// 		if (sizeof 0 < 2)
// 			return 1;
// 		if (sizeof(char) < 1)
// 			return 1;
// 		if (sizeof(i64) - 2 < 0)
// 			return 1;
// 		if (sizeof(&x) != sizeof p)
// 			return 1;
// 		return 0;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite39_mod() {
	suite.EqualProgramK(`
	i64
	main()
	{
		// we don't have void pointers, so we use i8*
		// void *p;
		i8 *p;
		i64 x;
		
		x = 2;
		p = (i8*)&x;
		
		if(*((i64*)p) != 2)
			return 1;
		return 0;
	}
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite40() {
// 	suite.EqualProgramK(`
// 	#include <stdlib.h>

// 	int N;
// 	int *t;

// 	int
// 	chk(int x, int y)
// 	{
// 			int i;
// 			int r;

// 			for (r=i=0; i<8; i++) {
// 					r = r + t[x + 8*i];
// 					r = r + t[i + 8*y];
// 					if (x+i < 8 & y+i < 8)
// 							r = r + t[x+i + 8*(y+i)];
// 					if (x+i < 8 & y-i >= 0)
// 							r = r + t[x+i + 8*(y-i)];
// 					if (x-i >= 0 & y+i < 8)
// 							r = r + t[x-i + 8*(y+i)];
// 					if (x-i >= 0 & y-i >= 0)
// 							r = r + t[x-i + 8*(y-i)];
// 			}
// 			return r;
// 	}

// 	int
// 	go(int n, int x, int y)
// 	{
// 			if (n == 8) {
// 					N++;
// 					return 0;
// 			}
// 			for (; y<8; y++) {
// 					for (; x<8; x++)
// 							if (chk(x, y) == 0) {
// 									t[x + 8*y]++;
// 									go(n+1, x, y);
// 									t[x + 8*y]--;
// 							}
// 					x = 0;
// 			}
// 		return 0;
// 	}

// 	int
// 	main()
// 	{
// 			t = calloc(64, sizeof(int));
// 			go(0, 0, 0);
// 			if(N != 92)
// 				return 1;
// 			return 0;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite41() {
	suite.EqualProgramK(`
	i64 printf(i8 *fmt,... );
	i64
	main() {
		i64 n;
		i64 t;
		i64 c;
		i64 p;

		c = 0;
		n = 2;
		while (n < 5000) {
			t = 2;
			p = 1;
			while (t*t <= n) {
				if (n % t == 0)
					p = 0;
				t++;
			}
			n++;
			// if statments condition must be a boolean
			if (p != 0)
				c++;
		}
		if (c != 669)
			return 1;
		return 0;
	}
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite42() {
// 	suite.EqualProgramK(`
// 	i64
// 	main()
// 	{
// 		union { i64 a; i64 b; } u;
// 		u.a = 1;
// 		u.b = 3;

// 		if (u.a != 3 || u.b != 3)
// 			return 1;
// 		return 0;
// 	}
// 	`, "")
// }

func (suite *CSuiteSuite) TestCSuite43_mod() {
	suite.EqualProgramK(`
	type struct {
		i64 x,
		struct {
			i64 y,
			i64 z,
		} nest,
	} s;
	
	i64
	main() {
		s v;
		v.x = 1;
		v.nest.y = 2;
		v.nest.z = 3;
		if (v.x + v.nest.y + v.nest.z != 6)
			return 1;
		return 0;
	}
	`, "")
}

// func (suite *CSuiteSuite) TestCSuite44() {
// 	suite.EqualProgramK(`
// 	struct T;

// 	struct T {
// 		i64 x;
// 	};

// 	i64
// 	main()
// 	{
// 		struct T v;
// 		{ struct T { i64 z; }; }
// 		v.x = 2;
// 		if(v.x != 2)
// 			return 1;
// 		return 0;
// 	}
// 	`, "")
// }

func TestCSuiteSuite(t *testing.T) {
	suite.Run(t, new(CSuiteSuite))
}
